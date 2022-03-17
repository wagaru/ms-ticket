package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/log"

	"github.com/wagaru/ticket/pkg/config"
	"github.com/wagaru/ticket/show/domain"
	"github.com/wagaru/ticket/show/endpoint"
	"github.com/wagaru/ticket/show/logging"
	"github.com/wagaru/ticket/show/repository"
	"github.com/wagaru/ticket/show/service"
	"github.com/wagaru/ticket/show/transport"
)

func main() {
	var (
		addr           = flag.String("http.addr", ":8080", "HTTP Listen Address")
		mysql_host     = flag.String("mysql.Host", "localhost", "mysql host")
		mysql_port     = flag.Int("mysql.port", 3306, "mysql port")
		mysql_user     = flag.String("mysql.user", "show_admin", "mysql user")
		mysql_password = flag.String("mysql.password", "show_admin", "mysql password")
		mysql_db       = flag.String("mysql.db", "shows", "mysql database")
	)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	flag.Parse()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", *mysql_user, *mysql_password, *mysql_host, *mysql_port, *mysql_db)
	repo := repository.NewMySQLRepo(dsn)

	// insertTestData(repo)

	service := service.NewService(repo)
	service = logging.NewLoggingMiddleware(logger)(service)

	endpts := endpoint.MakeEndpoints(service)

	r := transport.MakeHttpHandler(endpts, logger)

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

	logger.Log("error", <-errChan)

}

func insertTestData(repo repository.Repository) {
	var err error
	comeOutDate := time.Date(2022, time.April, 1, 0, 0, 0, 0, config.Location)
	if err = repo.StoreMovie(&domain.Movie{
		Title:       "The Avengers",
		Duration:    143,
		Desc:        "descxxxxxxxxx",
		ComeOutDate: &comeOutDate,
	}); err != nil {
		panic(err)
	}

	if err = repo.StoreCinema(&domain.Cinema{
		Name: "Viewshow",
		City: domain.HSINCHU_CITY,
		Halls: []*domain.CinemaHall{
			{
				Name: "hall_1",
				Seats: []*domain.CinemaSeat{
					{
						Number: "A1",
					},
					{
						Number: "A2",
					},
					{
						Number: "A3",
					},
					{
						Number: "B1",
					},
					{
						Number: "B2",
					},
					{
						Number: "B3",
					},
				},
			},
			{
				Name: "hall_2",
				Seats: []*domain.CinemaSeat{
					{
						Number: "A1",
					},
					{
						Number: "A2",
					},
					{
						Number: "A3",
					},
					{
						Number: "B1",
					},
					{
						Number: "B2",
					},
					{
						Number: "B3",
					},
				},
			},
		},
	}); err != nil {
		panic(err)
	}

	if err = repo.StoreShow(&domain.Show{
		MovieID:      1,
		CinemaHallID: 1,
		Date:         time.Date(2022, time.March, 17, 0, 0, 0, 0, config.Location),
		StartTime:    time.Date(2022, time.March, 17, 14, 0, 0, 0, config.Location),
		EndTime:      time.Date(2022, time.March, 17, 15, 0, 0, 0, config.Location),
	}); err != nil {
		panic(err)
	}

	if err = repo.StoreShow(&domain.Show{
		MovieID:      1,
		CinemaHallID: 1,
		Date:         time.Date(2022, time.March, 17, 0, 0, 0, 0, config.Location),
		StartTime:    time.Date(2022, time.March, 17, 15, 30, 0, 0, config.Location),
		EndTime:      time.Date(2022, time.March, 17, 16, 30, 0, 0, config.Location),
	}); err != nil {
		panic(err)
	}

	if err = repo.StoreShow(&domain.Show{
		MovieID:      1,
		CinemaHallID: 1,
		Date:         time.Date(2022, time.March, 18, 0, 0, 0, 0, config.Location),
		StartTime:    time.Date(2022, time.March, 18, 14, 0, 0, 0, config.Location),
		EndTime:      time.Date(2022, time.March, 18, 15, 0, 0, 0, config.Location),
	}); err != nil {
		panic(err)
	}

	if err = repo.StoreShow(&domain.Show{
		MovieID:      1,
		CinemaHallID: 1,
		Date:         time.Date(2022, time.March, 18, 0, 0, 0, 0, config.Location),
		StartTime:    time.Date(2022, time.March, 18, 15, 30, 0, 0, config.Location),
		EndTime:      time.Date(2022, time.March, 18, 16, 30, 0, 0, config.Location),
	}); err != nil {
		panic(err)
	}
}
