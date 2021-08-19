package utils

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const STAGGER_FRAMES = 7

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
		log.Println("Frame limits are out of image bounds.")
	}

	rec := image.Rect(x0, y0, x0+s.FrameWidth, y0+s.FrameHeight)
	subFrame := s.Image.SubImage(rec).(*ebiten.Image)

	return subFrame
}

/* Updates frames from spritesheet in order to simulate movement.
'animSeq' - A 2D slice, which each element is a {row, column} representing frame positions on spritesheet that compounds a complete animation;
Ex.: {{1,2}, {1,3}, {1,4}}.
'count' - Game counter.*/
func (s *SpriteSheet) UpdateFrame(animSeq [][2]int, count int) {
	index := (count / STAGGER_FRAMES) % len(animSeq)

	s.CurrentFrame = s.subImage(animSeq[index][0], animSeq[index][1])
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
