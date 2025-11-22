package main

import (
	"net/http"
	"os"
)

func main() {
	//TODO: init config

	//TODO: init logger

	//TODO: init storage

	if err := http.ListenAndServe(":80", nil); err != nil {
		os.Exit(1)
	}
}
