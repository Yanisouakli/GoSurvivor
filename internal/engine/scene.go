package engine

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	OnEnter()
	OnExit()
}

type SceneManager struct {
	current Scene
	next    Scene
}

// create a sceneManager
func NewSceneManger() *SceneManager {
	return &SceneManager{}
}

func (sm *SceneManager) SetScene(scene Scene) {
	sm.next = scene
}

func (sm *SceneManager) Update() error {
	if sm.next != nil {
		if sm.current != nil {
			sm.current.OnExit()
		}
		sm.current = sm.next
		sm.current.OnEnter()
		sm.next = nil
	}

	if sm.current != nil {
		return sm.current.Update()
	}

	return nil
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	if sm.current != nil {
		sm.current.Draw(screen)
	}
}
