package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/saadi925/rssagregator/internal/database"
	"github.com/saadi925/rssagregator/internal/handlers"
)

func dbInit() handlers.ApiConfig {
	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DB_URL not is in env")
	}

	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("error opening postgres ", err)
	}
	queries := database.New(conn)
	apiCfg := handlers.ApiConfig{
		DB: queries,
	}

	return apiCfg

}
