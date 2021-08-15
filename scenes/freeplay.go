package scenes

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/bouncy/config"
	"github.com/miniscruff/bouncy/controllers"
	"github.com/miniscruff/igloo"

	// "github.com/miniscruff/bouncy/models"
	"github.com/miniscruff/bouncy/visuals"
)

type FreeplayScene struct {
	game     igloo.Game
	content  *visuals.Content
	isPaused bool

	// controllers
	uiCamera *controllers.StaticCamera
	score    *controllers.Score
	balls    []*controllers.Ball

	// visuals
	fpsTracker     *visuals.FPSTracker
	ballRenderers  []*visuals.BallRenderer
	borderRenderer *visuals.BorderRenderer
	scoreRenderer  *visuals.ScoreRenderer
}

func (s *FreeplayScene) Setup(game igloo.Game, fs fs.FS) error {
	var err error

	s.game = game
	s.content, err = visuals.LoadContent(fs)
	if err != nil {
		return err
	}

	// controllers
	s.uiCamera = &controllers.StaticCamera{}
	s.score = controllers.NewScore()

	// visuals
	s.fpsTracker, err = visuals.NewFPSTracker(s.content)
	if err != nil {
		return err
	}

	for i := 0; i < 4; i++ {
		ball := controllers.NewBall(controllers.BallConfig{
			X:      float64(config.ScreenWidth / 5 * (i + 1)),
			Key:    config.Keybinds[i],
			Timing: config.Timings[i],
			Score:  s.score,
		})
		br := visuals.NewBallRenderer(s.content, ball)

		s.balls = append(s.balls, ball)
		s.ballRenderers = append(s.ballRenderers, br)
	}

	s.borderRenderer = visuals.NewBorderRenderer(s.content)
	s.scoreRenderer, err = visuals.NewScoreRenderer(s.content, s.score)

	return nil
}

// Update all game elements
func (s *FreeplayScene) Update(deltaTime float64) {
	// should add a debug toggle
	s.fpsTracker.Update(deltaTime)
	for _, b := range s.balls {
		b.Update(deltaTime)
	}
}

// Draw all game elements
func (s *FreeplayScene) Draw(screen *ebiten.Image) {
	s.fpsTracker.Draw(screen)
	s.borderRenderer.Draw(screen, s.uiCamera)
	for _, br := range s.ballRenderers {
		br.Draw(screen, s.uiCamera)
	}
	s.scoreRenderer.Draw(screen)
}

// Dispose of game content or data
func (s *FreeplayScene) Dispose() {
	s.content.Dispose()
}
