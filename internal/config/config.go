package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

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

func LoadYAML(fileName string) (*Config, error) {
	cfg := Default()
	path := filepath.Join("configs", fileName+".yaml")

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, fmt.Errorf("read config %q: %w", path, err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parse config %q: %w", path, err)
	}

	return cfg, nil
}
