package socket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	chat "github.com/timPrachasri/app-chat/server/app/modules/chat"
	entities "github.com/timPrachasri/app-chat/server/app/entities"
)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Clients represented Connected clients
var Clients = make(map[*websocket.Conn]bool)

var broadcast = make(chan entities.Message)

//ChatHTTPHandler represents a httphandler for chat module
type ChatHTTPHandler interface {
	HandleConnections(*gin.Context)
}

type chatHTTPHandler struct {
	chatUsecase chat.Usecase
}

// NewChatHTTPHandler is a function for creating client http controller instance
func NewChatHTTPHandler(chatUsecase chat.Usecase) ChatHTTPHandler {
	return &chatHTTPHandler{
		chatUsecase: chatUsecase,
	}
}

// ApplyRoutes is a function for apply any function to route
func ApplyRoutes(router *gin.Engine, chatHandler ChatHTTPHandler) {
	router.GET("/ws", chatHandler.HandleConnections)
}

// HandleMessages is a handler function for sending each messages to client
func HandleMessages() {
	for { 
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		log.Printf("%+v msg sending", msg)
		// Send it out to every client that is currently connected
		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(Clients, client)
			}
		}
	}
}
