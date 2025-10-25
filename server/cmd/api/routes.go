package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func routes() http.Handler {
	router := httprouter.New()
	router.Handler(http.MethodGet, "/", http.HandlerFunc(health))

	std := alice.New(secureHeaders)

	return std.Then(router)
}
