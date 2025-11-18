package entity

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X            float64
	Y            float64
	Speed        float64
	Radius       float64
	Color        color.Color
	Health       float64
	MaxHealth    float64
	GifFrames    []*ebiten.Image 
	GifDelays    []int       
	CurrentFrame int
	FrameTimer   float64
	IsMoving     bool
	Facing       string
}

func NewPlayer(x, y float64, gifPath string) *Player {
	player := &Player{
		X:         x,
		Y:         y,
		Speed:     100.0,
		Radius:    10.0,
		Color:     color.RGBA{50, 100, 255, 255},
		Health:    100.0,
		MaxHealth: 100.0,
		IsMoving:  false,
		Facing:    "right",
	}

	if gifPath != "" {
		err := player.LoadGif(gifPath)
		if err != nil {
			println("Failed to load GIF:", err.Error())
		}
	}

	return player
}

func (p *Player) LoadGif(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	gifData, err := gif.DecodeAll(file)
	if err != nil {
		return err
	}

	p.GifFrames = make([]*ebiten.Image, len(gifData.Image))
	p.GifDelays = gifData.Delay

	for i, frame := range gifData.Image {
		bounds := frame.Bounds()
		rgba := image.NewRGBA(bounds)
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				rgba.Set(x, y, frame.At(x, y))
			}
		}
		p.GifFrames[i] = ebiten.NewImageFromImage(rgba)
	}

	return nil
}

func (p *Player) Update(dt float64, enemies []*Enemy) {
	dx, dy := 0.0, 0.0
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		dy -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dy += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dx -= 1
		p.Facing = "left"
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dx += 1
		p.Facing = "right"
	}

	p.IsMoving = (dx != 0 || dy != 0)

	if p.IsMoving {
		length := math.Sqrt(dx*dx + dy*dy)
		dx /= length
		dy /= length

		p.X += dx * p.Speed * dt
		p.Y += dy * p.Speed * dt
	}

	for _, e := range enemies {
		dx := p.X - e.X
		dy := p.Y - e.Y
		dist := math.Sqrt(dx*dx + dy*dy)
		minDist := e.Radius
		if dist < minDist {
			p.Health = p.Health - e.Damage
			if p.Health < 0 {
				p.Health = 0
			}
		}
	}

	if len(p.GifFrames) > 0 && p.IsMoving {
		frameDelay := float64(p.GifDelays[p.CurrentFrame]) / 100.0
		p.FrameTimer += dt

		if p.FrameTimer >= frameDelay {
			p.FrameTimer = 0
			p.CurrentFrame = (p.CurrentFrame + 1) % len(p.GifFrames)
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	if len(p.GifFrames) > 0 {
		op := &ebiten.DrawImageOptions{}

		bounds := p.GifFrames[p.CurrentFrame].Bounds()
		w, h := float64(bounds.Dx()), float64(bounds.Dy())

		if p.Facing == "left" {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(w, 0)
		}

		op.GeoM.Translate(-w/2, -h/2)
		op.GeoM.Translate(p.X, p.Y)

		screen.DrawImage(p.GifFrames[p.CurrentFrame], op)
	} else {
		vector.FillCircle(screen, float32(p.X), float32(p.Y), float32(p.Radius), p.Color, false)
	}
}
