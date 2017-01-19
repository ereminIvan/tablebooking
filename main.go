package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/ereminIvan/tablebooking/dto"
	h "github.com/ereminIvan/tablebooking/handler"
	"github.com/ereminIvan/tablebooking/service"

	fb "github.com/ereminIvan/firebase"
)

var dbStorage service.ISource
var random service.IRandom
var appConfig *dto.Config
var routes map[h.IHandler]*regexp.Regexp

const configFilePath = "config.json"

func init() {
	//Read config
	appConfig = &dto.Config{}
	if err := appConfig.Read(configFilePath); err != nil {
		panic(err)
	}

	//Init Services
	random = service.NewRand()
	dbStorage = service.NewStorage(
		fb.NewDBClient(appConfig.FbDBPath, appConfig.FbDBSecret, false, nil),
		random,
	)

	//Init Routes
	routes = h.RouteList{
		{Path: "/guest/code", Handler: &h.GuestCode{Source: dbStorage}},
		{Path: "/guest/create", Handler: &h.GuestCreate{Source: dbStorage, Random: random}},
		{Path: "/event/create", Handler: &h.EventCreate{Source: dbStorage}},
		{Path: "/event/edit/", Handler: &h.EventEdit{Source: dbStorage}},
		{Path: "/event/delete/", Handler: &h.EventDelete{Source: dbStorage}},
		{Path: "/event/list", Handler: &h.EventList{Source: dbStorage}},
		{Path: "/event/(.*)/table", Handler: &h.EventTableCreate{Source: dbStorage}},
		//Serve static files //todo switch from "/" to "/static/" route. need to refactor React paths
		{Path: "/", Handler: http.FileServer(http.Dir("./static/build"))},
	}.Prepare()
}

func main() {
	http.Handle("/", &h.Router{routes})

	log.Printf("Listen and serve with config: %#v", *appConfig)

	if err := http.ListenAndServe(":"+appConfig.Port, nil); err != nil {
		panic(err)
	}
}
