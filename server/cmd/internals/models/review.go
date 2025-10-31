package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Review struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Image     string    `json:"image_url"`
	Summary   string    `json:"summary"`
	Rating    int       `json:"rating"`    // will be rated 1-5
	AuthorID  string    `json:"author_id"` // will be defaulted to Corbin for now
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReviewModel struct {
	DB *pgxpool.Pool
}
