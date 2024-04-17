package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	// fixed path
	// longer matches are served from priority
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}