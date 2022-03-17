package config

import (
	"os"
	"time"

	"github.com/go-kit/log"
)

var (
	ConsulHost = "localhost"
	ConsulPort = "8888"
	Logger     log.Logger
	Location   *time.Location
)

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
	Location, _ = time.LoadLocation("Asia/Taipei")
}
