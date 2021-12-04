package entity

import (
	"Gobang-2004A/config"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

type Game struct {
	Steps   [2][]*Step
	Prior   int //0或者1，谁先手
	Current int //0或者1，当前是谁的回合
}

func NewGame() *Game {
	result, _ := rand.Int(rand.Reader, big.NewInt(2))
	number := result.String()
	num, _ := strconv.Atoi(number)
	game := Game{
		Prior:   num,
		Current: num,
	}
	game.Steps[0] = make([]*Step, 0)
	game.Steps[1] = make([]*Step, 0)
	return &game
}

func findFive(list []bool) bool {
	i := 0
	for _, value := range list {
		if value {
			i++
			if i == 5 {
				return true
			}
		} else {
			i = 0
		}
	}
	return false
}

func (game *Game) AddStep(step *Step, player int) (victory bool, err string) {
	fmt.Println("进入addStep")
	if player != game.Current {
		victory = false
		err = "illegalStep"
		return
	}
	boardWidth := config.Config.Get("boardWidth").(int)
	columnsRowEqual := make([]bool, boardWidth)
	rowsColumnEqual := make([]bool, boardWidth)
	rowsSumEqual := make([]bool, boardWidth)
	rowsDiffEqual := make([]bool, boardWidth)
	row := step.Row
	column := step.Column
	sum := row + column
	diff := row - column
	state := 0
	for _, value := range game.Steps[1-player] {
		if value.Row == row && value.Column == column {
			victory = false
			err = "step illegal"
			return
		}
	}
	for _, value := range game.Steps[player] {
		state = 0
		if value.Row == row {
			columnsRowEqual[value.Column] = true
			if state == 0 {
				state = 1
			} else if state == 2 {
				state = 3
			}
		}
		if value.Column == column {
			rowsColumnEqual[value.Row] = true
			if state == 0 {
				state = 2
			} else if state == 1 {
				state = 3
			}
		}
		if state == 3 {
			victory = false
			err = "illegalStep"
			return
		} else if state == 0 {
			if value.Row+value.Column == sum {
				rowsSumEqual[value.Row] = true
			}
			if value.Row-value.Column == diff {
				rowsDiffEqual[value.Row] = true
			}
		}
	}
	game.Steps[player] = append(game.Steps[player], step)
	columnsRowEqual[column] = true
	rowsColumnEqual[row] = true
	rowsSumEqual[row] = true
	rowsDiffEqual[row] = true
	victory = findFive(columnsRowEqual) || findFive(rowsColumnEqual) || findFive(rowsSumEqual) || findFive(rowsDiffEqual)
	err = ""
	game.Current = 1 - player
	return
}
