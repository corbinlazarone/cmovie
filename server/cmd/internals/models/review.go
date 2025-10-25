package models

import "time"

type Review struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	Summary   string    `json:"summary"`
	Rating    int       `json:"rating"` // will be rated 1-5
	Author    string    `json:"author"` // will be defaulted to Corbin for now
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NOTE: will wrap a db command so we can communicate with the db table
// type ReviewModel struct {
// }

// NOTE: crud operations will be here
