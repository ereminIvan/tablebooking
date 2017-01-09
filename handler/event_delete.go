package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
)

type EventDelete struct {
	Source service.ISource
}

type EventDeleteRequest struct {
	EventTitle string `json:"event_title"`
}

func (h *EventDelete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Parse Request
	b := &EventDeleteRequest{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(b)
	if err != nil {
		invokeResponceErrorWithStatus(w, fmt.Errorf("Handler EventDelete parsing request Error: %s", err), http.StatusBadRequest)
	}
	defer r.Body.Close()

	if err := h.Source.DeleteEvent(dto.Event{Title: b.EventTitle}); err != nil {
		log.Printf("Firebase error: %s", err)
		invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	rsp := dto.Response{Data: "Event successfully deleted"}
	rb, _ := json.Marshal(rsp)
	w.Write(rb)
	return
}
