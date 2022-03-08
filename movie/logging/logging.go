package logging

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/wagaru/ticket/movie/service"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Service
}

type Middleware func(service.Service) service.Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (lm *loggingMiddleware) PostMovie(ctx context.Context, movie service.Movie) (err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"method", "PostMovie",
			"input", movie.Title,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return lm.next.PostMovie(ctx, movie)
}

func (lm *loggingMiddleware) GetMovies(ctx context.Context) (movies []service.Movie, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"method", "GetMovies",
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return lm.next.GetMovies(ctx)
}

func (lm *loggingMiddleware) GetMovie(ctx context.Context, ID string) (movie service.Movie, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"method", "GetMovie",
			"ID", ID,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return lm.next.GetMovie(ctx, ID)
}
