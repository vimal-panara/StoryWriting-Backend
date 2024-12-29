package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type CollaborationHub struct {
	Clients    map[*websocket.Conn]bool
	Broadcast  chan []byte
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
	Mutex      sync.Mutex
}
