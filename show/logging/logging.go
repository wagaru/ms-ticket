package logging

import (
	"time"

	"github.com/go-kit/log"

	"github.com/wagaru/ticket/show/domain"
	"github.com/wagaru/ticket/show/service"
)

type loggingMiddleware struct {
	logger log.Logger
	svc    service.Service
}

func NewLoggingMiddleware(logger log.Logger) service.ServiceMiddleware {
	return func(s service.Service) service.Service {
		return &loggingMiddleware{
			svc:    s,
			logger: logger,
		}
	}
}

func (l *loggingMiddleware) AddMovie(m *domain.Movie) (err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "AddMovie",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return l.svc.AddMovie(m)
}

func (l *loggingMiddleware) GetCinemaMovies(cinemaID uint) (_ []*domain.Movie, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "GetCinemaMovies",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return l.svc.GetCinemaMovies(cinemaID)
}

func (l *loggingMiddleware) AddCinema(cinema *domain.Cinema) (err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "AddCinema",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return l.svc.AddCinema(cinema)
}

func (l *loggingMiddleware) GetCinemas() (_ []*domain.Cinema, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "GetCinemas",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return l.svc.GetCinemas()
}

func (l *loggingMiddleware) AddShow(show *domain.Show) (err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "AddShow",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return l.svc.AddShow(show)
}

func (l *loggingMiddleware) GetShows(movieID uint, cinemaID uint) (_ []*domain.Show, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "GetShows",
			"movieID", movieID,
			"cinemaID", cinemaID,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return l.svc.GetShows(movieID, cinemaID)
}

func (l *loggingMiddleware) GetShowSeats(showID uint) (_ []*domain.CinemaSeat, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "GetShowSeats",
			"showID", showID,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return l.svc.GetShowSeats(showID)
}
