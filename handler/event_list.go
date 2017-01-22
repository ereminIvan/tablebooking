package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
)

type EventList struct {
	Source service.ISource
}

func (h *EventList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var errs []dto.Error
	events, err := h.Source.GetEvents()
	if err != nil {
		e := dto.ErrorInternalServer
		e.Code = dto.ErrorFBErrorCode
		errs = append(errs, e)

		log.Printf("During getting `events` from source error ocured: %s", err)
	}

	data := make([]struct {
		Id string `json:"id"`
		dto.Event
	}, 0, len(events))
	for idx, e := range events {
		data = append(data, struct {
			Id string `json:"id"`
			dto.Event
		}{idx, e})
	}

	response := dto.Response{
		Errors: errs,
		Data:   data,
	}
	b, _ := json.Marshal(response)
	if _, err = w.Write(b); err != nil {
		panic(err)
	}
}
