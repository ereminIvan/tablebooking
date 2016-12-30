package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
	"github.com/golang/go/src/pkg/encoding/json"
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
	if r.Method == http.MethodPost {
		//Parse Request
		gcr := &GuestCreateRequest{}
		d := json.NewDecoder(r.Body)
		err := d.Decode(gcr)
		if err != nil {
			log.Printf("Handler GuestCreate parsing request Error: %s", err)
		}
		defer r.Body.Close()
		log.Printf("%#v", gcr)
		if err := gcr.Validate(); err == nil {
			err := h.Source.CreateGuest(
				dto.Guest{IsVIP: gcr.IsVIP == "true", FirstName: gcr.Name, LastName: gcr.LastName, Code: h.Random.Runes(7)},
				dto.Event{Title: gcr.EventTitle},
			)
			if err != nil {
				log.Printf("Handler GuestCreate failed with %s", err)
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	} else {
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
}

// Validate validate request
func (r *GuestCreateRequest) Validate() error {
	if r.Name != "" {
		return Error{Value: "Не правильное имя гостя"}
	}
	if r.LastName != "" {
		return Error{Value: "Не правильная фамилия гостя"}
	}
	if r.EventTitle != "" {
		return Error{Value: "Событие не задано"}
	}
	return nil
}
