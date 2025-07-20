package config

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
