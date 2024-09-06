package music

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/tejashwikalptaru/go.run/resources/music"
)

const sampleRate = 44100 // Standard sample rate for audio playback

type Manager struct {
	audioContext    *audio.Context
	backgroundSound *audio.Player
	jumpSound       *audio.Player
	collisionSound  *audio.Player
}

func NewMusicManager() (*Manager, error) {
	// Initialize audio context
	audioContext := audio.NewContext(sampleRate)

	// Load and decode background music
	bgMusicStream, bgMusicStreamErr := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(music.Background))
	if bgMusicStreamErr != nil {
		return nil, bgMusicStreamErr
	}

	// Wrap the background stream in an infinite loop so that it plays continuously
	loopStream := audio.NewInfiniteLoop(bgMusicStream, bgMusicStream.Length())

	// Create a new audio player for background music
	bgMusicPlayer, bgMusicPlayerErr := audioContext.NewPlayer(loopStream)
	if bgMusicPlayerErr != nil {
		return nil, bgMusicPlayerErr
	}

	// Load and decode jump sound
	jumpStream, jumpStreamErr := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(music.Jump))
	if jumpStreamErr != nil {
		return nil, jumpStreamErr
	}

	// Create a new audio player for jump sound (not looped)
	jumpPlayer, jumpPlayerErr := audioContext.NewPlayer(jumpStream)
	if jumpPlayerErr != nil {
		return nil, jumpPlayerErr
	}

	// Load and decode collision sound
	collisionStream, collisionStreamErr := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(music.Collision))
	if collisionStreamErr != nil {
		return nil, collisionStreamErr
	}

	// Create a new audio player for collision sound (not looped)
	collisionPlayer, collisionPlayerErr := audioContext.NewPlayer(collisionStream)
	if collisionPlayerErr != nil {
		return nil, collisionPlayerErr
	}

	return &Manager{
		audioContext:    audioContext,
		backgroundSound: bgMusicPlayer,
		jumpSound:       jumpPlayer,
		collisionSound:  collisionPlayer,
	}, nil
}

// PlayBackground starts the background music if it is not already playing
func (m *Manager) PlayBackground() {
	if !m.backgroundSound.IsPlaying() {
		m.backgroundSound.Play()
	}
}

// StopBackground pauses the background music
func (m *Manager) StopBackground() {
	if m.backgroundSound.IsPlaying() {
		m.backgroundSound.Pause()
	}
}

// ResetBackground restarts the background music
func (m *Manager) ResetBackground() {
	err := m.backgroundSound.Rewind()
	if err != nil {
		fmt.Printf("failed to rewind background sound: %v", err)
	}
	m.PlayBackground()
}

// PlayJumpSound plays the jump sound effect
func (m *Manager) PlayJumpSound() {
	// Rewind ensures the jump sound starts from the beginning
	err := m.jumpSound.Rewind()
	if err != nil {
		fmt.Printf("failed to rewind jump sound: %v", err)
	}
	m.jumpSound.Play()
}

// PlayCollisionSound plays the collision sound effect
func (m *Manager) PlayCollisionSound() {
	// Rewind ensures the collision sound starts from the beginning
	err := m.collisionSound.Rewind()
	if err != nil {
		fmt.Printf("failed to rewind collision sound: %v", err)
	}
	m.collisionSound.Play()
}
