package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
)

type EventCreate struct {
	Source service.ISource
}

type EventCreateRequest struct {
	Title     string    `json:"event_title"`
	StartDate time.Time `json:"event_date_start"`
}

func (h *EventCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.post(w, r)
	} else {
		h.get(w, r)
	}
}

func (h *EventCreate) post(w http.ResponseWriter, r *http.Request) {
	//Parse Request
	ecr := &EventCreateRequest{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(ecr)
	if err != nil {
		log.Printf("Handler EventCreate parsing request Error: %s", err)
	}
	defer r.Body.Close()

	log.Printf("%#v", ecr)
	if err := ecr.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	events, err := h.Source.GetEvents()
	for _, e := range events {
		if e.Title == ecr.Title {
			log.Printf("Error: Event with this name already exist")
			return
		}
	}
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}
	h.Source.CreateEvent(dto.Event{Title: ecr.Title, StartDate: ecr.StartDate})
}

func (h *EventCreate) get(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/basic.html", "./templates/event/create/content.html"))
	if err := tpl.ExecuteTemplate(w, "basic.html", nil); err != nil {
		panic(err)
	}
}

func (r *EventCreateRequest) Validate() error {
	if r.Title == "" {
		return Error{Value: "Incorrect event title"}
	}
	return nil
}
