package config

import (
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Config struct {
	Fonts  Fonts
	Client Client
}

type Client struct {
	address string
	port    string
}

type Fonts struct {
	MplusNormal font.Face
}

func New() (*Config, error) {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		return nil, err
	}

	mpn, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	f := Fonts{
		MplusNormal: mpn,
	}
	c := Client{
		address: "localhost",
		port:    "8080",
	}
	cfg := &Config{
		Fonts:  f,
		Client: c,
	}

	return cfg, nil
}
