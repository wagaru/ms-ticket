package main

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentationMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	svc            Service
}

func NewInstrumentationMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) ServiceMiddleware {
	return func(s Service) Service {
		return &instrumentationMiddleware{
			requestCount:   requestCount,
			requestLatency: requestLatency,
			svc:            s,
		}
	}
}

func (im *instrumentationMiddleware) PostCinema(cinema *Cinema) error {
	defer func(begin time.Time) {
		im.requestCount.With("method", "PostCinema").Add(1)
		im.requestLatency.With("method", "PostCinema").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return im.svc.PostCinema(cinema)
}

func (im *instrumentationMiddleware) GetCinemasByCity(city City) ([]*Cinema, error) {
	defer func(begin time.Time) {
		im.requestCount.With("method", "GetCinemasByCity").Add(1)
		im.requestLatency.With("method", "GetCinemasByCity").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return im.svc.GetCinemasByCity(city)
}

func (im *instrumentationMiddleware) GetCinemas() ([]*Cinema, error) {
	defer func(begin time.Time) {
		im.requestCount.With("method", "GetCinemas").Add(1)
		im.requestLatency.With("method", "GetCinemas").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return im.svc.GetCinemas()
}

func (im *instrumentationMiddleware) GetCinema(ID string) (*Cinema, error) {
	defer func(begin time.Time) {
		im.requestCount.With("method", "GetCinema").Add(1)
		im.requestLatency.With("method", "GetCinema").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return im.svc.GetCinema(ID)
}
