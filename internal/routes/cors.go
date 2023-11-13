package routes

import (
	"net/http"

	"github.com/go-chi/cors"
)

// This will enable cors for  All Origins and it is just a basic template yo get started .
// you should update this as per your project. These are some basic Configurations
// You can Learn More about cors at http://www.google.com
func enableCors() func(http.Handler) http.Handler {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "application/json"},
	})
	return cors.Handler
}
