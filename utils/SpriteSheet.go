package utils

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SpriteSheet struct {
	Image *ebiten.Image
}

func NewSpriteSheet(path string) (*SpriteSheet, error) {
	eImage, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}
	return &SpriteSheet{
		Image: eImage,
	}, nil
}
