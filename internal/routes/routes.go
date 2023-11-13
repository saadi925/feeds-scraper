package routes

import (
	"github.com/go-chi/chi"
	"github.com/saadi925/rssagregator/internal/handlers"
)

// by default cors are enabled for all origins , you should manually update cors.
// Setup Routes uses Chi Router, You can add your middlewares here. This is a starting template
// feel free to customize it .
func SetupRoutes(apiCfg handlers.ApiConfig) *chi.Mux {
	corsHandler := enableCors()
	r := chi.NewRouter()
	r.Use(corsHandler)
	userRoutes(r, apiCfg)
	// Add Routes here .
	//
	// **    Can use Middlewares too . ** //
	return r
}

func userRoutes(r chi.Router, apiCfg handlers.ApiConfig) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/", apiCfg.ApiAuth(apiCfg.GetUserHandler))
		r.Post("/", apiCfg.CreateUserHandler)
	})
}
func feedRoutes(r chi.Router, apiCfg handlers.ApiConfig) {
	r.Route("/feeds", func(r chi.Router) {
		r.Post("/", apiCfg.ApiAuth(apiCfg.CreateFeed))
		r.Post("/", apiCfg.ApiAuth(apiCfg.GetFeeds))
	})
}
