package scene

import (
	"image/color"
	_ "image/png"
	"log"
	"math"
	"slices"
	"vampsur/internal/config"
	"vampsur/internal/entity"
	"vampsur/internal/weapon"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameplayScene struct {
	cfg            *config.Config
	player         *entity.Player
	gameTime       float64
	enemies        []*entity.Enemy
	projectiles    []*entity.Projectile
	weapons        []*weapon.Weapon
	enemiesKilled  int
	paused         bool
	lastUpdateTime float64
}

func NewGameplayScene(cfg *config.Config) *GameplayScene {
	return &GameplayScene{
		cfg: cfg,
	}
}

func (gs *GameplayScene) OnEnter() {

	centerX := float64(gs.cfg.ScreenWidth) / 2
	centerY := float64(gs.cfg.ScreenHeight) / 2
  playerSprite,_,err:= ebitenutil.NewImageFromFile("assets/sprite1.png")
  if err!= nil {
    log.Fatal(err)
  }

	gs.player = entity.NewPlayer(centerX, centerY,playerSprite)

	gs.enemies = []*entity.Enemy{}
	gs.projectiles = []*entity.Projectile{}

	gs.spawnInitialEnemies()
	gs.weapons = []*weapon.Weapon{
		weapon.NewWeapon(),
	}

}

func (gs *GameplayScene) UpdateProjectiles(dt float64) {
	for _, p := range gs.projectiles {
		p.Update(dt)
	}

	for _, p := range gs.projectiles {
		if !p.Active {
			continue
		}

		for _, e := range gs.enemies {
			if !e.IsAlive() {
				continue
			}

			dx := p.X - e.X
			dy := p.Y - e.Y
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist < e.Radius+p.Radius {
				e.Health -= p.Damage
				p.Active = false
				if !e.IsAlive() {
					gs.enemiesKilled++
				}
				break
			}
		}
	}

	alive := gs.projectiles[:0]
	for _, p := range gs.projectiles {
		if p.Active {
			alive = append(alive, p)
		}
	}
	gs.projectiles = alive
}

func (gs *GameplayScene) OnExit() {
	// Cleanup
}

func (gs *GameplayScene) Update() error {
	if gs.paused {
		return nil
	}

	dt := 1.0 / float64(gs.cfg.TPS)
	gs.gameTime += dt

	for _, w := range gs.weapons {
		w.Update(dt, gs.player, gs.enemies, &gs.projectiles)
	}

	if gs.player != nil {
		gs.player.Update(dt, gs.enemies)
    gs.UpdateProjectiles(dt)
	}

	for i, enemy := range gs.enemies {
		if enemy != nil {
			enemy.Update(dt, gs.player, gs.enemies)
      if !enemy.IsAlive() {
        gs.enemies = slices.Delete(gs.enemies,i,i+1)
      }
		}
	}

	return nil
}

func (gs *GameplayScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 30, 255})

	vector.FillRect(screen, float32(gs.player.X-50), float32(gs.player.Y-40), 100, 6, color.RGBA{255, 255, 255, 255}, false)

	healthPercentage := gs.player.Health / gs.player.MaxHealth
	vector.FillRect(screen, float32(gs.player.X-50), float32(gs.player.Y-40), float32(100*healthPercentage), 6, color.RGBA{255, 0, 0, 255}, false)

	for _, enemy := range gs.enemies {
		enemy.Draw(screen)
	}

	for _, p := range gs.projectiles {
		p.Draw(screen)
	}

	if gs.player != nil {
		gs.player.Draw(screen)
	}

}

func (gs *GameplayScene) spawnInitialEnemies() {
	screenW := float64(gs.cfg.ScreenWidth)
	screenH := float64(gs.cfg.ScreenHeight)

	gs.enemies = append(gs.enemies, entity.NewEnemy(100, 50))
	gs.enemies = append(gs.enemies, entity.NewEnemy(300, 50))
	gs.enemies = append(gs.enemies, entity.NewEnemy(500, 50))

	gs.enemies = append(gs.enemies, entity.NewEnemy(200, screenH-50))
	gs.enemies = append(gs.enemies, entity.NewEnemy(400, screenH-50))

	gs.enemies = append(gs.enemies, entity.NewEnemy(50, 200))
	gs.enemies = append(gs.enemies, entity.NewEnemy(50, 400))

	gs.enemies = append(gs.enemies, entity.NewEnemy(screenW-50, 250))
	gs.enemies = append(gs.enemies, entity.NewEnemy(screenW-50, 450))
}
