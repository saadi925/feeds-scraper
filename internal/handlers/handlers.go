package handlers

import (
	"errors"
	"fmt"
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

func TestHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, "routes are working , you are good to go", 200)
	return
}

// here you can add the handler functions for the routes .
// we are using utils/json.go for handling error response and json responses.
func (apiCfg *ApiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	parameter := struct {
		Name string `json:"name"`
	}{}
	if err := utils.ParseJSON(r, &parameter); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	if len(parameter.Name) > 4 {
		utils.RespondWithError(w, errors.New("name is too short to be valid"), http.StatusBadRequest)
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:      parameter.Name,
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
	return
}
func (apiCfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	parameters := struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{}
	if err := utils.ParseJSON(r, &parameters); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:      parameters.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		Url:       parameters.Url,
		ID:        uuid.New(),
	})
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, models.DBFeedToFeed(feed), http.StatusCreated)

}
func (apiCfg *ApiConfig) GetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, models.DBFeedsToFeeds(feeds), http.StatusOK)
}

func (apiCfg *ApiConfig) CreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	parameter := struct {
		FeedID string `json:"feed_id"`
	}{}

	if err := utils.ParseJSON(r, &parameter); err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	uuidFeedID, err := uuid.Parse(parameter.FeedID)
	if err != nil {
		utils.RespondWithError(w, errors.New("error parsing feed id"), http.StatusInternalServerError)
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		FeedID:    uuidFeedID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		ID:        uuid.New(),
	})
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, models.DBFeedFollowToFeedFollow(feedFollow), http.StatusCreated)

}

func (apiCfg *ApiConfig) GetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, models.DBFeedsFollowToFeedsFollow(feeds), http.StatusOK)
}
func (apiCfg *ApiConfig) DeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	parameter := struct {
		FeedID string `json:"feed_id"`
	}{}
	uuidFeedID, err := uuid.Parse(parameter.FeedID)
	if err != nil {
		utils.RespondWithError(w, errors.New("error parsing feed id"), http.StatusInternalServerError)
		return
	}
	if err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     uuidFeedID,
		UserID: user.ID,
	}); err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}
	message := struct {
		Message string `json:"message"`
		FeedID  string `json:"feed_id"`
	}{
		Message: fmt.Sprintf("the feed has been deleted sucessfully with the feed id %v ", parameter.FeedID),
	}
	utils.RespondWithJSON(w, message, http.StatusOK)
}
