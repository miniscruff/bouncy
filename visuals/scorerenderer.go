package visuals

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/bouncy/config"
	"github.com/miniscruff/bouncy/controllers"
	"github.com/miniscruff/igloo"
)

// StaticCamera is a fixed camera that does not move or effect positions.
// Useful for demos or user interface rendering.
type ScoreRenderer struct {
	label *igloo.Label
}

func NewScoreRenderer(content *Content, score *controllers.Score) (*ScoreRenderer, error) {
	label, err := igloo.NewLabel(igloo.LabelOptions{
		Font:      content.LargeFont,
		Transform: igloo.NewTransform(config.ScreenWidth/2, 10, 0),
		Anchor:    igloo.AnchorTopCenter,
		Color:     color.White,
	})
	score.AddWatcher(func(newScore int) {
		label.SetText(fmt.Sprintf("%v", newScore))
	})
	if err != nil {
		return nil, err
	}
	return &ScoreRenderer{
		label: label,
	}, nil
}

func (sr *ScoreRenderer) Draw(screen *ebiten.Image) {
	sr.label.Draw(screen)
}
