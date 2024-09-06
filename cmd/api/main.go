package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"go-breeders-remote/configuration"
)

const port = ":8081"

type application struct {
	App    *configuration.Application // a singleton which is exported, so we can get to it from other modules.
	config appConfig                  // configuration information for the app.
}

// appConfig is a type embedded into the application type. It holds things that no other part of the
// app needs to know about.
type appConfig struct {
	useCache bool
	dsn      string
}

// main is the entry point for our app.
func main() {
	var config appConfig

	// Env or default value
	dsn := os.Getenv("DSN")

	// read command line parameters, if any, and set sensible defaults for development
	flag.StringVar(&config.dsn, "dsn", dsn, "DSN")
	flag.Parse()

	// get database
	db, err := initMySQLDB(config.dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := application{
		App:    configuration.New(db),
		config: config,
	}

	// create http server
	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	log.Println("*** Starting server on port", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
