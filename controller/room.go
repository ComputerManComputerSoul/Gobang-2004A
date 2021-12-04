package controller

import (
	"Gobang-2004A/entity"
	"Gobang-2004A/service"
	"Gobang-2004A/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	fmt.Println("收到创建房间请求")
	ws, err := util.OpenWebSocket(c)
	if err != nil {
		return
	}
	defer ws.Close()
	player := entity.NewPlayer(ws)
	var room *entity.Room
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			if (*room).Roomid != "" {
				(*room).LeaveRoom(0)
			}
			break
		}
		service.HandleMessage(messageType, message, &room, player, 0)
	}
}

func EnterRoom(c *gin.Context) {
	fmt.Println("收到加入房间请求")
	ws, err := util.OpenWebSocket(c)
	if err != nil {
		return
	}
	defer ws.Close()
	player := entity.NewPlayer(ws)
	var room *entity.Room
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			if (*room).Roomid != "" {
				(*room).LeaveRoom(1)
			}
			break
		}
		service.HandleMessage(messageType, message, &room, player, 1)
	}
}
