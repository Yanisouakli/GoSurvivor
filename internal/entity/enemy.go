package entity

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	X      float64
	Y      float64
	Speed  float64
	Radius float64
	Color  color.Color
	Health float64
}

func NewEnemy(x, y float64) *Enemy {
	return &Enemy{
		X:      x,
		Y:      y,
		Color:  color.RGBA{255, 0, 0, 255},
		Radius: 10.0,
		Health: 10,
	}

}

func (e *Enemy) Update(dt float64) {
	// Enemy behavior will be added later
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	vector.FillCircle(screen, float32(e.X), float32(e.Y), float32(e.Radius), e.Color, false)
}

func (e *Enemy) IsAlive() bool {
	return e.Health > 0
}
