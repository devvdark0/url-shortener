package main

import (
	"flag"
	"fmt"
	"github.com/devvdark0/url-shortener/internal/config"
)

func main() {
	configPath := flag.String("config-path", "", "path to config file")
	flag.Parse()

	cfg := config.MustLoad(*configPath)
	//TODO: init logger

	//TODO: init storage

	//TODO: init router

	//TODO: run the server

}
