package websocket

import (
	"encoding/json"

	"github.com/Jasveer399/Chat-To/database"
	"github.com/Jasveer399/Chat-To/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	UserID uint
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
}

func (c *Client) ReadMessages() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, raw, err := c.Conn.ReadMessage()

		if err != nil {
			break
		}

		// Step 1: Parse JSON payload
		var payload struct {
			ReceiverID uint   `json:"receiver_id"`
			Content    string `json:"content"`
		}
		if err := json.Unmarshal(raw, &payload); err != nil {
			continue // skip bad payload
		}

		msg := &models.Message{
			SenderID:   c.UserID,
			ReceiverID: payload.ReceiverID,
			Content:    payload.Content,
		}

		if err := database.DB.Create(&msg).Error; err != nil {
			continue // optional: log or send error
		}

		c.Hub.Broadcast <- raw
	}
}

func (c *Client) WriteMessages() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}

		}

	}
}
