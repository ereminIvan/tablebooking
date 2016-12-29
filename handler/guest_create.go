package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ereminIvan/tablebooking/model"
	"github.com/ereminIvan/tablebooking/service"
)

type GuestCreate struct {
	Source service.ISource
	Random service.IRandom
}

func (h *GuestCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("guest_name")
	lastname := r.FormValue("guest_lastname")
	isVIP := r.FormValue("guest_is_vip")

	if name == "" || lastname == "" {
		tpl := template.Must(template.ParseFiles("./templates/basic.html", "./templates/guest/create/content.html"))
		if err := tpl.ExecuteTemplate(w, "basic.html", nil); err != nil {
			panic(err)
		}
	} else {
		err := h.Source.CreateGuest(
			model.Guest{IsVIP: isVIP == "true", FirstName: name, LastName: lastname},
			h.Random.Runes(7),
		)
		if err != nil {
			log.Printf("CreateGuest failed with %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
