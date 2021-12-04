package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func OpenWebSocket(c *gin.Context) (ws *websocket.Conn, err error) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err = upGrader.Upgrade(c.Writer, c.Request, nil)
	return
}
