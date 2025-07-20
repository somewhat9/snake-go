package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/somewhat9/snake-go/internal/config"
)

func (g *Game) Draw(screen *ebiten.Image) {
	
	screen.Fill(color.Gray{30})
	for y, col := range g.board {
		for x, value := range col {
			if g.Cfg.ViewGrid {
				vector.DrawFilledRect(screen, float32(x*g.Cfg.SquareSize+1+g.Cfg.GridWidthOffset()), float32(y*g.Cfg.SquareSize+1+g.Cfg.GridHeightOffset()), float32(g.Cfg.SquareSize-2), float32(g.Cfg.SquareSize-2), config.Colors[value], false)
			} else {
				vector.DrawFilledRect(screen, float32(x*g.Cfg.SquareSize+g.Cfg.GridWidthOffset()), float32(y*g.Cfg.SquareSize+g.Cfg.GridHeightOffset()), float32(g.Cfg.SquareSize), float32(g.Cfg.SquareSize), config.Colors[value], false)
			}
		}
	}
	text.Draw(screen, "Score: " + fmt.Sprint(len(g.body)-1) + "   High Score: " + fmt.Sprint(g.highScore), g.FontFace, g.Cfg.SquareSize, int(float32(g.Cfg.SquareSize)*1.5), color.White)
	if !g.status {
		message := "GAME OVER!"
		bounds := text.BoundString(g.FontFace, message)
		text.Draw(screen, message, g.FontFace, (g.Cfg.ScreenWidth()-bounds.Dx())/2, (g.Cfg.ScreenHeight()-bounds.Dy())/2, color.White)
	}
}
