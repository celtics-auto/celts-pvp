package objects

import (
	"image"

	"github.com/celtics-auto/ebiten-chat/client"
	"github.com/celtics-auto/ebiten-chat/utils"
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

func (p *Player) Update(sender chan client.UpdateJson) {
	oldPlayer := &Player{
		Position: &utils.Vector{
			X: p.Position.X,
			Y: p.Position.Y,
		},
		Width:  p.Width,
		Height: p.Height,
	}
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

	if p.Position.X != oldPlayer.Position.X || p.Position.Y != oldPlayer.Position.Y {
		uJson := client.UpdateJson{
			Player: &client.Player{
				Position: client.Vector{
					X: p.Position.X,
					Y: p.Position.Y,
				},
				Width:  p.Width,
				Height: p.Height,
			},
		}

		sender <- uJson
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	rec := image.Rect(0, 0, p.Width, p.Height)
	sub := p.sprite.Image.SubImage(rec).(*ebiten.Image)
	screen.DrawImage(sub, op)
}
