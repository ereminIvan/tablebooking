package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ereminIvan/tablebooking/service"
)

type EventList struct {
	Source service.ISource
}

func (h *EventList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := h.Source.GetEvents()
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(events)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(b)
	if err != nil {
		panic(err)
	}
}
