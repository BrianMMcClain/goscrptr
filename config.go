package main

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	WUKey string
}

func parseConfig(configPath string) Config {
	c := Config{}
	confJ, _ := ioutil.ReadFile(configPath)
	// TODO: Handle missing file/parse error
	json.Unmarshal(confJ, &c)
	return c
}