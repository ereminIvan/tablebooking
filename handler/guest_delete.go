package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
)

type GuestDelete struct {
	Source service.ISource
}

func (h *GuestDelete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var eTitle string
	var gCode string

	t := strings.Split(r.URL.Path, "/")

	gCode = t[len(t)-1]
	if strings.Trim(gCode, " ") == "" || strings.Trim(eTitle, " ") == "" {

	} else {
		e := dto.Event{Title: eTitle} //todo no idea how to do it currently
		g := dto.Guest{Code: gCode}
		if err := h.Source.DeleteGuest(g, e); err != nil {
			log.Print(err)
		}
	}
}
