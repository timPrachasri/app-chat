package socket

import (
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
	HandleMessages(c *gin.Context)
	ReadChatByRoomName(c *gin.Context)
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
	router.GET("messages.info", chatHandler.ReadChatByRoomName)
}

