package model

import "time"

type Guest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsVIP     bool   `json:"is_vip"`
}

type Code string

type Guests map[Code]Guest

type Table struct {
	Guests Guests `json:"guests"`
	IsVIP  bool   `json:"is_vip"`
	Size   int64  `json:"size"`
}

type Event struct {
	Title     string    `json:"title"`
	StartDate time.Time `json:"start_date"`
	Tables    []Table   `json:"tables"`
}

type Events map[string]Event
