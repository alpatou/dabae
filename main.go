package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from this app"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display snippet"))

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create snippet"))
}

func main() {

	// mux is like router synonym
	mux := http.NewServeMux()
	// subtree path, end with / (and starts on this case)
	mux.HandleFunc("/", home)
	// fixed path
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("start serving on  :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
