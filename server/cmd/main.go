package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "cmovie server is running!")
}

func main() {
	router := httprouter.New()
	router.GET("/", health)

	log.Println("Server running on port 4000...")
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		log.Fatal(err)
	}
}
