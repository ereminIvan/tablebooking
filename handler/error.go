package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	//Guest errors
	errInvalidGuestName     = errors.New("Invalid guest name")
	errInvalidGuestLastName = errors.New("Invalid guest last name")

	//Event errors
	errInvalidEventTitle   = errors.New("Invalid event title")
	errDuplicateEventTitle = errors.New("Event with this name already exist")
)

// write json responce with error and status code
func invokeResponceErrorWithStatus(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	rsp := Responce{Error: err.Error()}
	d, _ := json.Marshal(rsp)
	w.Write(d)
}
