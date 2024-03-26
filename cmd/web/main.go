package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	envFileNotLoaded := godotenv.Load()

	if envFileNotLoaded != nil {
		log.Fatal("Error loading .env")
	}

	addrFromEnv := os.Getenv("port")

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

	log.Printf("start serving on  %s and address from env is %s", *addr, addrFromEnv)
	err := http.ListenAndServe(addrFromEnv, mux)
	log.Fatal(err)
}
