package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/ereminIvan/tablebooking/handler"
	"github.com/ereminIvan/tablebooking/service"

	fb "github.com/ereminIvan/firebase"
)

var dbStorage service.ISource
var random service.IRandom
var appConfig *Config
var routeList map[handler.IHandler]*regexp.Regexp

type Config struct {
	Port       string `json:"port"`
	FbDBSecret string `json:"fb_db_secret"`
	FbDBPath   string `json:"fb_db_path"`
}

func ReadConfig() (*Config, error) {
	c := &Config{}
	cf, err := ioutil.ReadFile("config.json")
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(cf, c)
	return c, err
}

func init() {
	c, err := ReadConfig()
	if err != nil {
		panic(err)
	}
	appConfig = c
	dbStorage = service.NewStorage(fb.NewDBClient(appConfig.FbDBPath, appConfig.FbDBSecret, false, nil))
	random = service.NewRand()

	routeList = handler.RouteList{
		{
			Path:    "/guest/code",
			Handler: &handler.GuestCode{Source: dbStorage},
		},
		{
			Path:    "/event/create",
			Handler: &handler.EventCreate{Source: dbStorage},
		},
		{
			Path:    "/event/list",
			Handler: &handler.EventList{Source: dbStorage},
		},
		{
			Path:    "/guest/create",
			Handler: &handler.GuestCreate{Source: dbStorage, Random: random},
		},
	}.Build()
}

func main() {
	http.Handle("/", &handler.Router{RouteList: routeList})

	log.Printf("Listen and serve with config: %#v", *appConfig)

	http.ListenAndServe(":"+appConfig.Port, nil)
}
