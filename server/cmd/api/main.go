package main

import (
	"log"
	"net/http"
	"os"

	"github.com/corbinlazarone/cmovie/cmd/internals/models"
)

type application struct {
	reviews *models.ReviewModel
}

func main() {

	// db url is auto injected by docker-compose
	url := os.Getenv("DB_CONN")
	if url == "" {
		log.Fatal("DB_CONN env var not set")
	}

	// so we can share our CRUD operations across our app
	app := &application{
		reviews: &models.ReviewModel{
			// TODO: db connection goes here
		},
	}

	srv := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	log.Println("Server running on port 4000...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
