package handler

import (
	"html/template"
	"net/http"

	"github.com/ereminIvan/tablebooking/model"
	"github.com/ereminIvan/tablebooking/service"
)

type EventList struct {
	Source service.ISource
}

func (h *EventList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := h.Source.GetEvents()
	if err != nil {
		panic(err)
	}
	data := struct {
		Events model.Events
	}{
		Events: events,
	}
	tpl := template.Must(template.ParseFiles("./templates/basic.html", "./templates/event/list/content.html"))
	if err := tpl.ExecuteTemplate(w, "basic.html", data); err != nil {
		panic(err)
	}
}
