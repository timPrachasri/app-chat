package socket

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	entities "github.com/timPrachasri/app-chat/server/app/entities"
)

func(handler chatHTTPHandler) HandleConnections(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
	}
	Clients[ws] = true
	for {
		var msg entities.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		log.Printf("%+v msg comin", msg)
		if err != nil {
				log.Printf("error: %v", err)
				delete(Clients, ws)
				break
		}
		
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}