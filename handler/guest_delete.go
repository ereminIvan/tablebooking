package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
)

type GuestDelete struct {
	Source service.ISource
}

type GuestDeleteRequest struct {
	EventTitle string `json:"event_title"`
	GuestCode  string `json:"guest_code"`
}

func (h *GuestDelete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Parse Request
	gdr := &GuestDeleteRequest{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(gdr)
	if err != nil {
		log.Printf("Handler GuestDelete parsing request Error: %s", err)
	}
	defer r.Body.Close()

	if err := h.Source.DeleteGuest(dto.Guest{Code: gdr.GuestCode}, dto.Event{Title: gdr.EventTitle}); err != nil {
		log.Print(err)
		//todo return error struct
		//todo set header
	}
}
