package visuals

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/bouncy/controllers"
	"github.com/miniscruff/igloo"
)

// StaticCamera is a fixed camera that does not move or effect positions.
// Useful for demos or user interface rendering.
type BallRenderer struct {
	sprite *igloo.Sprite
	ball   *controllers.Ball
}

func NewBallRenderer(content *Content, ball *controllers.Ball) *BallRenderer {
	x, y := ball.Position()
	sprite := igloo.NewSprite(igloo.SpriteConfig{
		Image:     content.Ball,
		Transform: igloo.NewTransform(x, y, 0),
		Anchor:    igloo.AnchorMiddleCenter,
	})
	return &BallRenderer{
		sprite: sprite,
		ball:   ball,
	}
}

func (br *BallRenderer) Draw(screen *ebiten.Image, camera igloo.Camera) {
	x, y := br.ball.Position()
	br.sprite.Transform.SetPosition(x, y)
	br.sprite.Draw(screen, camera)
}
