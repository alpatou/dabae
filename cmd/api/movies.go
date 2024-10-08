package main

import (
	"fmt"
	"net/http"
)

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintln(w, "Movie fetched with id %d\n", id)
}

func (app *application) createMoviehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created a movie")
}
