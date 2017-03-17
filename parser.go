package main

import (
	"regexp"
	"strings"
	"time"
	"strconv"
	"fmt"
)

func parse(formatString string, config Config) string {
	
	// Get all format tags
	r, _ := regexp.Compile("{{([\\w\\s]*)}}")
	retString := formatString

	for _,match := range r.FindAllStringSubmatch(formatString, -1) {
		// Get method name
		var fSplit = strings.Split(match[1], " ")
		var replaceString = ""
		switch fSplit[0] {
			case "timeSpeech":
				replaceString = timeSpeech()
			case "city":
				replaceString = currentWeatherConditions(fSplit[1], config).Location.City
			case "temperature":
				replaceString = strconv.FormatFloat(currentWeatherConditions(fSplit[1], config).TempF, 'f', 1, 64)
			case "weatherCondition":
				replaceString = currentWeatherConditions(fSplit[1], config).Weather
			default:
				fmt.Printf("No method for %s\n", fSplit[0])
		}

		retString = strings.Replace(retString, match[0], replaceString, -1)
	}

	return retString
}

func timeSpeech() string {
	var now = time.Now()
	if (now.Minute() == 0) {
		return now.Format("3 PM")
	} else if (now.Minute() < 10) {
		return now.Format("3 oh 4 PM")
	} else {
		return now.Format("3 04 PM")
	}
}