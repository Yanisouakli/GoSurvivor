package main

import (
	"log"
	"vampsur/internal/config"
	"vampsur/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	cfg := config.Default()
	
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := engine.NewGame(cfg)
	//game.SetScene(scene.NewGameplayScene(cfg))

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
