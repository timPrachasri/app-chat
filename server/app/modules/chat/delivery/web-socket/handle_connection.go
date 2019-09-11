package socket

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	entities "github.com/timPrachasri/app-chat/server/app/entities"
	utils "github.com/timPrachasri/app-chat/server/utils"
)

// ReadChat is a  function for reading chats
func(handler chatHTTPHandler) ReadChatByRoomName(c *gin.Context) {
	roonName := c.Query("roomName")
	msgs, err  := handler.chatUsecase.ReadChat(c, roonName)
	if err != nil {
		c.JSON(http.StatusOK, utils.NewErrorResponse(err.Error()))
		return
	}
	response := map[string]interface{}{
		"messages": msgs,
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse(response))
}

// HandleMessages is a handler function for sending each messages to client
func(handler chatHTTPHandler) HandleMessages(c *gin.Context) {
	for { 
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		log.Printf("%+v msg sending", msg)
		// Send it out to every client that is currently connected
		for client := range Clients {
			msgs, err  := handler.chatUsecase.ReadChat(c, msg.RoomName)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(Clients, client)
			}
			if err := client.WriteJSON(msgs); err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(Clients, client)
			}
		}
	}
}

func(handler chatHTTPHandler) HandleConnections(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
	}
	Clients[ws] = true

	go handler.HandleMessages(c)

	for {
		var msg entities.Message
		// Read in a new message as JSON and map it to a Message object
		log.Printf("%+v msg comin", msg)
		if err := ws.ReadJSON(&msg); err != nil {
			log.Printf("error: %v", err)
				delete(Clients, ws)
				break
		}
		handler.chatUsecase.AddChat(c, msg)
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}