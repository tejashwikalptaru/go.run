package stage

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Level struct {
	countdownStart     time.Time
	fontFace           font.Face
	countdownAlpha     float64
	countdown          int
	levelJumpThreshold int
	level              int
	jumps              int
	score              int
	screenWidth        float64
	screenHeight       float64
	isFirstLevel       bool
	gameOver           bool
	inLevelGreeting    bool
}

func NewLevel(screenWidth, screenHeight float64, fontFace font.Face, levelJumpThreshold int) *Level {
	return &Level{
		gameOver:           false,
		inLevelGreeting:    true,
		isFirstLevel:       true,
		countdownStart:     time.Now(),
		countdown:          3,
		countdownAlpha:     1.0,
		levelJumpThreshold: levelJumpThreshold,
		level:              1,
		jumps:              0,
		score:              0,
		fontFace:           fontFace,
		screenWidth:        screenWidth,
		screenHeight:       screenHeight,
	}
}

func (l *Level) IsGreeting() bool {
	return l.inLevelGreeting
}

func (l *Level) Score() int {
	return l.score
}

func (l *Level) IncreaseScore() {
	l.jumps++
	l.score += 10
}

func (l *Level) Clear() bool {
	return l.jumps >= l.levelJumpThreshold
}

// Next moves the level to next stage
func (l *Level) Next() {
	l.isFirstLevel = false
	l.level++
	l.jumps = 0 // ResetToFirst the jump counter for the next stage
	l.inLevelGreeting = true
	l.countdown = 3 // Start countdown for new stage
	l.countdownAlpha = 1.0
	l.countdownStart = time.Now()
}

// handleCountdown manages the countdown before each stage
func (l *Level) handleCountdown() {
	elapsed := time.Since(l.countdownStart).Seconds()

	// Handle fade-in/out based on elapsed time
	if elapsed > 1 {
		l.countdown--
		l.countdownAlpha = 1.0
		l.countdownStart = time.Now()
	}

	if l.countdown < 0 {
		l.inLevelGreeting = false
	}
}

func (l *Level) Update() {
	// If we're in the stage greeting phase, manage the countdown
	if l.inLevelGreeting {
		l.handleCountdown()
		return
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	// If we're in the stage greeting phase, show the greeting and countdown
	if l.inLevelGreeting {
		msg := fmt.Sprintf("Welcome to Level %d", l.level)
		text.Draw(screen, msg, l.fontFace, int(l.screenWidth/4), int(l.screenHeight/3), color.RGBA{R: 255, G: 255, B: 255, A: 255})

		// Countdown logic: Fade-in/out based on alpha value
		countdownColor := color.RGBA{R: 255, G: 0, B: 0, A: uint8(l.countdownAlpha * 255)}
		countdownText := fmt.Sprintf("%d", l.countdown)
		text.Draw(screen, countdownText, l.fontFace, int(l.screenWidth/2), int(l.screenHeight/2), countdownColor)

		// Update alpha value for smooth fade in/out
		l.countdownAlpha -= 0.05
		if l.countdownAlpha < 0 {
			l.countdownAlpha = 0
		}
	}
}
