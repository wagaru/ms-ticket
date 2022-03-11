package main

import (
	"time"

	"github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

func NewLoggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(s Service) Service {
		return &loggingMiddleware{
			logger: logger,
			svc:    s,
		}
	}
}

func (lm *loggingMiddleware) PostCinema(cinema *Cinema) (err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"method", "PostCinema",
			"input", cinema,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return lm.svc.PostCinema(cinema)
}

func (lm *loggingMiddleware) GetCinemasByCity(city City) (_ []*Cinema, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"method", "GetCinemasByCity",
			"input", city,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return lm.svc.GetCinemasByCity(city)
}

func (lm *loggingMiddleware) GetCinemas() (_ []*Cinema, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"method", "GetCinemas",
			"input", "",
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return lm.svc.GetCinemas()
}

func (lm *loggingMiddleware) GetCinema(ID string) (_ *Cinema, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"method", "GetCinema",
			"input", ID,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return lm.svc.GetCinema(ID)
}
