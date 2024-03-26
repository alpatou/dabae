package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "obvious")
	flag.Parse()

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

	log.Printf("start serving on  %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
