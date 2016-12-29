package handler

import (
	"net/http"
	"regexp"
)

type Route struct {
	Path    string
	Handler IHandler
}

type RouteList []Route

func (rl RouteList) Build() map[IHandler]*regexp.Regexp {
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

type IHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	RouteList map[IHandler]*regexp.Regexp
}

func (h *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	for h, rx := range h.RouteList {
		if rx.MatchString(path) {
			h.ServeHTTP(w, r)
			return
		}
	}
}
