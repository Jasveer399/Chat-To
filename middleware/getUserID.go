package middleware

import (
	"net/http"
)

func GetUserID(r *http.Request) (uint, bool) {
	userID := r.Context().Value(UserIDKey)
	if uid, ok := userID.(uint); ok {
		return uid, true
	}
	return 0, false
}
