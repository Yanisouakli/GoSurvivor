package entity

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	X      float64
	Y      float64
	Speed  float64
	Radius float64
	Color  color.Color
	Health float64
  Damage float64
}

func NewEnemy(x, y float64) *Enemy {
	return &Enemy{
		X:      x,
		Y:      y,
		Speed:  100.0,
		Color:  color.RGBA{255, 0, 0, 255},
		Radius: 10.0,
		Health: 10,
    Damage: 1,
	}

}

func (e *Enemy) Update(dt float64, p *Player, enemies []*Enemy) {
	directionX := p.X - e.X
	directionY := p.Y - e.Y

	randomFactor := 0.3

	magnitude := math.Sqrt(directionX*directionX + directionY*directionY)
	normalizedDX := directionX / magnitude
	normalizedDY := directionY / magnitude

	normalizedDX += (rand.Float64() - 0.5) * randomFactor
	normalizedDY += (rand.Float64() - 0.5) * randomFactor

	e.X += normalizedDX * e.Speed * dt
	e.Y += normalizedDY * e.Speed * dt

	for _, otherEnemy := range enemies {
		if otherEnemy != e {
			dx := e.X - otherEnemy.X
			dy := e.Y - otherEnemy.Y
			dist := math.Sqrt(dx*dx + dy*dy)
			minDist := e.Radius + otherEnemy.Radius
			if dist < minDist {
				overlap := minDist - dist
				normalDX := dx / dist
				normalDY := dy / dist

				e.X += normalDX * overlap * 0.5
				e.Y += normalDY * overlap * 0.5
				otherEnemy.X -= normalDX * overlap * 0.5
				otherEnemy.Y -= normalDY * overlap * 0.5

			}
		}

	}

}

func (e *Enemy) Draw(screen *ebiten.Image) {
	vector.FillCircle(screen, float32(e.X), float32(e.Y), float32(e.Radius), e.Color, false)
}

func (e *Enemy) IsAlive() bool {
	return e.Health > 0
}


