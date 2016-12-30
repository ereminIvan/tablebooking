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
	IsVIP      bool   `json:"guest_is_vip"`
	EventTitle string `json:"event_title"`
}

func (h *GuestCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//Parse Request
		gcr := &GuestCreateRequest{}
		d := json.NewDecoder(r.Body)
		err := d.Decode(gcr)
		if err != nil {
			log.Printf("Handler GuestCreate parsing request Error: %s", err.Error())
		}
		defer r.Body.Close()

		if gcr.Name != "" && gcr.LastName != "" && gcr.EventTitle != "" {
			err := h.Source.CreateGuest(
				dto.Guest{IsVIP: gcr.IsVIP, FirstName: gcr.Name, LastName: gcr.LastName, Code: h.Random.Runes(7)},
				dto.Event{Title: gcr.EventTitle},
			)
			if err != nil {
				log.Printf("Handler GuestCreate failed with %s", err.Error())
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	} else {
		evs, err := h.Source.GetEvents()
		if err != nil {
			log.Printf("Handler GuestCreate Error: %s", err.Error())
		}
		for t, e := range evs {
			log.Printf("%v   :   %v", t, e)
		}
		tpl := template.Must(template.ParseFiles(
			"./templates/basic.html",
			"./templates/guest/create/content.html",
		))
		if err := tpl.ExecuteTemplate(w, "basic.html", nil); err != nil {
			panic(err)
		}
	}
}
