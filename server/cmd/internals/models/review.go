package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

type ReviewModel struct {
	DB *pgxpool.Pool
}
