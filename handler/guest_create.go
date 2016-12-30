package handler

import (
	"encoding/json"
	"html/template"
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
	IsVIP      string `json:"guest_is_vip"`
	EventTitle string `json:"event_title"`
}

func (h *GuestCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Create new
	if r.Method == http.MethodPost {
		h.post(w, r)
		//Open creation form
	} else {
		h.get(w, r)
	}
}

func (h *GuestCreate) post(w http.ResponseWriter, r *http.Request) {
	//Parse Request
	gcr := &GuestCreateRequest{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(gcr)
	if err != nil {
		log.Printf("Handler GuestCreate parsing request Error: %s", err)
	}
	defer r.Body.Close()

	log.Printf("%#v", gcr)
	if err := gcr.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.Source.CreateGuest(
		dto.Guest{IsVIP: gcr.IsVIP == "true", FirstName: gcr.Name, LastName: gcr.LastName, Code: h.Random.Runes(7)},
		dto.Event{Title: gcr.EventTitle},
	)
	if err != nil {
		log.Printf("Handler GuestCreate failed with %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *GuestCreate) get(w http.ResponseWriter, r *http.Request) {
	evs, err := h.Source.GetEvents()
	if err != nil {
		log.Printf("Handler GuestCreate Error: %s", err.Error())
	}
	tpl := template.Must(template.ParseFiles(
		"./templates/basic.html",
		"./templates/guest/create/content.html",
	))
	if err := tpl.ExecuteTemplate(w, "basic.html", evs); err != nil {
		panic(err)
	}
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
