package app

import (
	"log/slog"
	"net/http"
)

type Application struct {
	logger *slog.Logger
	config Config
}

func NewApplication(logger *slog.Logger, config Config) *Application {
	return &Application{
		logger: logger,
		config: config,
	}
}

func (a *Application) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(a.config.StaticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", a.NewHome(a.logger))
	mux.HandleFunc("/snippet/create", a.NewSnippetCreate(a.logger))
	mux.HandleFunc("/snippet/view", a.NewSnippetView(a.logger))

	return mux
}
