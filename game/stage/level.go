package stage

import (
	"fmt"
	"image/color"
	"time"

	"github.com/tejashwikalptaru/go.run/resources/fonts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Level struct {
	countdownStart     time.Time
	textFaceSource     *text.GoTextFaceSource
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

func NewLevel(screenWidth, screenHeight float64, textFaceSource *text.GoTextFaceSource, levelJumpThreshold int) *Level {
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
		textFaceSource:     textFaceSource,
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
		msg := fmt.Sprintf("Level %d", l.level)
		op := &text.DrawOptions{}
		op.GeoM.Translate(l.screenWidth/3, l.screenHeight/6)
		op.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 255, B: 255, A: 255})
		text.Draw(screen, msg, &text.GoTextFace{
			Source: l.textFaceSource,
			Size:   fonts.DefaultTextSize,
		}, op)

		// Countdown logic: Fade-in/out based on alpha value
		countdownText := fmt.Sprintf("Ready... %d", l.countdown)
		op1 := &text.DrawOptions{}
		op1.GeoM.Translate(l.screenWidth/3, l.screenHeight/3)
		op1.ColorScale.ScaleAlpha(float32(l.countdownAlpha * 255))
		op1.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 0, B: 0})
		text.Draw(screen, countdownText, &text.GoTextFace{
			Source: l.textFaceSource,
			Size:   fonts.DefaultTextSize,
		}, op1)

		// Update alpha value for smooth fade in/out
		l.countdownAlpha -= 0.05
		if l.countdownAlpha < 0 {
			l.countdownAlpha = 0
		}
	}
}
