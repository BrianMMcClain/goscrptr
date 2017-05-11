package main

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	WUKey string `json:"WUKey"`
	SpotifyServiceIP string `json:"SpotifyServiceIP"`
	SpotifyServicePort string `json:"SpotifyServicePort"`
}

func parseConfig(configPath string) Config {
	c := Config{}
	confJ, _ := ioutil.ReadFile(configPath)
	// TODO: Handle missing file/parse error
	json.Unmarshal(confJ, &c)
	return c
}
