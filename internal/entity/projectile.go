package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
)

type Projectile struct {
	X, Y       float64
	VelX, VelY float64
	Speed      float64
	Radius     float64
	Color      color.Color
	Active     bool
	Damage     float64
}

func NewProjectile(x, y, targetX, targetY, speed, damage float64) *Projectile {

	dx := targetX - x
	dy := targetY - y
	dist := math.Sqrt(dx*dx + dy*dy)
	if dist == 0 {
		dist = 1
	}

	return &Projectile{
		X:      x,
		Y:      y,
		VelX:   (dx / dist) * speed,
		VelY:   (dy / dist) * speed,
		Speed:  speed,
		Radius: 3,
		Color:  color.RGBA{255, 255, 0, 255},
		Active: true,
		Damage: damage,
	}

}

func (p *Projectile) Update(dt float64) {
	if !p.Active {
		return
	}
	p.X += p.VelX * dt
	p.Y += p.VelY * dt

}

func (p *Projectile) Draw(screen *ebiten.Image) {
	if !p.Active {
		return
	}
	vector.FillCircle(screen, float32(p.X), float32(p.Y), float32(p.Radius), p.Color, false)
}
