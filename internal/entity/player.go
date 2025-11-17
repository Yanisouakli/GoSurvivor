package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"math"
)

type Player struct {
	X         float64
	Y         float64
	Speed     float64
	Radius    float64
	Color     color.Color
	Health    float64
	MaxHealth float64
	Sprite    *ebiten.Image
}

func NewPlayer(x, y float64, sprite *ebiten.Image) *Player {
	return &Player{
		X:         x,
		Y:         y,
		Speed:     200.0,
		Radius:    10.0,
		Color:     color.RGBA{50, 100, 255, 255},
		Health:    100.0,
		MaxHealth: 100.0,
		Sprite:    sprite,
	}
}

func (p *Player) Update(dt float64, enemies []*Enemy) {
	dx, dy := 0.0, 0.0

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		dy -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dy += 1

	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dx -= 1

	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dx += 1
	}

	if dx != 0 || dy != 0 {
		length := math.Sqrt(dx*dx + dy*dy)
		dx /= length
		dy /= length
	}

	p.X += dx * p.Speed * dt

	p.Y += dy * p.Speed * dt

	for _, e := range enemies {
		dx := p.X - e.X
		dy := p.Y - e.Y
		dist := math.Sqrt(dx*dx + dy*dy)
		minDist := e.Radius
		if dist < minDist {
			p.Health = p.Health - e.Damage
			if p.Health < 0 {
				p.Health = 0

			}
		}
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	if p.Sprite != nil {
		op := &ebiten.DrawImageOptions{}

		bounds := p.Sprite.Bounds()
		w, h := float64(bounds.Dx()), float64(bounds.Dy())
		op.GeoM.Translate(-w/2, -h/2)
		op.GeoM.Translate(p.X, p.Y)

		screen.DrawImage(p.Sprite, op)
	} else {
		vector.FillCircle(screen, float32(p.X), float32(p.Y), float32(p.Radius), p.Color, false)
	}
}
