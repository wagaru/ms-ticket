package config

import (
	"os"

	"github.com/go-kit/log"
)

var (
	ConsulHost = "localhost"
	ConsulPort = "8888"
	Logger     log.Logger
)

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
}
