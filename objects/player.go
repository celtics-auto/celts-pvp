package objects

import (
	"fmt"
	"math"

	"github.com/celtics-auto/ebiten-chat/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Position  *utils.Vector
	sprite    *utils.SpriteSheet
	HitBox    *utils.BoundingBox
	Width     int
	Height    int
	Animation int    // 0 = 'static', 1 = 'moving', (?) 3 = 'Attacking' (?), ...
	Face      string // 'N', 'S', 'E', 'W', 'NE', 'NW' ...
	Speed     int
}

func NewPlayer(x, y int, s *utils.SpriteSheet) *Player {
	pl := &Player{
		Position:  utils.NewVector(x, y),
		sprite:    s,
		HitBox:    utils.NewBoundigBox(utils.Vector{X: x - s.FrameWidth/2, Y: y - s.FrameHeight/2}, utils.Vector{X: x + s.FrameWidth/2, Y: y + s.FrameHeight/2}),
		Width:     s.FrameWidth,
		Height:    s.FrameHeight,
		Animation: 0,
		Face:      "S",
		Speed:     10,
	}
	return pl
}

func (p *Player) Update() bool {
	oldPlayer := &Player{
		Position: &utils.Vector{
			X: p.Position.X,
			Y: p.Position.Y,
		},
		Width:     p.Width,
		Height:    p.Height,
		Animation: p.Animation,
		Face:      p.Face,
	}
	p.Animation = 0 // Default position = 'static'
	face := ""

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		face = fmt.Sprintf("%s%s", face, "N")
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		face = fmt.Sprintf("%s%s", face, "S")
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		face = fmt.Sprintf("%s%s", face, "E")
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		face = fmt.Sprintf("%s%s", face, "W")
	}

	switch face {
	case "N":
		p.Position.Y -= p.Speed
	case "S":
		p.Position.Y += p.Speed
	case "E":
		p.Position.X += p.Speed
	case "W":
		p.Position.X -= p.Speed
	case "NE":
		p.Position.X += int(float64(p.Speed) * math.Sin(math.Pi/4))
		p.Position.Y -= int(float64(p.Speed) * math.Sin(math.Pi/4))
	case "NW":
		p.Position.X -= int(float64(p.Speed) * math.Sin(math.Pi/4))
		p.Position.Y -= int(float64(p.Speed) * math.Sin(math.Pi/4))
	case "SE":
		p.Position.X += int(float64(p.Speed) * math.Sin(math.Pi/4))
		p.Position.Y += int(float64(p.Speed) * math.Sin(math.Pi/4))
	case "SW":
		p.Position.X -= int(float64(p.Speed) * math.Sin(math.Pi/4))
		p.Position.Y += int(float64(p.Speed) * math.Sin(math.Pi/4))
	}

	if p.Position.X != oldPlayer.Position.X || p.Position.Y != oldPlayer.Position.Y {
		p.Face = face
		p.Animation = 1
		p.HitBox.V0.X = p.Position.X - p.Width/2
		p.HitBox.V0.Y = p.Position.Y - p.Height/2
		p.HitBox.V1.X = p.Position.X + p.Width/2
		p.HitBox.V1.Y = p.Position.Y + p.Height/2

		return true
	}

	return false
}

func (p *Player) updatePlayerFrame(count int) {
	animSeq := make([][2]int, 0) // Spritesheet frames {row, col} indexes

	m := make(map[string]int) // Spritesheet row indexes for each direction
	m["N"] = 0
	m["S"] = 1
	m["E"] = 3
	m["W"] = 6
	m["NE"] = 4
	m["NW"] = 7
	m["SE"] = 2
	m["SW"] = 5

	switch p.Animation {
	case 0:
		animSeq = append(animSeq, [2]int{m[p.Face], 4})
	case 1:
		for i := 0; i <= 3; i++ {
			animSeq = append(animSeq, [2]int{m[p.Face], i})
		}
	}

	p.sprite.UpdateFrame(animSeq, count)

}

func (p *Player) Draw(screen *ebiten.Image, count int) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(p.Position.X-p.sprite.FrameWidth/2), float64(p.Position.Y-p.sprite.FrameHeight/2))

	p.updatePlayerFrame(count)

	screen.DrawImage(p.sprite.CurrentFrame, op)
}
