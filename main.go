package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/saadi925/rssagregator/internal/database"
	"github.com/saadi925/rssagregator/internal/handlers"
	"github.com/saadi925/rssagregator/internal/routes"
)

func main() {
	loadEnv()
	port := os.Getenv("PORT")

	apiCfg := dbInit()

	r := routes.SetupRoutes(apiCfg)
	fmt.Printf("listening on port is %v", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// loads env variables
func loadEnv() {
	envPath := ".env"
	godotenv.Load(envPath)
}

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
