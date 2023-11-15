package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/saadi925/rssagregator/internal/routes"
)

func main() {
	loadEnv()
	port := os.Getenv("PORT")
	apiCfg := dbInit()
	go startScrapping(apiCfg.DB, 10, time.Minute)
	r := routes.SetupRoutes(apiCfg)
	fmt.Printf("listening on port is %v", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// loads env variables
func loadEnv() {
	envPath := ".env"
	godotenv.Load(envPath)
}
