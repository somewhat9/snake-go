package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ViewGrid bool = false
	SquareSize = 30

	GridSquaresWidth = 40
	GridSquaresHeight = 30

	GridWidth = GridSquaresWidth*SquareSize
	GridHeight = GridSquaresHeight*SquareSize

	GridHeightOffset = SquareSize*1
	GridWidthOffset = SquareSize*0

	ScreenWidth = GridWidth + GridWidthOffset
	ScreenHeight = GridHeight + GridHeightOffset
)

var Board [GridSquaresHeight][GridSquaresWidth]uint8

var Colors = map[uint8]color.Color{
	0: color.Black, 
	1: color.RGBA{0, 255, 0, 255}, 
	2: color.RGBA{255, 0, 0, 255},
}

type Game struct{
	Snake
}

type Snake struct{
	Position
}

type Position struct {
	X, Y float32
}

func (g *Game) Update() error {
	return nil 
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for y, col := range Board {
		for x, value := range col {
			if ViewGrid {
				vector.DrawFilledRect(screen, float32(x*SquareSize)+1+GridWidthOffset, float32(y*SquareSize)+1+GridHeightOffset, SquareSize-2, SquareSize-2, Colors[value], false)
			} else {
				vector.DrawFilledRect(screen, float32(x*SquareSize)+GridWidthOffset, float32(y*SquareSize)+GridHeightOffset, SquareSize, SquareSize, Colors[value], false)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Snake")
	Board[ScreenHeight/SquareSize/2][ScreenWidth/SquareSize/2] = 1

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
