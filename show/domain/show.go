package domain

import (
	"time"

	"gorm.io/gorm"
)

type City int

const (
	TAIPEI_CITY City = iota
	HSINCHU_CITY
	TAICHUNG_CITY
)

type Movie struct {
	gorm.Model
	Title       string     `json:"title"`
	Duration    int        `json:"duration"`
	Desc        string     `json:"desc"`
	ComeOutDate *time.Time `json:"come_out_date,omitempty"` // use pointer since zero value for time.Time is struct, and omitempty not
}

type Cinema struct {
	gorm.Model
	Name  string        `json:"name"`
	City  City          `json:"city"`
	Halls []*CinemaHall `json:"halls"`
}

type CinemaHall struct {
	gorm.Model
	Name     string        `json:"name"`
	CinemaID uint          `json:"-"`
	Seats    []*CinemaSeat `json:"seats"`
}

type CinemaSeat struct {
	gorm.Model
	CinemaHallID uint   `json:"-"`
	Number       string `json:"number"`
}

type Show struct {
	gorm.Model
	MovieID      uint      `json:"movie_id"`
	CinemaHallID uint      `json:"cinema_hall_id"`
	Date         time.Time `json:"date"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
}
