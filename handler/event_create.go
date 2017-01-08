package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
	"github.com/golang/go/src/pkg/strings"
)

type EventCreate struct {
	Source service.ISource
}

type EventCreateRequest struct {
	Title     string `json:"event_title"`
	StartDate string `json:"event_date_start"`
}

type Responce struct {
	Error   string
	Message string
}

func (h *EventCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.post(w, r)
	} else {
		h.get(w, r)
	}
}

func (h *EventCreate) post(w http.ResponseWriter, r *http.Request) {
	//Parse Request
	ecr := &EventCreateRequest{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(ecr)
	if err != nil {
		invokeResponceErrorWithStatus(w,
			errors.New(fmt.Sprintf("Handler EventCreate parsing request Error: %s", err)), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//Validate
	if err := ecr.Validate(); err != nil {
		invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse(time.RFC822, ecr.StartDate)
	if err != nil {
		invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
		return
	}
	err, id := h.Source.CreateEvent(dto.Event{Title: ecr.Title, StartDate: startDate})
	if err != nil {
		invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
		return
	}

	rsp := Responce{Message: id}
	rd, _ := json.Marshal(rsp)
	w.Write(rd)
}

func (h *EventCreate) get(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/basic.html", "./templates/event/create/content.html"))
	if err := tpl.ExecuteTemplate(w, "basic.html", nil); err != nil {
		panic(err)
	}
}

func (r *EventCreateRequest) Validate() error {
	if strings.Trim(r.Title, " ") == "" {
		return errInvalidEventTitle
	}
	return nil
}
