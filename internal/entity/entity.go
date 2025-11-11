package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Draw(screen *ebiten.Image)
	Update(dt float64)
	Destroy()
	IsActive() bool
}

type BaseEntity struct {
	Active bool
}

func NewBaseEntity(x, y float64) *BaseEntity {
	return &BaseEntity{
		Active: true,
	}
}

func (e *BaseEntity) IsActive() bool {
	return e.Active
}

func (e *BaseEntity) Destroy() {
	e.Active = false
}
