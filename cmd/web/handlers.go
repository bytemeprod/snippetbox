package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Failed to parse template files: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := ts.ExecuteTemplate(w, "base", nil); err != nil {
		log.Printf("Failed to execute template: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write([]byte(fmt.Sprintf("view snippet %d", id)))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create snippet"))
}
