package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"vampsur/internal/config"
)

type Game struct {
	config      *config.Config
	sceneManager *SceneManager
}


func NewGame(cfg *config.Config) *Game{
  return &Game{
    config:   cfg,
    sceneManager:  NewSceneManger(),
  }
}

//implement the scne interface methods


func (g *Game) SetScene(scene Scene){
  g.sceneManager.SetScene(scene)
}


func (g *Game) Update() error {
	return g.sceneManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.config.ScreenWidth, g.config.ScreenHeight
}


