package config

import (
	"io/ioutil"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Config struct {
	Fonts  Fonts
	Client Client
	Screen Screen
	Env    string
}

type Client struct {
	Address string
	Port    string
	Path    string
	Host    string
}

type Screen struct {
	Height int
	Width  int
}

type Fonts struct {
	MplusNormal font.Face
}

func New() (*Config, error) {
	runescapeFont, fileReadErr := ioutil.ReadFile("./fonts/runescape_uf.ttf")
	if fileReadErr != nil {
		return nil, fileReadErr
	}
	tt, err := opentype.Parse(runescapeFont)
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
		Path: "/connection",
		Host: "localhost:3000",
	}
	s := Screen{
		Height: 768,
		Width:  1366,
	}
	cfg := &Config{
		Fonts:  f,
		Client: c,
		Screen: s,
		Env:    "production",
	}

	return cfg, nil
}
