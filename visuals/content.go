package visuals

import (
	"image/color"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/igloo"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const pathPrefix = "resources/"

type Content struct {
	// fonts
	LargeFont font.Face
	SmallFont font.Face

	// images
	Ball  *ebiten.Image
	Pixel *ebiten.Image
}

func LoadContent(fs fs.FS) (*Content, error) {
	var err error
	content := &Content{}

	// assets
	fontPath := pathPrefix + "fonts/Sono-Regular.ttf"
	content.LargeFont, err = igloo.LoadOpenType(fs, fontPath, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}

	content.SmallFont, err = igloo.LoadOpenType(fs, fontPath, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}

	content.Ball, err = igloo.LoadImage(fs, pathPrefix+"sprites/ball.png")
	if err != nil {
		return nil, err
	}

	content.Pixel = ebiten.NewImage(1, 1)
	content.Pixel.Set(0, 0, color.White)

	return content, nil
}

func (c *Content) Dispose() {
	// fonts
	c.SmallFont.Close()
	c.LargeFont.Close()

	// images
	c.Ball.Dispose()
}
