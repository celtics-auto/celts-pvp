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
	/* \/   ADDED   \/ */
	animation int  // 0 = 'static', 1 = 'moving', (?) 3 = 'Attacking' (?), ...
	faceRight bool // true = 'facing right', false = 'facing left'
}

func NewPlayer(x, y int, s *utils.SpriteSheet) *Player {
	pl := &Player{
		Position:  utils.NewVector(x, y),
		sprite:    s,
		Width:     64,
		Height:    44,
		animation: 0,
		faceRight: true,
	}
	return pl
}

func (p *Player) Update(sender chan client.UpdateJson, env string) {
	oldPlayer := &Player{
		Position: &utils.Vector{
			X: p.Position.X,
			Y: p.Position.Y,
		},
		Width:     p.Width,
		Height:    p.Height,
		animation: p.animation,
		faceRight: p.faceRight,
	}
	animation := 0 // Default position = 'static'
	faceRight := p.faceRight

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Position.X -= 10
		animation = 1
		faceRight = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Position.X += 10
		animation = 1
		faceRight = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Position.Y -= 10
		animation = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Position.Y += 10
		animation = 1
	}

	p.animation = animation

	if p.Position.X != oldPlayer.Position.X || p.Position.Y != oldPlayer.Position.Y {
		p.faceRight = faceRight

		uJson := client.UpdateJson{
			Player: &client.Player{
				Position: client.Vector{
					X: p.Position.X,
					Y: p.Position.Y,
				},
				Width:     p.Width,
				Height:    p.Height,
				Animation: p.animation,
				FaceRight: p.faceRight,
			},
		}

		if env != "development" {
			sender <- uJson
		}
	}
}

func (p *Player) subImage(lin int, col int) *ebiten.Image {
	x0 := col * p.Width
	y0 := lin * p.Height

	rec := image.Rect(x0, y0, x0+p.Width, y0+p.Height)
	subFrame := p.sprite.Image.SubImage(rec).(*ebiten.Image)

	return subFrame
}

func (p *Player) updateImage(count int) *ebiten.Image {
	var subFrame *ebiten.Image

	if p.animation == 0 {
		col := (count / 5) % 6
		subFrame = p.subImage(3, col) // Static animation
	} else if p.animation == 1 {
		col := (count / 5) % 8
		subFrame = p.subImage(2, col) // Moving animation
	}

	return subFrame
}

func (p *Player) Draw(screen *ebiten.Image, count int) {
	op := &ebiten.DrawImageOptions{}

	if !p.faceRight {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(p.Width), 0)
	}

	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))

	sub := p.updateImage(count)

	screen.DrawImage(sub, op)
}
