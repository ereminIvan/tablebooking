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
	//Read config
	appConfig = &dto.Config{}
	if err := appConfig.Read("config.json"); err != nil {
		panic(err)
	}
	//Init Services
	random = service.NewRand()
	dbStorage = service.NewStorage(
		fb.NewDBClient(appConfig.FbDBPath, appConfig.FbDBSecret, false, nil),
		random,
	)

	//Init Routes
	routeList = handler.RouteList{
		{Path: "/guest/code", Handler: &handler.GuestCode{Source: dbStorage}},
		{Path: "/guest/create", Handler: &handler.GuestCreate{Source: dbStorage, Random: random}},
		{Path: "/event/create", Handler: &handler.EventCreate{Source: dbStorage}},
		{Path: "/event/edit/", Handler: &handler.EventEdit{Source: dbStorage}},
		{Path: "/event/delete/", Handler: &handler.EventDelete{Source: dbStorage}},
		{Path: "/event/list", Handler: &handler.EventList{Source: dbStorage}},
		{Path: "/event/(.*)/table", Handler: &handler.EventTableCreate{Source: dbStorage}},

		{Path: "/", Handler: http.FileServer(http.Dir("./static/build"))},
	}.Prepare()
}

func main() {
	http.Handle("/", &handler.Router{RouteList: routeList})

	log.Printf("Listen and serve with config: %#v", *appConfig)

	if err := http.ListenAndServe(":"+appConfig.Port, nil); err != nil {
		panic(err)
	}
}
