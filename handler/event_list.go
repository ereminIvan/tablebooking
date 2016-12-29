package handler

import (
	"html/template"
	"net/http"

	"github.com/ereminIvan/tablebooking/service"
)

type EventList struct {
	Source service.ISource
}

func (h *EventList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/basic.html", "./templates/event/list/content.html"))
	if err := tpl.ExecuteTemplate(w, "basic.html", nil); err != nil {
		panic(err)
	}
}
