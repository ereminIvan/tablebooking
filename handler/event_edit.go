package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
	"html/template"
)

type EventEdit struct {
	Source service.ISource
}

type EventEditRequest struct {
	EventTitle string `json:"event_title"`
}

func (h *EventEdit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == http.MethodPost {
		//Parse Request
		b := &EventEditRequest{}
		d := json.NewDecoder(r.Body)
		err := d.Decode(b)
		if err != nil {
			invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
		if err := b.Validate(); err != nil {
			invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
		//Update event
		if err := h.Source.UpdateEvent(dto.Event{Title: b.EventTitle}); err != nil {
			log.Print(err)
			//todo return error struct
			//todo set header
		}
	} else {
		u, err := url.Parse(r.RequestURI)
		if err != nil {
			invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
		id := path.Base(u.Path)
		event, err := h.Source.GetEvent(id)
		if err != nil {
			invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
		data := struct {
			EventId string
			Event   dto.Event
		}{
			EventId: id,
			Event:   event,
		}
		tpl := template.Must(template.ParseFiles("./templates/basic.html", "./templates/event/edit/content.html"))
		if err := tpl.ExecuteTemplate(w, "basic.html", data); err != nil {
			panic(err)
		}
		log.Printf("%#v", event)
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
