package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ereminIvan/tablebooking/handler"
	"github.com/ereminIvan/tablebooking/service"

	fb "github.com/ereminIvan/firebase"
)

var dbStorage service.ISource
var random service.IRandom

type Config struct {
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
	dbStorage = service.NewStorage(fb.NewDBClient(c.FbDBPath, c.FbDBSecret, false, nil))
	random = service.NewRand()
}

func main() {
	http.Handle("/guest/code", &handler.GuestCode{Source: dbStorage})
	http.Handle("/event/create", &handler.EventCreate{Source: dbStorage})
	http.Handle("/event/list", &handler.EventList{Source: dbStorage})
	http.Handle("/guest/create", &handler.GuestCreate{Source: dbStorage, Random: random})

	http.ListenAndServe(":8090", nil)
}
