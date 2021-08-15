package visuals

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/bouncy/config"
	"github.com/miniscruff/igloo"
)

// StaticCamera is a fixed camera that does not move or effect positions.
// Useful for demos or user interface rendering.
type BorderRenderer struct {
	topSprite    *igloo.Sprite
	bottomSprite *igloo.Sprite
}

func NewBorderRenderer(content *Content) *BorderRenderer {
	topY := config.BounceTop - config.BallRadius - 1
	botY := config.BounceBottom + config.BallRadius - 1
	return &BorderRenderer{
		topSprite: igloo.NewSprite(igloo.SpriteConfig{
			Image:     content.Pixel,
			Transform: igloo.NewTransform(0, topY, 0),
			Width:     config.ScreenWidth,
			Height:    3,
		}),
		bottomSprite: igloo.NewSprite(igloo.SpriteConfig{
			Image:     content.Pixel,
			Transform: igloo.NewTransform(0, botY, 0),
			Width:     config.ScreenWidth,
			Height:    3,
		}),
	}
}

func (br *BorderRenderer) Draw(screen *ebiten.Image, camera igloo.Camera) {
	br.topSprite.Draw(screen, camera)
	br.bottomSprite.Draw(screen, camera)
}
