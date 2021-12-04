package entity

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type Room struct {
	Roomid   string
	Players  [2]*Player
	Game     *Game
	RoomLock sync.Mutex
}

func NewRoom(player *Player) *Room {
	room := new(Room)
	room.Roomid = uuid.NewV4().String()
	room.Players[0] = player
	room.Game = nil
	return room
}

func (room *Room) EnterRoom(player *Player) error {
	room.RoomLock.Lock()
	defer room.RoomLock.Unlock()
	if room.Players[1] == nil {
		room.Players[1] = player
	} else {
		return errors.New("room is full")
	}
	return nil
}

func (room *Room) LeaveRoom(playnum int) {
	room.RoomLock.Lock()
	defer room.RoomLock.Unlock()
	if playnum == 0 {
		room.Players[0].WebSocket.Close()
		room.Players[0] = nil
		room.Players[1].WebSocket.Close()
		room.Players[1] = nil
	} else if playnum == 1 {
		room.Players[1].WebSocket.Close()
		room.Players[1] = nil
	} else {
		room.Players[0].WebSocket.Close()
		room.Players[0] = nil
		room.Players[1].WebSocket.Close()
		room.Players[1] = nil
	}
}

func (room *Room) SetGame(game *Game) {
	room.Game = game
}

func (room *Room) Lock() {
	room.RoomLock.Lock()
}

func (room *Room) Unlock() {
	room.RoomLock.Unlock()
}
