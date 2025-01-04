package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

// RoomConnection represents a user's connection to a specific story room
type RoomConnection struct {
	StoryID string
	Conn    *websocket.Conn
}

// CollaborationMessage represents a message being sent in a room
type CollaborationMessage struct {
	StoryID string `json:"storyId"`
	Content string `json:"content"`
	UserID  string `json:"userId"`
	Type    string `json:"type"` // e.g., "edit", "cursor", etc.
}

type CollaborationHub struct {
	Rooms      map[string]map[*websocket.Conn]bool // Rooms mapped by story ID
	Broadcast  chan CollaborationMessage           // Channel for broadcasting messages
	Register   chan RoomConnection                 // Channel for new connections
	Unregister chan RoomConnection                 // Channel for disconnecting
	Mutex      sync.Mutex                          // Mutex for thread-safe operations
}
