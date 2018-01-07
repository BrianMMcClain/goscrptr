package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SpotifyStatus struct {
	Playing    bool   `json:"playing"`
	DeviceName string `json:"deviceName"`
	Track      string `json:"track"`
	Artist     string `json:"artist"`
	Shuffle    bool   `json:"shuffle"`
}

func spotifyStatus(conf Config) SpotifyStatus {
	spotifyServiceIP := conf.SpotifyServiceIP
	spotifyServicePort := conf.SpotifyServicePort
	statusURI := fmt.Sprintf("http://%s:%d/status", spotifyServiceIP, spotifyServicePort)

	resp, err := http.Get(statusURI)
	if err != nil {
		fmt.Println("Error getting current Spotify status")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// Parse the JSON
	s := SpotifyStatus{}
	json.Unmarshal(body, &s)

	return s
}
