package music

import (
	"sync"
	
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type ctx struct {
	ctx *audio.Context
}

var (
	instance *ctx
	once     sync.Once
)

// ctx returns the singleton instance of audio context
func audioCtx() *ctx {
	once.Do(func() {
		instance = &ctx{
			ctx: audio.NewContext(44100),
		}
	})
	return instance
}

func (a *ctx) get() *audio.Context {
	return a.ctx
}
