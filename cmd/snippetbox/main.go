package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bytemeprod/snippetbox/internal/handlers"
)

type config struct {
	addr      string
	staticDir string
}

var cfg config

// Parsing flags
func init() {
	flag.StringVar(&cfg.addr, "addr", ":4040", "HTTP network address to start server")
	flag.StringVar(&cfg.staticDir, "staticDir", "./frontend/static", "Path to static files")
	flag.Parse()
}

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./frontend/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet/create", handlers.SnippetCreate)
	mux.HandleFunc("/snippet/view", handlers.SnippetView)

	log.Printf("Server starting on %s\n", cfg.addr)

	server := http.Server{
		Addr:    cfg.addr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
