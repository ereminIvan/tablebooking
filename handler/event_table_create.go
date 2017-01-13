package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/service"
	"net/http"
)

type EventTableCreate struct {
	Source service.ISource
}

type EventTableCreateRequest struct {
	EventId  string `json:"event_id"`
	Capacity int64  `json:"table_capacity"`
	IsVip    bool   `json:"table_is_vip"`
}

type EventTableCreateResponse struct {
	Id string `json:"id"`
}

func (h *EventTableCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.post(w, r)
	}
}

func (h *EventTableCreate) post(w http.ResponseWriter, r *http.Request) {
	//Parse Request
	ecr := &EventTableCreateRequest{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(ecr)
	if err != nil {
		invokeResponceErrorWithStatus(w,
			errors.New(fmt.Sprintf("Handler EventCreateTable parsing request Error: %s", err)), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//Validate
	if err := ecr.Validate(); err != nil {
		invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
		return
	}
	table := dto.Table{Capacity: ecr.Capacity, IsVIP: ecr.IsVip}
	idx := int(0)
	err = h.Source.CreateEventTables(ecr.EventId, dto.Tables{idx: table})
	if err != nil {
		invokeResponceErrorWithStatus(w, err, http.StatusBadRequest)
		return
	}
	rsp := dto.Response{Data: EventCreateResponse{Id: string(idx)}}
	rd, _ := json.Marshal(rsp)
	w.Write(rd)
}

func (r *EventTableCreateRequest) Validate() error {
	return nil
}
