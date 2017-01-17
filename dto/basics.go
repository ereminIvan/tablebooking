package dto

import "time"

// Guest DTO model
type Guest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsVIP     bool   `json:"is_vip"`
	Code      string `json:"code"`
}

// Guests list DTO model
type Guests map[string]Guest

// Table DTO model
type Table struct {
	Guests   Guests `json:"guests"`
	IsVIP    bool   `json:"is_vip"`
	Capacity int64  `json:"capacity"`
}

// Tables list DTO model
type Tables map[int]Table

// Event DTO model
type Event struct {
	Title     string    `json:"title"`
	StartDate time.Time `json:"start_date"`
	Tables    Tables    `json:"tables"`
}

// Events list DTO model
type Events map[string]Event
