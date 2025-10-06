package main

import (
	"log"
	"net/http"

	"github.com/bytemeprod/snippetbox/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./frontend/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet/create", handlers.SnippetCreate)
	mux.HandleFunc("/snippet/view", handlers.SnippetView)

	log.Println("Server starting on :4000")

	http.ListenAndServe(":4000", mux)
}
