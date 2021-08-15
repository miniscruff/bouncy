package main

import (
	"embed"
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/miniscruff/bouncy/config"
	"github.com/miniscruff/bouncy/scenes"
	"github.com/miniscruff/igloo"
)

var (
	//go:embed resources
	resourcesFS embed.FS
)

type Game struct {
	scene igloo.Scene
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func (g *Game) Update() error {
	if g.scene == nil {
		return nil
	}

	deltaTime := 1 / ebiten.CurrentTPS()
	if math.IsInf(deltaTime, 0) {
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return fmt.Errorf("Escape shutdown")
	}

	g.scene.Update(deltaTime)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.scene == nil {
		return
	}
	deltaTime := 1.0 / ebiten.CurrentTPS()
	if math.IsInf(deltaTime, 0) {
		return
	}
	g.scene.Draw(screen)
}

func (g *Game) LoadScene(newScene igloo.Scene) error {
	if g.scene != nil {
		g.scene.Dispose()
	}
	g.scene = newScene
	return g.scene.Setup(g, resourcesFS)
}

func main() {
	ebiten.SetWindowTitle("Bouncy")
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)

	g := &Game{}
	fmt.Println("loading scene")
	if err := g.LoadScene(&scenes.FreeplayScene{}); err != nil {
		fmt.Printf("starting scene failed to load: %s", err.Error())
		return
	}
	if err := ebiten.RunGame(g); err != nil {
		fmt.Printf("run game: %s", err.Error())
	}
}
