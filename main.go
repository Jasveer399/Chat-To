package main

import (
	"log"
	"net/http"

	controllers "github.com/Jasveer399/Chat-To/controllers/auth"
	"github.com/Jasveer399/Chat-To/database"
	"github.com/Jasveer399/Chat-To/middleware"
	"github.com/Jasveer399/Chat-To/websocket"
)

func main() {
	database.InitDatabase()
	hub := websocket.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", middleware.JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	}))
	http.HandleFunc("/register", controllers.RegisterHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/get-user", middleware.JWTMiddleware(controllers.GetAllUsersHandler))

	log.Println("Server started on :3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
