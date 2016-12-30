package handler

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
)

type EventCreate struct {
	Source service.ISource
}

func (h *EventCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	eventTitle := strings.Trim(r.FormValue("event_title"), " ")
	events, err := h.Source.GetEvents()
	for _, e := range events {
		if e.Title == eventTitle {
			log.Printf("Error: Event with this name already exist")
			return
		}
	}
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}
	if eventTitle == "" {
		tpl := template.Must(template.ParseFiles("./templates/basic.html", "./templates/event/create/content.html"))
		if err := tpl.ExecuteTemplate(w, "basic.html", events); err != nil {
			panic(err)
		}
	} else {
		log.Printf("Create event with title: %s", eventTitle)
		h.Source.CreateEvent(dto.Event{Title: eventTitle})
	}
}
