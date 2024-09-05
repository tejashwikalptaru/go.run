package main

import (
	"github.com/tejashwikalptaru/go-dino/resources/fonts"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tejashwikalptaru/go-dino/game"
)

func main() {
	// Initialize game and load font
	fontFace := fonts.LoadFont(fonts.ManaSpace)

	// Create game instance and set font
	g := game.NewGame(fontFace)

	// Set up the window size and title
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Dino Game in Go")

	// Run the game loop
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
