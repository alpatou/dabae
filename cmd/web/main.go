package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	envFileNotLoaded := godotenv.Load()

	if envFileNotLoaded != nil {
		log.Fatal("Error loading .env")
	}

	// addr := flag.String("addr", ":4000", "obvious")

	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "Port")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errLog,
		infoLog:  infolog,
	}

	// mux is like router synonym
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// subtree path, end with / (and starts on this case)
	mux.HandleFunc("/", app.home)
	// fixed path
	// longer matches are served from priority
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infolog.Printf("start serving on  %s", cfg.addr)
	err := srv.ListenAndServe()
	// log to a file someday
	errLog.Fatal(err)
}
