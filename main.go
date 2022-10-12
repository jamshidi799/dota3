package main

import (
	"game/cmd"
	"game/messenger/dto"
	"game/messenger/event"
	"log"
)

func main() {
	a(2)
	a(event.NewGameStartedEvent([]dto.PlayerDto{}))
	return
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func a(x any) {
	switch x.(type) {
	case int:
		println("int")
	case *event.GameStartedEvent:
		println("string")
	}
}
