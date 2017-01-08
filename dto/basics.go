package dto

import "time"

type Guest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsVIP     bool   `json:"is_vip"`
	Code      string `json:"code"`
}

type Guests map[string]Guest

type Table struct {
	Guests   Guests `json:"guests"`
	IsVIP    bool   `json:"is_vip"`
	Capacity int64  `json:"capacity"`
}

type Event struct {
	Title     string    `json:"title"`
	StartDate time.Time `json:"start_date"`
	Tables    []Table   `json:"tables"`
}

type Events map[string]Event
