package main

import (
	"log"
	"vampsur/internal/config"
	"vampsur/internal/engine"
	"vampsur/internal/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	cfg := config.Default()

	game := engine.NewGame(cfg)

	gameplayScene := scene.NewGameplayScene(cfg)
	game.SetScene(gameplayScene)

	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetTPS(cfg.TPS)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
