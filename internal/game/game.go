package game

import (
	"math/rand"
	"time"

	"github.com/somewhat9/snake/internal/config"
	"golang.org/x/image/font"
)

type Game struct{
	Snake
	Board [][]uint8
	LastTick time.Time
	Status bool
	HighScore uint
	Cfg *config.Config
	FontFace font.Face
}

type Snake struct{
	Body []Position
	Dir Position
}

type Position struct {
	X, Y float32
}

func (g *Game) Tick() {
	head := g.Body[0]
	if !(g.Dir.X == 0 && g.Dir.Y == 0) {
		newHead := Position{X: head.X + g.Dir.X, Y: head.Y + g.Dir.Y}
		
		if g.Cfg.GridSquaresHeight <= int(newHead.Y) || newHead.Y < 0 || g.Cfg.GridSquaresWidth <= int(newHead.X) || newHead.X < 0 {
			g.Status = false
			return
		}

		positionValue := &g.Board[int(newHead.Y)][int(newHead.X)]

		/* Debug Prints
		fmt.Print(newHead)
		fmt.Print(" ")
		fmt.Print(head)
		fmt.Print(" ")
		fmt.Println(positionValue)
		*/

		switch *positionValue {
		case 0:
			g.Body = append([]Position{newHead}, g.Body...)
			tail := g.Body[len(g.Body)-1]
			g.Body = g.Body[:len(g.Body)-1]
			*positionValue = 1
			g.Board[int(tail.Y)][int(tail.X)] = 0
		case 1:
			// end game
			g.Status = false
		case 2:
			g.Body = append([]Position{newHead}, g.Body...)
			*positionValue = 1
			g.PlaceApple()
			g.UpdateHighScore()
		}
	}
}

func (g *Game) RandomEmpty() (y, x int) {
	var coords [][2]int
	for y, row := range g.Board {
		for x, value := range row {
			if value == 0 {
				coords = append(coords, [2]int{y, x})
			}
		}
	}
	
	if len(coords) == 0 {
		// end game
		g.Status = false
	}

	c := coords[rand.Intn(len(coords))]
	return c[0], c[1]
}

func (g *Game) PlaceApple() {
	y, x := g.RandomEmpty()
	g.Board[y][x] = 2
}

func (g *Game) UpdateHighScore() {
	if g.HighScore < uint(len(g.Body)-1) {
		g.HighScore = uint(len(g.Body)-1)
	}
}

func (g *Game) Setup() {
	g.Board = make([][]uint8, g.Cfg.GridSquaresHeight)
	for y := range g.Board {
		g.Board[y] = make([]uint8, g.Cfg.GridSquaresWidth)
	}
	startY, startX := g.RandomEmpty()
	g.Body = []Position{Position{X: float32(startX), Y: float32(startY)}}
	g.Board[startY][startX] = 1
	g.PlaceApple()
	g.Status = true
}
