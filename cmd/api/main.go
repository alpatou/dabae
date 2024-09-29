package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"log"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	// declare config struct instance
	var cfg config

	// read from command line args, if no default vals
	// port, env and then partse
	flag.IntVar(&cfg.port, "port", 4000, "server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|test|staging|production)")
	flag.Parse()

	// define logger

	logger := log.New(os.Stdout, "Log", log.Ldate|log.Ltime)

	// pointer instance of application
	app := &application{
		config: cfg,
		logger: logger,
	}

	// server mux and health check EP

	// POST movies
	// GET movies/:id

	// launch server , prin error if any

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("start  %s server at %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
