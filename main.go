package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/ereminIvan/tablebooking/dto"
	"github.com/ereminIvan/tablebooking/handler"
	"github.com/ereminIvan/tablebooking/service"

	fb "github.com/ereminIvan/firebase"
)

var dbStorage service.ISource
var random service.IRandom
var appConfig *dto.Config
var routeList map[handler.IHandler]*regexp.Regexp

func init() {
	appConfig = &dto.Config{}
	if err := appConfig.Read("config.json"); err != nil {
		panic(err)
	}
	dbStorage = service.NewStorage(fb.NewDBClient(appConfig.FbDBPath, appConfig.FbDBSecret, false, nil))
	random = service.NewRand()

	routeList = handler.RouteList{
		{Path: "/guest/code", Handler: &handler.GuestCode{Source: dbStorage}},
		{Path: "/event/create", Handler: &handler.EventCreate{Source: dbStorage}},
		{Path: "/event/list", Handler: &handler.EventList{Source: dbStorage}},
		{Path: "/guest/create", Handler: &handler.GuestCreate{Source: dbStorage, Random: random}},
		{Path: "/event/edit/(.*)", Handler: &handler.GuestEdit{Source: dbStorage}},
		{Path: "/event/delete/(.*)", Handler: &handler.GuestDelete{Source: dbStorage}},
		{Path: "/static/", Handler: http.FileServer(http.Dir("./"))},
	}.Prepare()
}

func main() {
	http.Handle("/", &handler.Router{RouteList: routeList})

	log.Printf("Listen and serve with config: %#v", *appConfig)

	http.ListenAndServe(":"+appConfig.Port, nil)
}
