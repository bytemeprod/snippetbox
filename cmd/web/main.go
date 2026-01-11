package main

import (
	"log"
	"net/http"
)

const SERVER_ADDR = "localhost:8080"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /view", snippetView)
	mux.HandleFunc("POST /create", snippetCreate)

	log.Printf("Starting server (%s)\n", SERVER_ADDR)
	http.ListenAndServe(SERVER_ADDR, mux)
}
