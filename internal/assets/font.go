package assets

import (
	_ "embed"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed Bitcount-Regular.ttf
var fontBytes []byte

func LoadFont(size float64) font.Face {
	tt, err := opentype.Parse(fontBytes); 
	if err != nil {
		log.Fatal(err)
	}
	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size: size,
		DPI: 72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return face
}
