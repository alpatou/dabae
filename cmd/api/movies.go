package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintln(w, "Movie fetched with id %d\n", id)
}

func (app *application) createMoviehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created a movie")
}
