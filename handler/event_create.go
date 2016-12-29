package handler

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/ereminIvan/tablebooking/model"
	"github.com/ereminIvan/tablebooking/service"
)

type EventCreate struct {
	Source service.ISource
}

func (h *EventCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	eventTitle := r.FormValue("event_title")
	events, err := h.Source.GetEvents()
	if err != nil || len(events) == 0 {
		panic("Ooops")
	}
	if strings.Trim(eventTitle, " ") == "" {
		tpl := template.Must(template.ParseFiles("./templates/basic.html", "./templates/event/create/content.html"))
		if err := tpl.ExecuteTemplate(w, "basic.html", events); err != nil {
			panic(err)
		}
	} else {
		log.Printf("Create event with title: %s", eventTitle)
		h.Source.CreateEvent(model.Event{Title: eventTitle})
	}
}
