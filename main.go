package main

import (
	"Gobang-2004A/router"

	"Gobang-2004A/config"
)

func main() {
	config.Init()
	router.InitRouter()
}
