package weapon

import (
	"math"
	"vampsur/internal/entity"
)

type Weapon struct {
	Damage   float64
	CoolDown float64
	timer    float64
}

func NewWeapon() *Weapon {
	return &Weapon{
		Damage:   5,
		CoolDown: 0.6,
	}
}

func (w *Weapon) Fire(player *entity.Player, enemies []*entity.Enemy, projectiles *[]*entity.Projectile) {
	if len(enemies) == 0 {
		return
	}

	var nearest *entity.Enemy
	shortestDist := math.MaxFloat64

	for _, e := range enemies {
		dx := player.X - e.X
		dy := player.Y - e.Y
		dist := math.Sqrt(dx*dx + dy*dy)
		if dist < shortestDist {
			shortestDist = dist
			nearest = e
		}
	}

	if nearest == nil {
		return
	}

	p := entity.NewProjectile(
		player.X, player.Y,
		nearest.X, nearest.Y,
		400,
		w.Damage,
	)

	*projectiles = append(*projectiles, p)
}

func (w *Weapon) Update(dt float64, player *entity.Player, enemies []*entity.Enemy, projectiles *[]*entity.Projectile) {
	w.timer += dt

	if w.timer >= w.CoolDown {
		w.timer = 0
		w.Fire(player, enemies, projectiles)
	}

}
