package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/wagaru/ticket/movie/endpoint"
	"github.com/wagaru/ticket/movie/logging"
	"github.com/wagaru/ticket/movie/repository"
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

	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "ticket_system", // space is not allowed
		Subsystem: "movie",
		Name:      "request_count",
		Help:      "Number of request received",
	}, []string{"method"})

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "ticket_system",
		Subsystem: "movie",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds",
	}, []string{"method"})

	repo := repository.NewInMemRepository()

	svc := service.NewService(repo)
	svc = logging.LoggingMiddleware(logger)(svc)
	svc = logging.InstrumentationMiddleware(requestCount, requestLatency)(svc)

	endpoints := endpoint.MakeEndpoints(svc)

	r := transport.MakeHttpHandler(endpoints, logger)

	errChan := make(chan error)
	go func() {
		logger.Log("msg", "start server", "tcp", "http", "listen", *addr)
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
