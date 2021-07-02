package main

import (
	"image/color"
	"log"

	"github.com/celtics-auto/ebiten-chat/client"
	"github.com/celtics-auto/ebiten-chat/config"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

type Game struct {
	text   string
	config *config.Config
	conn   *websocket.Conn
}

func newGame(cfg *config.Config, client *client.Client) *Game {
	conn := client.Connect()

	return &Game{
		config: cfg,
		conn:   conn,
	}
}

func (g *Game) Update() error {
	// if backspace was pressed
	//   delete one char from g.text

	g.text += string(ebiten.InputChars())

	if err := g.conn.WriteMessage(websocket.TextMessage, []byte(g.text)); err != nil {
		log.Println("write error:", err)
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, g.text, g.config.Fonts.MplusNormal, 0, 100, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("CHAT")

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	c := client.New()

	myGame := newGame(cfg, c)
	if err := ebiten.RunGame(myGame); err != nil {
		log.Fatal(err)
	}
}
