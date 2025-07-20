package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/somewhat9/snake/internal/assets"
	"github.com/somewhat9/snake/internal/config"
	"github.com/somewhat9/snake/internal/game"
)

func main() {
	gameInstance := &game.Game{}
	cfg, err := config.LoadYAML("settings")
	gameInstance.Cfg = cfg
	if err != nil {
		log.Fatalf("could not load .yaml: %v", err)
	}
	ebiten.SetWindowSize(gameInstance.Cfg.ScreenWidth(), gameInstance.Cfg.ScreenHeight())
	ebiten.SetWindowTitle("Snake")
	gameInstance.FontFace = assets.LoadFont(float64(gameInstance.Cfg.FontSize))

	gameInstance.Setup()

	if err := ebiten.RunGame(gameInstance); err != nil {
		log.Fatal(err)
	}
}
