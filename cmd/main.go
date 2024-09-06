package main

import (
	"log"

	"github.com/tejashwikalptaru/go.run/resources/fonts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go.run/game"
)

func main() {
	// Initialize game and load font
	fontFace := fonts.LoadFont(fonts.ManaSpace)

	// Create game instance and set font
	g, err := game.NewGame(fontFace)
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
