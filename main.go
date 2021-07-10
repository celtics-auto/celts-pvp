package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/celtics-auto/ebiten-chat/client"
	"github.com/celtics-auto/ebiten-chat/config"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

type Message struct {
	address string
	text    string
}

func (m *Message) String() string {
	return fmt.Sprintf("%s: %s", m.address, m.text)
}

type Game struct {
	text    string
	history []Message
	config  *config.Config
	conn    *websocket.Conn
}

func newGame(cfg *config.Config, client *client.Client) *Game {
	conn := client.Connect()

	return &Game{
		config:  cfg,
		history: []Message{},
		conn:    conn,
	}
}

func (g *Game) Update() error {
	// if backspace was pressed
	//   delete one char from g.text

	g.text += string(ebiten.InputChars())

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if err := g.conn.WriteMessage(websocket.TextMessage, []byte(g.text)); err != nil {
			log.Println("write error:", err)
			return err
		}

		g.history = append(g.history, Message{
			address: g.conn.LocalAddr().String(),
			text:    g.text,
		})
		g.text = ""
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if len(g.history) > 0 {
		lineHeight := 60
		for i := len(g.history) - 1; i >= 0; i-- {
			text.Draw(screen, g.history[i].String(), g.config.Fonts.MplusNormal, 10, SCREEN_HEIGHT-lineHeight, color.White)
			lineHeight += 30
		}
	}
	text.Draw(screen, g.text, g.config.Fonts.MplusNormal, 10, SCREEN_HEIGHT-20, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
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

	/*
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			sig := <-sigChan
			fmt.Println(sig, "EXITING")

			data := websocket.FormatCloseMessage(websocket.CloseNormalClosure, fmt.Sprintf("Client %s exiting.", myGame.conn.LocalAddr()))
			myGame.conn.WriteMessage(websocket.CloseMessage, data)
		}()
	*/
}
