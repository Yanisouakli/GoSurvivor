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
}

func NewPlayer(x, y float64) *Player {
	return &Player{
		X:         x,
		Y:         y,
		Speed:     200.0,
		Radius:    10.0,
		Color:     color.RGBA{50, 100, 255, 255},
		Health:    100.0,
		MaxHealth: 100.0,
	}
}

func (p *Player) Update(dt float64) {
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
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.FillCircle(screen, float32(p.X), float32(p.Y), float32(p.Radius), p.Color, false)

}
