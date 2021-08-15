package config

import "github.com/hajimehoshi/ebiten/v2"

// some configs like video, audio, preferences should be loaded from file
// other gameplay settings should be constants

const (
	// video
	ScreenWidth  = 1280
	ScreenHeight = 720

	// ball
	BounceTop  = float64(100)
	BallRadius = float64(32)

	// score
	MaxDistance    = 20
	PixelsPerPoint = 2
)

var (
	// ball
	BounceBottom = float64(ScreenHeight - 100)
	Keybinds     = []ebiten.Key{
		ebiten.KeyQ,
		ebiten.KeyW,
		ebiten.KeyE,
		ebiten.KeyR,
	}
	Timings = []float64{
		2,
		2.3,
		3.1,
		3.5,
	}
)
