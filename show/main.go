package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"

	"github.com/wagaru/ticket/show/endpoint"
	"github.com/wagaru/ticket/show/repository"
	"github.com/wagaru/ticket/show/service"
	"github.com/wagaru/ticket/show/transport"
)

func main() {
	var (
		addr           = flag.String("http.addr", ":8080", "HTTP Listen Address")
		mysql_host     = flag.String("mysql.Host", "", "mysql host")
		mysql_port     = flag.Int("mysql.port", 0, "mysql port")
		mysql_user     = flag.String("mysql.user", "", "mysql user")
		mysql_password = flag.String("mysql.password", "", "mysql password")
		mysql_db       = flag.String("mysql.db", "", "mysql database")
	)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	flag.Parse()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", *mysql_user, *mysql_password, *mysql_host, mysql_port, *mysql_db)
	repo := repository.NewMySQLRepo(dsn)

	service := service.NewService(repo)

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
