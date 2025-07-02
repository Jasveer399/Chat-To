package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/Jasveer399/web-service-gin/common"
	"github.com/Jasveer399/web-service-gin/database"
	"github.com/Jasveer399/web-service-gin/models"
	"github.com/Jasveer399/web-service-gin/utils"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"

var jwtSecret = []byte("DBUI28BHJPWU0298VN3I230JWLD982NDWO029")

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		if tokenString == "" {
			utils.SendError(w, http.StatusUnauthorized, "Missing or invalid token", nil, nil)
			log.Printf("Missing or invalid token in request: %s", r.URL.Path, r.Method, tokenString)
			return
		}

		claims := &common.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			utils.SendError(w, http.StatusUnauthorized, "Invalid token", err, nil)
			return
		}

		var user models.User
		if err := database.DB.Where("username = ?", claims.Username).First(&user).Error; err != nil {
			utils.SendError(w, http.StatusUnauthorized, "User not found", err, nil)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, user.ID)
		next(w, r.WithContext(ctx))
	}
}

func extractToken(r *http.Request) string {
	// Try query param first (for WebSocket)
	if token := r.URL.Query().Get("token"); token != "" {
		return token
	}
	// Then check Authorization header
	authHeader := r.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return ""
}
