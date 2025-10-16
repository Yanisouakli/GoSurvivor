package scene

import (
  "vampsur/internal/config"
)

type GameplayScene struct {	
	gameTime        float64
	enemiesKilled   int
	paused          bool
}

func NewGameplayScene(cfg *config.Config) *GameplayScene {
	return &GameplayScene{
	}
}

func (gs *GameplayScene) OnEnter() {
	// Initialize scene
}

func (gs *GameplayScene) OnExit() {
	// Cleanup
}

