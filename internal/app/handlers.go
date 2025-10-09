package app

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/bytemeprod/snippetbox/pkg/prettylog"
)

func (a *Application) NewHome(plog *slog.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" || r.Method != http.MethodGet {
			a.notFound(w)
			return
		}

		files := []string{
			"./frontend/html/base.html",
			"./frontend/html/pages/index.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			plog.Error("Failed to parse templates: ", prettylog.Error(err))
			a.serverError(w, err)
			return
		}

		if err := ts.ExecuteTemplate(w, "base", nil); err != nil {
			plog.Error("Failed to execute template: ", prettylog.Error(err))
			a.serverError(w, err)
			return
		}
	}
}

func (a *Application) NewSnippetCreate(plog *slog.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			a.clientError(w, http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte("Creating snippet..."))
	}
}

func (a *Application) NewSnippetView(plog *slog.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			a.notFound(w)
			return
		}

		fmt.Fprintf(w, "Display a snippet with id %d", id)
	}
}
