package main

import (
	"fmt"
	"net/http"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "cmovie server is running!")
}

func (app *application) AllReviews() {

}

func (app *application) SingleReview() {

}
