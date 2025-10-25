package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "cmovie server is running!")
}

func AllReviews() {

}

func SingleReview() {

}
