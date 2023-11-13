package handlers

import (
	"net/http"

	"github.com/saadi925/rssagregator/internal/auth"
	"github.com/saadi925/rssagregator/internal/database"
	"github.com/saadi925/rssagregator/internal/utils"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) ApiAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKeyFromHeader(r.Header)
		if err != nil {
			utils.RespondWithError(w, err, http.StatusUnauthorized)
		}
		user, err := apiCfg.DB.GetUsersByApiKey(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, err, http.StatusInternalServerError)
		}
		handler(w, r, user)
	}
}
