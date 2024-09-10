package main

import (
	"flag"
	"image"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game"
)

func main() {
	// flags
	debug := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	// register image format
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	// Create game instance
	g, err := game.NewGame(*debug)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the window size and title
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("The Go Runner")

	// Run the game loop
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
