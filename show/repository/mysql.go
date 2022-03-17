package repository

import (
	"fmt"
	"time"

	"github.com/wagaru/ticket/pkg/config"
	"github.com/wagaru/ticket/pkg/utils"
	"github.com/wagaru/ticket/show/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDB struct {
	db *gorm.DB
}

func NewMySQLRepo(dsn string) Repository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &mysqlDB{
		db: db,
	}
}

func (m *mysqlDB) StoreMovie(movie *domain.Movie) error {
	res := m.db.Create(movie)
	if err := res.Error; err != nil {
		return fmt.Errorf("StoreMovie failed, %w", err)
	}
	return nil
}

func (m *mysqlDB) FetchMovies(cinemaID uint) (movies []*domain.Movie, err error) {
	var result []*domain.Movie
	res := m.db.Raw(`
		SELECT DISTINCT movies.*
		FROM movies JOIN shows ON movies.id = shows.movie_id
		WHERE shows.start_time BETWEEN ? AND ? AND shows.cinema_hall_id IN (
			SELECT id FROM cinema_halls where cinema_id = ?
		)`, utils.StartOfDay(time.Now().In(config.Location)), utils.EndOfDay(time.Now().Add(time.Hour*24*7).In(config.Location)), cinemaID).Scan(&result)
	if err = res.Error; err != nil {
		return nil, fmt.Errorf("FetchMovies failed, %w", err)
	}
	return result, nil
}

func (m *mysqlDB) StoreCinema(cinema *domain.Cinema) error {
	res := m.db.Create(cinema)
	if err := res.Error; err != nil {
		return fmt.Errorf("StoreCinema failed, %w", err)
	}
	return nil
}

func (m *mysqlDB) FetchAllCinemas() (cinemas []*domain.Cinema, err error) {
	res := m.db.Preload("Halls").Preload("Halls.Seats").Find(&cinemas)
	if err := res.Error; err != nil {
		return nil, fmt.Errorf("FetchAllCinemas failed, %w", err)
	}
	return
}

func (m *mysqlDB) StoreShow(show *domain.Show) error {
	res := m.db.Create(show)
	if err := res.Error; err != nil {
		return fmt.Errorf("StoreShow failed, %w", err)
	}
	return nil
}

func (m *mysqlDB) FetchShows(movieID uint, cinemaID uint) (shows []*domain.Show, err error) {
	res := m.db.Table("shows").Where("movie_id = ?", movieID).Where("cinema_hall_id IN (?)", m.db.Table("cinema_halls").Select("id").Where("cinema_id = ?", cinemaID)).Find(&shows)
	if err = res.Error; err != nil {
		return nil, fmt.Errorf("FetchShows failed, %w", err)
	}
	return
}

func (m *mysqlDB) FetchShowSeats(showID uint) (seats []*domain.CinemaSeat, err error) {
	res := m.db.Table("cinema_seats").Where("cinema_hall_id = (?)", m.db.Table("shows").Select("cinema_hall_id").Where("id = ?", showID)).Find(&seats)
	if err = res.Error; err != nil {
		return nil, fmt.Errorf("FetchShowSeats failed, %w", err)
	}
	return
}
