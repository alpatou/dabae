package main

import (
	"flag"
	"fmt"
	"os"

	"log"
)

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
	flag.IntVar(*cfg.port, "port", 4000, "server port")
	flag.StringVar(*cfg.env, "env", "development", "Environment (development|test|staging|production)")
	flag.Parse()

	// define logger

	logger := log.New(os.Stdout, "Log", log.Ldate|log.Ltime)

	// pointer instance of application
	application & application{
		config: cfg,
		logger: logger,
	}

	// server mux and health check EP

	// launch server , prin error if any

}
