package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
)

type EventDelete struct {
	Source service.ISource
}

type EventDeleteRequest struct {
	Id string `json:"id"`
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

	b.Id = strings.Trim(b.Id, " ")
	if err := b.Validate(); err != nil {
		invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
		return
	}
	if err := h.Source.DeleteEvent(b.Id); err != nil {
		log.Printf("Firebase error: %s", err)
		invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	rsp := dto.Response{Data: struct {
		Message string `json:"message"`
	}{Message: "Event successfully deleted"}}
	rb, _ := json.Marshal(rsp)
	w.Write(rb)
	return
}

func (r *EventDeleteRequest) Validate() error {
	if r.Id == "" {
		return errors.New("Invalid event Id")
	}
	return nil
}
