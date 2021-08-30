package devutils

import (
	"image/color"

	"github.com/celtics-auto/ebiten-chat/objects"
	"github.com/celtics-auto/ebiten-chat/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawBox(screen *ebiten.Image, b *utils.BoundingBox, c color.RGBA) {
	ebitenutil.DrawLine(screen, float64(b.V0.X), float64(b.V0.Y), float64(b.V1.X), float64(b.V0.Y), c)
	ebitenutil.DrawLine(screen, float64(b.V0.X), float64(b.V0.Y), float64(b.V0.X), float64(b.V1.Y), c)
	ebitenutil.DrawLine(screen, float64(b.V1.X), float64(b.V1.Y), float64(b.V0.X), float64(b.V1.Y), c)
	ebitenutil.DrawLine(screen, float64(b.V1.X), float64(b.V1.Y), float64(b.V1.X), float64(b.V0.Y), c)
}

func Draw(screen *ebiten.Image, p *objects.Player) {
	//colorGreen := color.RGBA{0, 255, 0, 150}

	var color color.RGBA
	color.A = 150
	color.G = 255

	// [test] Random box
	x := screen.Bounds().Dx() / 2
	y := screen.Bounds().Dy() / 2
	bBox := utils.NewBoundigBox(utils.Vector{X: x - 25, Y: y - 25}, utils.Vector{X: x + 25, Y: y + 25})

	if utils.CheckBoxCollision(bBox, p.HitBox) {
		color.G = 0
		color.R = 255
	}

	DrawBox(screen, bBox, color)
	DrawBox(screen, p.HitBox, color)
}
