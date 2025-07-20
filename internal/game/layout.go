package game

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Cfg.ScreenWidth(), g.Cfg.ScreenHeight()
}