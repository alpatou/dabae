package main

import (
	"log"
	"net/http"
)

func main() {
	// mux is like router synonym
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// subtree path, end with / (and starts on this case)
	mux.HandleFunc("/", home)
	// fixed path
	// longer matches are served from priority
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("start serving on  :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
