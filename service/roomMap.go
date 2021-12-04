package service

import (
	"Gobang-2004A/entity"
	"sync"
)

var (
	RoomMap sync.Map
	lock    sync.Mutex
)

func StoreRoom(room *entity.Room) {
	RoomMap.Store(room.Roomid, room)
}

func DeleteRoom(roomid string) {
	lock.Lock()
	defer lock.Unlock()
	RoomMap.Delete(roomid)
}

func LoadRoom(roomid string) (room *entity.Room, ok bool) {
	_room, ok := RoomMap.Load(roomid)
	room = _room.(*entity.Room)
	return
}
