package entity

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	NickName  string
	WebSocket *websocket.Conn
}

func NewPlayer(WebSocket *websocket.Conn) *Player {
	return &Player{
		NickName:  "",
		WebSocket: WebSocket,
	}
}
