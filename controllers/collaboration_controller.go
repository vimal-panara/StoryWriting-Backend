package controllers

import (
	"fmt"
	"net/http"
	"story-plateform/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Initialize the Collaboration Hub
var hub = models.CollaborationHub{
	Rooms:      make(map[string]map[*websocket.Conn]bool),
	Broadcast:  make(chan models.CollaborationMessage),
	Register:   make(chan models.RoomConnection),
	Unregister: make(chan models.RoomConnection),
}

// handleMessages listens for messages from a specific connection
func handleMessages(storyID string, conn *websocket.Conn) {
	defer func() {
		// Unregister the connection when it is closed
		hub.Unregister <- models.RoomConnection{
			StoryID: storyID,
			Conn:    conn,
		}
		conn.Close()
	}()

	for {
		var message models.CollaborationMessage
		err := conn.ReadJSON(&message)

		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		// Set the story ID in the message and send it to the broadcast channel
		message.StoryID = storyID
		hub.Broadcast <- message
	}
}

// WebsocketHandler handles WebSocket connections for collaboration
func WebsocketHandler(c *gin.Context) {
	// Get the story ID from the URL parameter
	storyID := c.Param("storyId")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println("Error in upgrading connection:", err)
		return
	}

	// Register the connection to the specific story room
	hub.Register <- models.RoomConnection{
		StoryID: storyID,
		Conn:    conn,
	}

	// Start listening for messages from the client
	go handleMessages(storyID, conn)
}

// RunHub runs the main collaboration hub logic
func RunHub() {
	for {
		select {
		case connection := <-hub.Register:
			hub.Mutex.Lock()
			// Create a new room if it doesn't exist
			if _, exists := hub.Rooms[connection.StoryID]; !exists {
				hub.Rooms[connection.StoryID] = make(map[*websocket.Conn]bool)
			}
			hub.Rooms[connection.StoryID][connection.Conn] = true
			hub.Mutex.Unlock()
			fmt.Printf("Client connected to story %s\n", connection.StoryID)

		case connection := <-hub.Unregister:
			hub.Mutex.Lock()
			// Remove the connection from the room
			if clients, exists := hub.Rooms[connection.StoryID]; exists {
				delete(clients, connection.Conn)
				// Delete the room if it is empty
				if len(clients) == 0 {
					delete(hub.Rooms, connection.StoryID)
				}
			}
			hub.Mutex.Unlock()
			fmt.Printf("Client disconnected from story %s\n", connection.StoryID)

		case message := <-hub.Broadcast:
			hub.Mutex.Lock()
			// Broadcast the message to all clients in the room
			if clients, exists := hub.Rooms[message.StoryID]; exists {
				for conn := range clients {
					err := conn.WriteJSON(message)
					if err != nil {
						fmt.Println("Error sending message:", err)
						conn.Close()
						delete(clients, conn)
					}
				}
			}
			hub.Mutex.Unlock()
		}
	}
}
