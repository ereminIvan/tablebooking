package handler

import (
	"encoding/json"
	"errors"
	"github.com/ereminIvan/tablebooking/dto"
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
func invokeResponceErrorWithStatus(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	rsp := dto.Response{
		Errors: []dto.Error{
			{
				Title:  err.Error(),
				Detail: err.Error(),
				Code:   string(code),
			},
		},
	}
	d, _ := json.Marshal(rsp)
	w.Write(d)
}
