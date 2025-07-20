package game

import (
	"math/rand"
	"time"

	"github.com/somewhat9/snake/internal/config"
	"golang.org/x/image/font"
)

type Game struct{
	snake
	board [][]uint8
	lastTick time.Time
	status bool
	highScore uint
	Cfg *config.Config
	FontFace font.Face
}

type snake struct{
	body []position
	dir position
}

type position struct {
	X, Y float32
}

func (g *Game) tick() {
	head := g.body[0]
	if !(g.dir.X == 0 && g.dir.Y == 0) {
		newHead := position{X: head.X + g.dir.X, Y: head.Y + g.dir.Y}
		
		if g.Cfg.GridSquaresHeight <= int(newHead.Y) || newHead.Y < 0 || g.Cfg.GridSquaresWidth <= int(newHead.X) || newHead.X < 0 {
			g.status = false
			return
		}

		positionValue := &g.board[int(newHead.Y)][int(newHead.X)]

		switch *positionValue {
		case 0:
			g.body = append([]position{newHead}, g.body...)
			tail := g.body[len(g.body)-1]
			g.body = g.body[:len(g.body)-1]
			*positionValue = 1
			g.board[int(tail.Y)][int(tail.X)] = 0
		case 1:
			// end game
			g.status = false
		case 2:
			g.body = append([]position{newHead}, g.body...)
			*positionValue = 1
			g.placeApple()
			g.updatehighScore()
		}
	}
}

func (g *Game) randomEmpty() (y, x int) {
	var coords [][2]int
	for y, row := range g.board {
		for x, value := range row {
			if value == 0 {
				coords = append(coords, [2]int{y, x})
			}
		}
	}
	
	if len(coords) == 0 {
		// end game
		g.status = false
	}

	c := coords[rand.Intn(len(coords))]
	return c[0], c[1]
}

func (g *Game) placeApple() {
	y, x := g.randomEmpty()
	g.board[y][x] = 2
}

func (g *Game) updatehighScore() {
	if g.highScore < uint(len(g.body)-1) {
		g.highScore = uint(len(g.body)-1)
	}
}

func (g *Game) Setup() {
	g.board = make([][]uint8, g.Cfg.GridSquaresHeight)
	for y := range g.board {
		g.board[y] = make([]uint8, g.Cfg.GridSquaresWidth)
	}
	startY, startX := g.randomEmpty()
	g.body = []position{position{X: float32(startX), Y: float32(startY)}}
	g.board[startY][startX] = 1
	g.placeApple()
	g.status = true
}
