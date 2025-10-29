package main

import (
	"fmt"
	"net/http"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "cmovie server is running!")
}

// NOTE: public
func (app *application) AllReviews() {
}

// NOTE:  public
func (app *application) SingleReview(w http.ResponseWriter, r *http.Request) {
}

// NOTE: admin
func (app *application) CreateReview() {
}

// NOTE: admin
func (app *application) EditReview() {
}

// NOTE: admin
func (app *application) DeleteReview() {
}
