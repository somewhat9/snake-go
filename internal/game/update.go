package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	if g.Status {
		switch {
		case ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.Dir.Y == 0:
			g.Dir = Position{X: 0, Y: -1}
		case ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.Dir.Y == 0:
			g.Dir = Position{X: 0, Y: 1}
		case ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.Dir.X == 0:
			g.Dir = Position{X: 1, Y: 0}
		case ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.Dir.X == 0:
			g.Dir = Position{X: -1, Y: 0}
		}

		now := time.Now()
		if now.Sub(g.LastTick) >= time.Second/4 {
			g.LastTick = now
			g.Tick()
		}
	} else {
		if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyR) {
			g.Setup()
		}
	}

	return nil 
}
