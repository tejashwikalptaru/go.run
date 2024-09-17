package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game"
)

func main() {
	// flags
	debug := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	// Create game instance
	g, err := game.NewGame(*debug)
	if err != nil {
		log.Fatal(err)
	}
	// Run the game loop
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
