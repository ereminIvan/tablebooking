package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
)

type EventEdit struct {
	Source service.ISource
}

type EventEditRequest struct {
	EventTitle string `json:"event_title"`
}

func (h *EventEdit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Parse Request
	b := &EventEditRequest{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(b)
	if err != nil {
		//Validate
		if err := b.Validate(); err != nil {
			invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
	}
	defer r.Body.Close()

	if r.Method == http.MethodPost {
		//Update event
		if err := h.Source.UpdateEvent(dto.Event{Title: b.EventTitle}); err != nil {
			log.Print(err)
			//todo return error struct
			//todo set header
		}
	} else {
		//h.get(w, r)
	}
}

func (eer *EventEditRequest) Validate() error {
	if eer.EventTitle == "" {
		return errInvalidEventTitle
	}
	return nil
}

//
//
//func (h *EventEdit) get (w http.ResponseWriter, r *http.Request) {
//
//}
