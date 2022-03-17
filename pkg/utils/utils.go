package utils

import (
	"time"

	"github.com/wagaru/ticket/pkg/config"
)

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, config.Location)
}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 11, 59, 59, 999, config.Location)
}
