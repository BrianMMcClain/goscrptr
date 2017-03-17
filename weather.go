package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Conditions struct {
	Observation Observation `json:"current_observation"`
}

type Observation struct {
	Location Location `json:"display_location"`
	TempF float64 `json:"temp_f"`
	TempC float64 `json:"temp_c"`
	Weather string `json:"weather"`
	WindDir string `json:"wind_dir"`
	WindMPH float64 `json:"wind+mph"`
}

type Location struct {
	City string `json:"city"`
	State string `json:"state_name"`
}

func currentWeatherConditions(zip string, conf Config) Observation {
	key := conf.WUKey
	currentConditionsFormat := "http://api.wunderground.com/api/%s/conditions/q/%s.json"

	// Get the weather from Weather Underground
	callURL := fmt.Sprintf(currentConditionsFormat, key, zip)
	resp, err := http.Get(callURL)
	if (err != nil) {
		fmt.Println("Error getting current weather conditions")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	
	// Parse the JSON
	c := Conditions{}
	json.Unmarshal(body, &c)

	return c.Observation
}