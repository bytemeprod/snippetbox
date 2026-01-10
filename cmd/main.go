package main

import (
	"log"
	"net/http"
)

const SERVER_ADDR = "localhost:8080"

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("home"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("view snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create snippet"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /view", snippetView)
	mux.HandleFunc("POST /create", snippetCreate)

	log.Printf("Starting server (%s)\n", SERVER_ADDR)
	http.ListenAndServe(SERVER_ADDR, mux)
}
