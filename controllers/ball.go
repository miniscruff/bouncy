package controllers

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/miniscruff/bouncy/config"
)

type BallConfig struct {
	Key    ebiten.Key
	X      float64
	Timing float64
	Score  *Score
}

type Ball struct {
	x         float64
	y         float64
	velocity  float64
	direction float64
	key       ebiten.Key
	score     *Score
}

func NewBall(ballConfig BallConfig) *Ball {
	return &Ball{
		key:       ballConfig.Key,
		x:         ballConfig.X,
		y:         config.BounceBottom,
		velocity:  float64((config.BounceBottom - config.BounceTop) / ballConfig.Timing),
		direction: -1,
		score:     ballConfig.Score,
	}
}

func (b *Ball) Position() (float64, float64) {
	return b.x, b.y
}

func (b *Ball) Update(deltaTime float64) {
	b.y += float64(b.velocity * b.direction * deltaTime)

	if b.direction == 1 && b.y >= config.BounceBottom {
		b.direction = -1
		b.y = config.BounceBottom
	} else if b.direction == -1 && b.y <= config.BounceTop {
		b.direction = 1
		b.y = config.BounceTop
	}

	if inpututil.IsKeyJustPressed(b.key) {
		distance := int(math.Abs(config.BounceBottom - b.y))
		if distance > config.MaxDistance {
			fmt.Printf("Too far: %v\n", distance)
		} else {
			points := int((config.MaxDistance - distance) / config.PixelsPerPoint)
			b.score.Add(points)
			fmt.Printf("Score +%v = %v\n", points, b.score.Value())
		}
	}
}
