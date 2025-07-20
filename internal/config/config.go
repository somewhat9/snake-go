package config

type Config struct {
	ViewGrid			bool	`yaml:"view_grid"`
	SquareSize          int		`yaml:"square_size"`
	GridSquaresWidth    int		`yaml:"grid_squares_width"`
	GridSquaresHeight	int		`yaml:"grid_squares_height"`
	FontSize			int		`yaml:"font_size"`
}

func(c *Config) GridWidth() int {
	return c.SquareSize * c.GridSquaresWidth
}

func (c *Config) GridHeight() int {
	return c.SquareSize * c.GridSquaresHeight
}

func (c *Config) GridWidthOffset() int {
	return 0
}

func (c *Config) GridHeightOffset() int {
	return c.SquareSize*2
}

func (c *Config) ScreenWidth() int {
	return c.GridWidth() + c.GridWidthOffset()
}

func (c *Config) ScreenHeight() int {
	return c.GridHeight() + c.GridHeightOffset()
}

func Default() *Config {
	return &Config {
		ViewGrid: false,
		SquareSize: 30,
		GridSquaresWidth: 40,
		GridSquaresHeight: 30,
		FontSize: 36,
	}
}
