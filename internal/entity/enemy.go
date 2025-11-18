package entity

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
  "image"
	"math"
	"math/rand"
	"image/gif"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	X            float64
	Y            float64
	Speed        float64
	Radius       float64
	Color        color.Color
	Health       float64
	Damage       float64
	GifFrames    []*ebiten.Image
	GifDelays    []int
	CurrentFrame int
	FrameTimer   float64
	Facing       string
}

func (e *Enemy) LoadGif(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	gifData, err := gif.DecodeAll(file)
	if err != nil {
		return err
	}

	e.GifFrames = make([]*ebiten.Image, len(gifData.Image))
	e.GifDelays = gifData.Delay

	for i, frame := range gifData.Image {
		bounds := frame.Bounds()
		rgba := image.NewRGBA(bounds)
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				rgba.Set(x, y, frame.At(x, y))
			}
		}
		e.GifFrames[i] = ebiten.NewImageFromImage(rgba)
	}

	return nil
}

func NewEnemy(x, y float64, animatedGifPath string) *Enemy {
	enemy := &Enemy{
		X:           x,
		Y:           y,
		Speed:       50.0,
		Color:       color.RGBA{255, 0, 0, 255},
		Radius:      10.0,
		Health:      10,
		Damage:      1,
		Facing:      "right",
	}

	if animatedGifPath != "" {
		err := enemy.LoadGif(animatedGifPath)
		if err != nil {
			println("Failed to load GIF:", err.Error())
		}
	}

  return enemy

}

func (e *Enemy) Update(dt float64, p *Player, enemies []*Enemy) {
	directionX := p.X - e.X
	directionY := p.Y - e.Y


  if directionX < 0 {
    e.Facing = "left"
  } else {
    e.Facing = "right"
  }

	randomFactor := 0.3

	magnitude := math.Sqrt(directionX*directionX + directionY*directionY)
	normalizedDX := directionX / magnitude
	normalizedDY := directionY / magnitude

	normalizedDX += (rand.Float64() - 0.5) * randomFactor
	normalizedDY += (rand.Float64() - 0.5) * randomFactor

	e.X += normalizedDX * e.Speed * dt
	e.Y += normalizedDY * e.Speed * dt

  


	if len(e.GifFrames) > 0 {
		frameDelay := float64(e.GifDelays[e.CurrentFrame]) / e.Speed
		e.FrameTimer += dt

		if e.FrameTimer >= frameDelay {
			e.FrameTimer = 0
			e.CurrentFrame = (e.CurrentFrame + 1) % len(e.GifFrames)
		}
	}


	for _, otherEnemy := range enemies {
		if otherEnemy != e {
			dx := e.X - otherEnemy.X
			dy := e.Y - otherEnemy.Y
			dist := math.Sqrt(dx*dx + dy*dy)
			minDist := e.Radius + otherEnemy.Radius
			if dist < minDist {
				overlap := minDist - dist
				normalDX := dx / dist
				normalDY := dy / dist

				e.X += normalDX * overlap * 0.5
				e.Y += normalDY * overlap * 0.5
				otherEnemy.X -= normalDX * overlap * 0.5
				otherEnemy.Y -= normalDY * overlap * 0.5

			}
		}

	}

}

func (e *Enemy) Draw(screen *ebiten.Image) {

	if len(e.GifFrames) > 0 {
		op := &ebiten.DrawImageOptions{}

		bounds := e.GifFrames[e.CurrentFrame].Bounds()
		w, h := float64(bounds.Dx()), float64(bounds.Dy())

		if e.Facing == "left" {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(w, 0)
		}

		op.GeoM.Translate(-w/2, -h/2)
		op.GeoM.Translate(e.X, e.Y)

		screen.DrawImage(e.GifFrames[e.CurrentFrame], op)
	} else {
	vector.FillCircle(screen, float32(e.X), float32(e.Y), float32(e.Radius), e.Color, false)
}
}

func (e *Enemy) IsAlive() bool {
	return e.Health > 0
}
