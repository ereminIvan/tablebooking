package dto

import "time"

// Guest dto model
type Guest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsVIP     bool   `json:"is_vip"`
	Code      string `json:"code"`
}

// Guests list dto model
type Guests map[string]Guest

// Table dto model
type Table struct {
	Guests   Guests `json:"guests"`
	IsVIP    bool   `json:"is_vip"`
	Capacity int64  `json:"capacity"`
}

// Tables list dto model
type Tables map[int]Table

// Event dto model
type Event struct {
	Title     string    `json:"title"`
	StartDate time.Time `json:"start_date"`
	Tables    Tables    `json:"tables"`
}

// Events list dto model
type Events map[string]Event
