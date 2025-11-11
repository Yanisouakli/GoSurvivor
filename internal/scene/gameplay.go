package scene

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"vampsur/internal/config"
	"vampsur/internal/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameplayScene struct {
	cfg            *config.Config
	player         *entity.Player
	gameTime       float64
	enemies        []*entity.Enemy
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
	gs.player = entity.NewPlayer(centerX, centerY)

	gs.enemies = []*entity.Enemy{}

	gs.spawnInitialEnemies()

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

	if gs.player != nil {
		gs.player.Update(dt)
	}

	for _, enemy := range gs.enemies {
		if enemy != nil {
			enemy.Update(dt, gs.player, gs.enemies)
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
