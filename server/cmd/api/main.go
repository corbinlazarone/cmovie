package main

import (
	"log"
	"net/http"
)

func main() {

	srv := &http.Server{
		Addr:    ":4000",
		Handler: routes(),
	}

	log.Println("Server running on port 4000...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
