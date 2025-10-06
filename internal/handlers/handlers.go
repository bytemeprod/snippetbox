package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./frontend/html/base.html",
		"./frontend/html/pages/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Failed to parse templates: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := ts.ExecuteTemplate(w, "base", nil); err != nil {
		log.Println("Failed to execute template: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Creating snippet..."))
}

func SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a snippet with id %d", id)
}
