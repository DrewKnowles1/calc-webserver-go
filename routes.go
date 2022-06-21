package main

import (
	"fmt"
	"net/http"
)

type application struct {
}

func startAPI() {
	fmt.Println("Starting Webserver...")

	mux := http.NewServeMux()
	app := application{}

	mux.HandleFunc("/calc", app.calculate)
	http.ListenAndServe(":8080", mux)

}
