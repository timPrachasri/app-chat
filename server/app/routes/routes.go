package routes

import (
	"github.com/gin-gonic/gin"
	chatWebsocket "github.com/timPrachasri/app-chat/server/app/modules/chat/delivery/web-socket"
	container  "github.com/timPrachasri/app-chat/server/app/interfaces/containers"
)

// ApplyRoutes is a function for grouping and applying api routes
func ApplyRoutes(r *gin.Engine) {
	chatContainer := container.NewChatContainer()
	chatWebsocket.ApplyRoutes(r, chatWebsocket.NewChatHTTPHandler(chatContainer.BootUsecase()))
}
