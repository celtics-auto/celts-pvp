package objects

import (
	"image"

	"github.com/celtics-auto/ebiten-chat/utils"
	"github.com/google/go-cmp/cmp"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Position *utils.Vector
	Width    int
	Height   int
	sprite   *utils.SpriteSheet
}

func NewPlayer(x, y int, s *utils.SpriteSheet) *Player {
	pl := &Player{
		Position: utils.NewVector(x, y),
		sprite:   s,
		Width:    64,
		Height:   44,
	}
	return pl
}

func (p *Player) Update() *Player {
	oldPlayer := *p
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Position.X -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Position.X += 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Position.Y -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Position.Y += 10
	}

	// FIXME: exclude Player.sprite from compare
	if !cmp.Equal(oldPlayer, *p) {
		return p
	}
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	rec := image.Rect(0, 0, p.Width, p.Height)
	sub := p.sprite.Image.SubImage(rec).(*ebiten.Image)
	screen.DrawImage(sub, op)
}
