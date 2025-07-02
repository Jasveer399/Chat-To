package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jasveer399/web-service-gin/common"
	"github.com/Jasveer399/web-service-gin/database"
	"github.com/Jasveer399/web-service-gin/middleware"
	"github.com/Jasveer399/web-service-gin/models"
	"github.com/Jasveer399/web-service-gin/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("DBUI28BHJPWU0298VN3I230JWLD982NDWO029")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MessageResponse struct {
	ID         uint   `json:"id"`
	SenderID   uint   `json:"sender_id"`
	ReceiverID uint   `json:"receiver_id"`
	Content    string `json:"content"`
}

type PublicUser struct {
	ID       uint              `json:"id"`
	Username string            `json:"username"`
	Messages []MessageResponse `json:"messages"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload", err, nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error hashing password", err, nil)
		return
	}

	user := models.User{Username: creds.Username, Password: string(hashedPassword)}
	db := database.DB

	// Check if the user already exists
	var existingUser models.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		utils.SendError(w, http.StatusConflict, "User already exists", nil, nil)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error creating user", err, nil)
		return
	}

	utils.SendResponse(w, http.StatusCreated, "User registered successfully", nil, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload", err, nil)
		return
	}

	var user models.User
	db := database.DB
	if err := db.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		utils.SendError(w, http.StatusUnauthorized, "User not found", err, nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Invalid password", err, nil)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &common.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error generating token", err, nil)
		return
	}

	utils.SendResponse(w, http.StatusOK, "Login successful", map[string]string{"token": tokenString}, nil)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	senderID, ok := middleware.GetUserID(r)

	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized", nil, nil)
		return
	}
	db := database.DB
	var users []models.User
	if err := db.
		Where("id != ?", senderID).
		Preload("MessagesSent", "receiver_id = ?", senderID).
		Preload("MessagesReceived", "sender_id = ?", senderID).
		Find(&users).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error fetching users", err, nil)
		return
	}

	utils.SendResponse(w, http.StatusOK, "Users fetched successfully", users, nil)
}
