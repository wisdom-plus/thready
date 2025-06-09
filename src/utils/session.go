package utils

import (
	"net/http"
	"strconv"
)

func GetCurrentUserID(r *http.Request) (int, error) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		return 0, err
	}
	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
