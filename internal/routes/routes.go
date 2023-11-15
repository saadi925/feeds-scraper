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
	r.Get("/test", handlers.TestHandler)
	userRoutes(r, apiCfg)
	feedRoutes(r, apiCfg)
	feedFollowRoutes(r, apiCfg)
	// Add Routes here .
	return r
}

func userRoutes(r chi.Router, apiCfg handlers.ApiConfig) {
	r.Route("/user", func(r chi.Router) {
		r.Post("/", apiCfg.CreateUserHandler)
		r.Get("/", apiCfg.ApiAuth(apiCfg.GetUserHandler))
	})
}

func feedRoutes(r chi.Router, apiCfg handlers.ApiConfig) {
	r.Route("/feeds", func(r chi.Router) {
		r.Post("/", apiCfg.ApiAuth(apiCfg.CreateFeed))
		r.Get("/", apiCfg.ApiAuth(apiCfg.GetFeeds))
	})
}
func feedFollowRoutes(r chi.Router, apiCfg handlers.ApiConfig) {
	r.Route("/feed_follows", func(r chi.Router) {
		r.Get("/", apiCfg.ApiAuth(apiCfg.GetFeedFollows))
		r.Post("/", apiCfg.ApiAuth(apiCfg.CreateFeedFollow))
		r.Delete("/", apiCfg.ApiAuth(apiCfg.DeleteFeedFollow))
	})
}
