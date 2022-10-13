package main

import (
	"game/cmd"
	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
