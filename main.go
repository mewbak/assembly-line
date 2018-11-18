package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.Run(game.Update, 500, 500, 1, "Assembly line")
}
