package main

import (
	"flag"
	"fmt"
)

var configPath string
var startServer bool
var evalString string

func main() {
	flag.StringVar(&configPath, "config", "config.json", "Path to scrptr config")
	flag.BoolVar(&startServer, "server", false, "Enable REST API")
	flag.StringVar(&evalString, "eval", "", "String to evaluate")
	flag.Parse()

	conf := parseConfig(configPath)

	if len(evalString) > 0 {
		fmt.Printf("%s\n", parse(evalString, conf))
	}

	if startServer {
		fmt.Println("Starting HTTP API")
	}
}
