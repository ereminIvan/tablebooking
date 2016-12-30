package dto

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port       string `json:"port"`
	FbDBSecret string `json:"fb_db_secret"`
	FbDBPath   string `json:"fb_db_path"`
}

func (c *Config) Read(path string) error {
	var err error
	if cf, err := ioutil.ReadFile(path); err == nil {
		return json.Unmarshal(cf, c)
	}
	return err
}
