package music

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type Manager struct {
	player       *audio.Player
	audioContext *audio.Context
}

func NewLoopMusic(data *bytes.Reader) *Manager {
	aCtx := audioCtx().get()
	stream, streamErr := mp3.DecodeWithSampleRate(aCtx.SampleRate(), data)
	if streamErr != nil {
		panic(streamErr)
	}
	player, playerErr := aCtx.NewPlayer(audio.NewInfiniteLoop(stream, stream.Length()))
	if playerErr != nil {
		panic(playerErr)
	}

	return &Manager{
		audioContext: aCtx,
		player:       player,
	}
}

func (m *Manager) Play() {
	_ = m.player.Rewind()
	if !m.player.IsPlaying() {
		m.player.Play()
	}
}

func (m *Manager) Pause() {
	if m.player.IsPlaying() {
		m.player.Pause()
	}
}

func (m *Manager) Stop() {
	if m.player.IsPlaying() {
		_ = m.player.Close()
	}
}
