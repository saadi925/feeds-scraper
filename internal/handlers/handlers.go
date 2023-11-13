package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/saadi925/rssagregator/internal/database"
	"github.com/saadi925/rssagregator/internal/models"
	"github.com/saadi925/rssagregator/internal/utils"
)

type ApiConfig struct {
	DB *database.Queries
}

// here you can add the handler functions for the routes .
// we are using utils/json.go for handling error response and json responses.
func (apiCfg *ApiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	parameter := struct{ name string }{}
	if err := utils.ParseJSON(r, &parameter); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:      parameter.name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        uuid.New(),
	})
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	utils.RespondWithJSON(w, models.DbUserToUser(user), http.StatusCreated)

}
func (apiCfg *ApiConfig) GetUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, models.DbUserToUser(user), http.StatusOK)
}
func (apiCfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	var parameters *models.Feed
	if err := utils.ParseJSON(r, &parameters); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:      parameters.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
	}
	utils.RespondWithJSON(w, models.DBFeedToFeed(feed), http.StatusCreated)

}
func (apiCfg *ApiConfig) GetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
	}
	utils.RespondWithJSON(w, models.DBFeedsToFeeds(feeds), http.StatusOK)
}
