package api

import (
	// "backend/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetToiletWebsocket(c *gin.Context) {
	w := c.Writer
	r := c.Request
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	count := 0

	for {
		t, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
		count += 1
		responseMessage := []byte(strconv.Itoa(count))
		conn.WriteMessage(t, responseMessage)
	}
}
