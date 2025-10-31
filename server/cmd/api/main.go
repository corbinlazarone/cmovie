package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/corbinlazarone/cmovie/cmd/internals/migrations"
	"github.com/corbinlazarone/cmovie/cmd/internals/models"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib" // to load the pgx driver for database/sql
)

type application struct {
	reviews *models.ReviewModel
}

func main() {

	// db url is auto injected by docker-compose
	dbUrl := os.Getenv("DB_CONN")
	if dbUrl == "" {
		log.Fatal("DB_CONN env var not set")
	}

	dbPool, err := initDB(dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer dbPool.Close()

	// create sql db connection for goose migrations
	sqlDB, err := sql.Open("pgx", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	err = migrations.RunMigrations(sqlDB)
	if err != nil {
		log.Fatal("Failed to load migrations", err)
	}

	// so we can share our CRUD operations across our app
	app := &application{
		reviews: &models.ReviewModel{
			DB: dbPool,
		},
	}

	srv := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	log.Println("Server running on port 4000...")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func initDB(dataSource string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dataSource)

	if err != nil {
		return nil, err
	}

	// pinging the connection to see if it was successful
	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
