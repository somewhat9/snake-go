package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ViewGrid bool = false
	SquareSize = 30

	GridSquaresWidth = 40
	GridSquaresHeight = 30

	GridWidth = GridSquaresWidth*SquareSize
	GridHeight = GridSquaresHeight*SquareSize

	GridHeightOffset = SquareSize*2
	GridWidthOffset = SquareSize*0

	ScreenWidth = GridWidth + GridWidthOffset
	ScreenHeight = GridHeight + GridHeightOffset
)

//go:embed assets/Bitcount-Regular.ttf
var fontBytes []byte

var fontFace font.Face

func loadFont() font.Face {
	tt, err := opentype.Parse(fontBytes); 
	if err != nil {
		log.Fatal(err)
	}
	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 36,
		DPI: 72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return face
}

var Colors = map[uint8]color.Color{
	0: color.Black, 
	1: color.RGBA{0, 255, 0, 255}, 
	2: color.RGBA{255, 0, 0, 255},
}

type Game struct{
	Snake
	Board [GridSquaresHeight][GridSquaresWidth]uint8
	LastTick time.Time
}

type Snake struct{
	Body []Position
	Dir Position
}

func (g *Game) Tick() {
	head := g.Body[0]
	newHead := Position{X: head.X + g.Dir.X, Y: head.Y + g.Dir.Y}

	/* Debug Prints
	fmt.Print(newHead)
	fmt.Print(" ")
	fmt.Print(head)
	fmt.Print(" ")
	fmt.Println(g.Board[int(newHead.Y)][int(newHead.X)])
	*/

	switch g.Board[int(newHead.Y)][int(newHead.X)] {
	case 0:
		g.Body = append([]Position{newHead}, g.Body...)
		tail := g.Body[len(g.Body)-1]
		g.Body = g.Body[:len(g.Body)-1]
		g.Board[int(newHead.Y)][int(newHead.X)] = 1
		g.Board[int(tail.Y)][int(tail.X)] = 0
	case 1:
		// end game
	case 2:
		g.Body = append([]Position{newHead}, g.Body...)
		g.Board[int(newHead.Y)][int(newHead.X)] = 1
	}
}

type Position struct {
	X, Y float32
}

func (g *Game) Update() error {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
		g.Dir = Position{X: 0, Y: -1}
	case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
		g.Dir = Position{X: 0, Y: 1}
	case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
		g.Dir = Position{X: 1, Y: 0}
	case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
		g.Dir = Position{X: -1, Y: 0}
	}

	now := time.Now()
	if now.Sub(g.LastTick) >= time.Second/4 {
		g.LastTick = now
		g.Tick()
	}

	return nil 
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{30})
	for y, col := range g.Board {
		for x, value := range col {
			if ViewGrid {
				vector.DrawFilledRect(screen, float32(x*SquareSize)+1+GridWidthOffset, float32(y*SquareSize)+1+GridHeightOffset, SquareSize-2, SquareSize-2, Colors[value], false)
			} else {
				vector.DrawFilledRect(screen, float32(x*SquareSize)+GridWidthOffset, float32(y*SquareSize)+GridHeightOffset, SquareSize, SquareSize, Colors[value], false)
			}
		}
	}
	text.Draw(screen, "Score: " + fmt.Sprint(len(g.Body)-1), fontFace, SquareSize, SquareSize*1.5, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Snake")
	fontFace = loadFont()
	game.Body = append(game.Body, Position{X: GridWidth/SquareSize/2, Y: GridHeight/SquareSize/2})
	game.Board[GridHeight/SquareSize/2][GridWidth/SquareSize/2] = 1

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}