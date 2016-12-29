package handler

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/ereminIvan/tablebooking/service"
)

type GuestCode struct {
	Source service.ISource
}

func (h *GuestCode) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := r.FormValue("guest_code")

	if strings.Trim(code, " ") == "" {
		tpl := template.Must(template.ParseFiles("./templates/basic.html", "./templates/guest/code/content.html"))
		if err := tpl.ExecuteTemplate(w, "basic.html", nil); err != nil {
			panic(err)
		}
	} else {
		guest, err := h.Source.GetGuestByCode(code) //todo assigned to cookie
		if err != nil {
			log.Print(err)
		}
		if guest.FirstName == "" {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
