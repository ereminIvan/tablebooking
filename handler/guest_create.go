package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
)

type GuestCreate struct {
	Source service.ISource
	Random service.IRandom
}

type GuestCreateRequest struct {
	Name       string `json:"guest_name"`
	LastName   string `json:"guest_last_name"`
	IsVIP      bool   `json:"guest_is_vip"`
	EventTitle string `json:"event_title"`
}

func (h *GuestCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Parse Request
	var errs []dto.Error
	gcr := &GuestCreateRequest{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(gcr)
	if err != nil {
		e := dto.ErrorInternalServer
		e.Code = dto.ErrorFBErrorCode
		e.Detail = err.Error()
		errs = append(errs, e)
		log.Printf("Handler GuestCreate parsing request Error: %s", err)
	}
	defer r.Body.Close()

	if err := gcr.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.Response{
			Errors: errs,
			Data:   nil,
		}
		b, _ := json.Marshal(response)
		if _, err = w.Write(b); err != nil {
			panic(err)
		}
		return
	}
	err = h.Source.CreateGuest(
		dto.Guest{IsVIP: gcr.IsVIP, FirstName: gcr.Name, LastName: gcr.LastName, Code: h.Random.Runes(7)},
		dto.Event{Title: gcr.EventTitle},
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := dto.ErrorInternalServer
		e.Code = dto.ErrorFBErrorCode
		errs = append(errs, e)
		log.Printf("Handler GuestCreate failed with %s", err)
		response := dto.Response{
			Errors: errs,
			Data:   nil,
		}
		b, _ := json.Marshal(response)
		if _, err = w.Write(b); err != nil {
			panic(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Validate validate request
func (r *GuestCreateRequest) Validate() error {
	if r.Name == "" {
		return errInvalidGuestName
	}
	if r.LastName == "" {
		return errInvalidGuestLastName
	}
	if r.EventTitle == "" {
		return errInvalidEventTitle
	}
	return nil
}
