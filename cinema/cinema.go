package main

type City int

const (
	TAIPEI City = iota
	HSINCHU
	TAICHUNG
)

type SeatStatus int

const (
	Available SeatStatus = iota
	NotAvailable
)

type Cinema struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	City  City    `json:"city"`
	Halls []*Hall `json:"halls"`
}

type Hall struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Seats []*Seat `json:"seats"`
}

type Seat struct {
	ID     string     `json:"id"`
	Number string     `json:"number"`
	Status SeatStatus `json:"status"`
}
