package logging

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/wagaru/ticket/movie/domain"
	"github.com/wagaru/ticket/movie/service"
)

type instrumentationMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           service.Service
}

func InstrumentationMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) Middleware {
	return func(next service.Service) service.Service {
		return &instrumentationMiddleware{
			requestCount:   requestCount,
			requestLatency: requestLatency,
			next:           next,
		}
	}
}

func (im *instrumentationMiddleware) PostMovie(ctx context.Context, movie *domain.Movie) error {
	defer func(begin time.Time) {
		im.requestCount.With("method", "PostMovie").Add(1)
		im.requestLatency.With("method", "PostMovie").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return im.next.PostMovie(ctx, movie)
}

func (im *instrumentationMiddleware) GetMovies(ctx context.Context) (movies []*domain.Movie, err error) {
	defer func(begin time.Time) {
		im.requestCount.With("method", "GetMovies").Add(1)
		im.requestLatency.With("method", "GetMovies").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return im.next.GetMovies(ctx)
}

func (im *instrumentationMiddleware) GetMovie(ctx context.Context, ID string) (movie *domain.Movie, err error) {
	defer func(begin time.Time) {
		im.requestCount.With("method", "GetMovie").Add(1)
		im.requestLatency.With("method", "GetMovie").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return im.next.GetMovie(ctx, ID)
}
