package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/saadi925/rssagregator/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func GetApiKeyFromHeader(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication information found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}

	return vals[1], nil
}
