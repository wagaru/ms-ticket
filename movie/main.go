package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"

	"github.com/wagaru/ticket/movie/endpoint"
	"github.com/wagaru/ticket/movie/logging"
	"github.com/wagaru/ticket/movie/service"
	"github.com/wagaru/ticket/movie/transport"
)

func main() {
	var (
		addr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// fieldKeys := []string{"method", "error"}
	// requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
	// 	Namespace: "my_group",
	// 	Subsystem: "movie",
	// 	Name:      "request_count",
	// 	Help:      "Number of requests received.",
	// }, fieldKeys)
	// requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
	// 	Namespace: "my_group",
	// 	Subsystem: "movie",
	// 	Name:      "request_latency_microseconds",
	// 	Help:      "Total duration of requests in microseconds.",
	// }, fieldKeys)
	// countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
	// 	Namespace: "my_group",
	// 	Subsystem: "movie",
	// 	Name:      "count_result",
	// 	Help:      "The result of each count method.",
	// }, []string{})

	svc := service.NewInMemService()
	svc = logging.LoggingMiddleware(logger)(svc)
	// svc = logging.NewInstrumentationMiddleware(requestCount, requestLatency, countResult, svc)

	endpoints := endpoint.MakeEndpoints(svc)

	r := transport.MakeHttpHandler(endpoints, logger)

	errChan := make(chan error)
	go func() {
		errChan <- http.ListenAndServe(*addr, r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	logger.Log("err", error)
}
