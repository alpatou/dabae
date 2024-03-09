package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" messag
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from this app"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display snippet"))

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create snippet"))
}

func main() {

	// mux is like router synonym
	mux := http.NewServeMux()
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
