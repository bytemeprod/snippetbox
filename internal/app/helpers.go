package app

import (
	"net/http"
)

func (a *Application) serverError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (a *Application) clientError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func (a *Application) notFound(w http.ResponseWriter) {
	a.clientError(w, http.StatusNotFound)
}
