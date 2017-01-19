package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/ereminIvan/tablebooking/dto"
)

// Route ...
type Route struct {
	Path    string
	Handler IHandler
}

// RouteList ...
type RouteList []Route

// Prepare route list for matching
func (rl RouteList) Prepare() map[IHandler]*regexp.Regexp {
	list := make(map[IHandler]*regexp.Regexp, len(rl))
	for _, r := range rl {
		rx, err := regexp.Compile(r.Path)
		if err != nil {
			panic(err)
		}
		list[r.Handler] = rx
	}
	return list
}

// IHandler ...
type IHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// Router ...
type Router struct {
	RouteList map[IHandler]*regexp.Regexp
}

// ServeHTTP
func (h *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	for h, rx := range h.RouteList {
		if rx.MatchString(path) {
			log.Printf("Serve route path: %s", path)
			//header for ajax cross-domain requests
			w.Header().Add("Access-Control-Allow-Origin", "*")
			h.ServeHTTP(w, r)
			return
		}
	}
	h.invokeNotFound(w, path)
}

func (h *Router) invokeNotFound(w http.ResponseWriter, path string) {
	w.WriteHeader(http.StatusNotFound)
	r := dto.Response{Errors: []dto.Error{
		{
			Title:  "This is not the droid you're looking for",
			Status: http.StatusNotFound,
		},
	}}
	b, _ := json.Marshal(r)
	w.Write(b)
	log.Printf("Route not found: %s", path)
}
