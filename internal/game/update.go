package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	if g.status {
		switch {
		case ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.dir.Y == 0:
			g.dir = position{X: 0, Y: -1}
		case ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.dir.Y == 0:
			g.dir = position{X: 0, Y: 1}
		case ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.dir.X == 0:
			g.dir = position{X: 1, Y: 0}
		case ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.dir.X == 0:
			g.dir = position{X: -1, Y: 0}
		}

		now := time.Now()
		if now.Sub(g.lastTick) >= time.Second/4 {
			g.lastTick = now
			g.tick()
		}
	} else {
		if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyR) {
			g.Setup()
		}
	}

	return nil 
}
