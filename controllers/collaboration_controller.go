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

var hub = models.CollaborationHub{
	Clients:    make(map[*websocket.Conn]bool),
	Broadcast:  make(chan []byte),
	Register:   make(chan *websocket.Conn),
	Unregister: make(chan *websocket.Conn),
}

func WebsocketHanler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println("Error in upgrading connection:", err)
		return
	}

	hub.Register <- conn

	go handleMessages(conn)
}

func handleMessages(conn *websocket.Conn) {
	defer func() {
		hub.Unregister <- conn
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("error in reading message:", err)
			break
		}

		hub.Broadcast <- message
	}
}

func RunHub() {
	for {
		select {
		case conn := <-hub.Register:
			hub.Mutex.Lock()
			hub.Clients[conn] = true
			hub.Mutex.Unlock()
			fmt.Println("client connected")

		case conn := <-hub.Unregister:
			hub.Mutex.Lock()
			delete(hub.Clients, conn)
			hub.Mutex.Unlock()
			fmt.Println("client disconnected")

		case message := <-hub.Broadcast:
			hub.Mutex.Lock()
			for client := range hub.Clients {
				if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
					fmt.Println("error sending message:", err)
					client.Close()
					delete(hub.Clients, client)
				}
			}
			hub.Mutex.Unlock()
		}
	}

}
