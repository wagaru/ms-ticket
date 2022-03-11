package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/go-kit/log"
)

func main() {
	var (
		addr = flag.String("http.addr", ":8081", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	repo := NewInMemRepository()

	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "ticket_system",
		Subsystem: "cinema",
		Name:      "request_count",
		Help:      "Number of request received",
	}, []string{"method"})

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "ticket_system",
		Subsystem: "cinema",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds",
	}, []string{"method"})

	svc := NewService(repo)
	svc = NewLoggingMiddleware(logger)(svc)
	svc = NewInstrumentationMiddleware(requestCount, requestLatency)(svc)

	endPts := MakeEndpoints(svc)

	r := MakeHttpHandler(endPts, logger)

	insertData(repo)

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
	logger.Log("error", error)
}

func insertData(repo Repository) {

}
