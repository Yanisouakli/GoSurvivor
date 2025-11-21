package scene

import (
	"image/color"
	_ "image/png"
	"math"
	"math/rand"
	"slices"
	"vampsur/internal/config"
	"vampsur/internal/entity"
	"vampsur/internal/weapon"

	"github.com/hajimehoshi/ebiten/v2"
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

	camX float64
	camY float64
}

const SpawnMinDistance = 500
const SpawnMaxDistance = 700

func NewGameplayScene(cfg *config.Config) *GameplayScene {
	return &GameplayScene{
		cfg: cfg,
	}
}

func (gs *GameplayScene) OnEnter() {

	centerX := float64(gs.cfg.ScreenWidth) / 2
	centerY := float64(gs.cfg.ScreenHeight) / 2

	gifPath := "assets/sprite.gif"
	enemyGifPath := "assets/mortaccio.gif"

	gs.player = entity.NewPlayer(centerX, centerY, gifPath)

	gs.enemies = []*entity.Enemy{}
	gs.projectiles = []*entity.Projectile{}

	gs.spawnInitialEnemies(enemyGifPath)
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
	enemyGifPath := "assets/mortaccio.gif"
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

	gs.UpdateCamera()

	gs.spawnEnemies(enemyGifPath, dt)

	for i, enemy := range gs.enemies {
		if enemy != nil {
			enemy.Update(dt, gs.player, gs.enemies)
			if !enemy.IsAlive() {
				gs.enemies = slices.Delete(gs.enemies, i, i+1)
			}
		}
	}

	return nil
}

func (gs *GameplayScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 30, 255})

	vector.FillRect(screen, float32(gs.player.X-50-gs.camX), float32(gs.player.Y-40-gs.camY), 100, 6, color.RGBA{255, 255, 255, 255}, false)

	healthPercentage := gs.player.Health / gs.player.MaxHealth
	vector.FillRect(screen, float32(gs.player.X-50-gs.camX), float32(gs.player.Y-40-gs.camY), float32(100*healthPercentage), 6, color.RGBA{255, 0, 0, 255}, false)

	for _, enemy := range gs.enemies {
		enemy.Draw(screen, gs.camX, gs.camY)
	}

	for _, p := range gs.projectiles {
		p.Draw(screen, gs.camX, gs.camY)
	}

	if gs.player != nil {
		gs.player.Draw(screen, gs.camX, gs.camY)
	}

}

func (gs *GameplayScene) spawnEnemies(enemyGifPath string, dt float64) {
	if len(gs.enemies) >= 5 {
		return
	}

	for i := 0; i < 5; i++ {

		angle := rand.Float64() * 2 * math.Pi

		dist := SpawnMinDistance + rand.Float64()*(SpawnMaxDistance-SpawnMinDistance)

		spawnX := gs.player.X + math.Cos(angle)*dist
		spawnY := gs.player.Y + math.Sin(angle)*dist

		gs.enemies = append(gs.enemies, entity.NewEnemy(spawnX, spawnY, enemyGifPath))
	}
}

func (gs *GameplayScene) spawnInitialEnemies(enemyGifPath string) {
	screenW := float64(gs.cfg.ScreenWidth)
	screenH := float64(gs.cfg.ScreenHeight)

	gs.enemies = append(gs.enemies, entity.NewEnemy(100, 50, enemyGifPath))
	gs.enemies = append(gs.enemies, entity.NewEnemy(300, 50, enemyGifPath))
	gs.enemies = append(gs.enemies, entity.NewEnemy(500, 50, enemyGifPath))

	gs.enemies = append(gs.enemies, entity.NewEnemy(200, screenH-50, enemyGifPath))
	gs.enemies = append(gs.enemies, entity.NewEnemy(400, screenH-50, enemyGifPath))

	gs.enemies = append(gs.enemies, entity.NewEnemy(50, 200, enemyGifPath))
	gs.enemies = append(gs.enemies, entity.NewEnemy(50, 400, enemyGifPath))

	gs.enemies = append(gs.enemies, entity.NewEnemy(screenW-50, 250, enemyGifPath))
	gs.enemies = append(gs.enemies, entity.NewEnemy(screenW-50, 450, enemyGifPath))
}

func (gs *GameplayScene) UpdateCamera() {
	gs.camX = gs.player.X - float64(gs.cfg.ScreenWidth)/2
	gs.camY = gs.player.Y - float64(gs.cfg.ScreenHeight)/2
}
