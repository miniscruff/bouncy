package visuals

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/igloo"
)

// StaticCamera is a fixed camera that does not move or effect positions.
// Useful for demos or user interface rendering.
type FPSTracker struct {
	label  *igloo.Label
	frames int64
	timer  float64
}

func NewFPSTracker(content *Content) (*FPSTracker, error) {
	label, err := igloo.NewLabel(igloo.LabelOptions{
		Font:      content.SmallFont,
		Transform: igloo.NewTransform(5, 5, 0),
		Color:     color.White,
	})
	if err != nil {
		return nil, err
	}
	return &FPSTracker{
		label: label,
	}, nil
}

func (fps *FPSTracker) Update(deltaTime float64) {
	fps.frames++
	fps.timer += deltaTime
	if fps.timer > 0.25 {
		fps.timer -= 0.25
		fps.label.SetText(fmt.Sprintf("%v", fps.frames*4))
		fps.frames = 0
	}
}

func (fps *FPSTracker) Draw(screen *ebiten.Image) {
	fps.label.Draw(screen)
}
