package handler

import (
	"net/http"

	"github.com/ereminIvan/tablebooking/service"
)

type GuestEdit struct {
	Source service.ISource
}

func (h *GuestEdit) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
