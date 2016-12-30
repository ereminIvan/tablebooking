package handler

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
)

type GuestCode struct {
	Source service.ISource
}

func (h *GuestCode) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := strings.Trim(r.FormValue("guest_code"), " ")
	log.Printf("Code: %s", code)
	if code == "" {
		log.Printf("Code is empty")
		tpl(w)
	} else {
		for _, c := range r.Cookies() {
			log.Printf("%v", c)
		}
		evs, err := h.Source.GetEvents()
		if err != nil {
			log.Print(err)
			tpl(w)
			//todo alert not found events
			return
		}
		//try to get guest by given code from all events
		var guest dto.Guest
		for _, ev := range evs {
			guest, err := h.Source.GetGuestByCode(code, dto.Event{Title: ev.Title}) //todo assigned to cookie
			if err != nil {
				log.Print(err)
			}
			if guest.FirstName != "" {
				break
			}
		}
		//If no ne guest is found in all events
		if guest.FirstName == "" {
			tpl(w)
			//w.WriteHeader(http.StatusNotFound)
		}
	}
}

func tpl(w http.ResponseWriter) {
	tpl := template.Must(template.ParseFiles(
		"./templates/basic.html",
		"./templates/guest/code/content.html",
	))
	if err := tpl.ExecuteTemplate(w, "basic.html", nil); err != nil {
		panic(err)
	}
}
