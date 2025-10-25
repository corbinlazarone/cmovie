package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes() http.Handler {
	router := httprouter.New()
	router.Handler(http.MethodGet, "/", http.HandlerFunc(health))

	return router
}
