package config

import "image/color"

var Colors = map[uint8]color.Color{
	0: color.Black, 
	1: color.RGBA{0, 255, 0, 255}, 
	2: color.RGBA{255, 0, 0, 255},
}