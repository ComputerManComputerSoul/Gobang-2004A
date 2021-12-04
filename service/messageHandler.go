package service

import (
	"Gobang-2004A/entity"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func HandleMessage(messageType int, originMessage []byte, room **entity.Room, player *entity.Player, playerNum int) error {
	message := entity.Message{}
	err := json.Unmarshal(originMessage, &message)
	fmt.Println("收到了" + player.NickName + "消息")
	fmt.Println(message)
	if err != nil {
		return err
	}
	if message.Sender == "client" {
		switch message.Name {
		case "createRoom":
			player.NickName = message.Content[0]
			*room = entity.NewRoom(player)
			StoreRoom(*room)
			SendMessage(player, entity.NewMessage("roomNumber", (*room).Roomid))
			return nil
		case "enterRoom":
			player.NickName = message.Content[0]
			_room, ok := LoadRoom(message.Content[1])
			if !ok {
				return errors.New("room error")
			}
			*room = _room
			(*room).Players[1] = player
			(*room).Game = entity.NewGame()
			var host string
			var guest string
			if (*room).Game.Prior == 0 {
				host = "black"
				guest = "white"
			} else {
				host = "white"
				guest = "black"
			}
			SendMessage(player, entity.NewMessage("guestStartGame", (*room).Players[0].NickName, guest))
			SendMessage((*room).Players[0], entity.NewMessage("hostStartGame", player.NickName, host))
			return nil
		case "step":
			_step, err := strconv.Atoi(message.Content[0])
			if err != nil {
				return err
			}
			victory, _error := (*room).Game.AddStep(entity.NewStep(_step), playerNum)
			_victory := ""
			if _error == "" {
				if victory {
					_victory = "victory"
					SendMessage(player, entity.NewMessage("victory"))
				}
				SendMessage((*room).Players[1-playerNum], entity.NewMessage("opponentStep", message.Content[0], _victory))
			} else {
				SendMessage(player, entity.NewMessage("error", _error))
			}
			return nil
		case "gameFinish":
			(*room).LeaveRoom(2)
		}
	} else {
		return errors.New("sender error")
	}
	return nil
}

func SendMessage(player *entity.Player, message *entity.Message) {
	fmt.Println("给" + player.NickName + "发送了消息")
	fmt.Println(*message)
	player.WebSocket.WriteJSON(message)
}
