// package main
package test

import (
	"Gobang-2004A/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WebsocketTest(c *gin.Context) {
	fmt.Println("ws收到连接请求")
	ws, err := util.OpenWebSocket(c)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("ws接收消息遇到了问题")
			break
		}
		stringMessage := string(message)
		fmt.Println("ws收到信息,类型：" + strconv.Itoa(messageType) + "，内容：" + stringMessage)

		byteMessage := []byte(stringMessage)
		err = ws.WriteMessage(messageType, byteMessage)
		if err != nil {
			fmt.Println("ws发送消息遇到了问题")
			break
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建房间",
	})
}

// func main() {
// 	router := gin.Default()

// 	router.GET("/wstest", WebsocketTest)
// 	fmt.Println("localhost:8080")
// 	router.Run(":8080")
// }
