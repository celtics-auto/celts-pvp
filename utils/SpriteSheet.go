package utils

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SpriteSheet struct {
	Image        *ebiten.Image
	CurrentFrame *ebiten.Image
	FrameWidth   int
	FrameHeight  int
	imageWidth   int
	imageHeight  int
}

func (s *SpriteSheet) subImage(lin int, col int) *ebiten.Image {
	x0 := col * s.FrameWidth
	y0 := lin * s.FrameHeight

	if x0 > s.imageWidth || x0+s.FrameWidth > s.imageWidth ||
		y0 > s.imageHeight || y0+s.FrameHeight > s.imageHeight {
		log.Println("Frame limits are out of image bounds.", x0, y0, s.imageWidth, s.imageHeight)
	}

	rec := image.Rect(x0, y0, x0+s.FrameWidth, y0+s.FrameHeight)
	subFrame := s.Image.SubImage(rec).(*ebiten.Image)

	return subFrame
}

func (s *SpriteSheet) UpdatePlayerFrame(face string, animation int, count int) {
	m := make(map[string]int)

	m["N"] = 0
	m["S"] = 1
	m["E"] = 3
	m["W"] = 6
	m["NE"] = 4
	m["NW"] = 7
	m["SE"] = 2
	m["SW"] = 5

	if animation == 0 {
		// Static animation
		col := 4
		s.CurrentFrame = s.subImage(m[face], col)
	} else if animation == 1 {
		// Moving animation
		col := (count / 7) % 4
		s.CurrentFrame = s.subImage(m[face], col)
	}
}

func NewSpriteSheet(path string, w int, h int) (*SpriteSheet, error) {
	eImage, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}
	imgWidth, imgHeight := eImage.Size()
	return &SpriteSheet{
		Image:       eImage,
		FrameWidth:  w,
		FrameHeight: h,
		imageWidth:  imgWidth,
		imageHeight: imgHeight,
	}, nil
}
