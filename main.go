package main

import (
	"fmt"
	"image/color"

	"log"

	"github.com/celtics-auto/ebiten-chat/chat"
	"github.com/celtics-auto/ebiten-chat/client"
	"github.com/celtics-auto/ebiten-chat/config"
	"github.com/celtics-auto/ebiten-chat/objects"
	"github.com/celtics-auto/ebiten-chat/utils"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	SCREEN_WIDTH  = 1366
	SCREEN_HEIGHT = 768
)

type Message struct {
	address string
	text    string
}

func (m *Message) String() string {
	return fmt.Sprintf("%s: %s", m.address, m.text)
}

type Game struct {
	text        string
	history     []Message
	config      *config.Config
	client      *client.Client
	messageChan chan *client.MessageJson
	player      *objects.Player
	chat        *chat.Chat
}

func newGame(cfg *config.Config, c *client.Client) *Game {
	msgChan := make(chan *client.MessageJson)

	return &Game{
		config:      cfg,
		history:     []Message{},
		client:      c,
		messageChan: msgChan,
	}
}

func (g *Game) Update() error {
	select {
	case mJson := <-g.messageChan:
		msgString := string(mJson.Message[:])
		g.history = append(g.history, Message{
			address: mJson.Address,
			text:    msgString,
		})
	default:
	}

	// if backspace was pressed
	//   delete one char from g.text

	g.text += string(ebiten.InputChars())

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if err := g.client.Conn.WriteMessage(websocket.TextMessage, []byte(g.text)); err != nil {
			log.Println("write error:", err)
			return err
		}
		g.text = ""
	}

	updatedPlayer := g.player.Update()
	if updatedPlayer != nil {
		g.client.SendMessage(updatedPlayer)
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
	g.player.Draw(screen)
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

	spriteSheet, _ := utils.NewSpriteSheet("./images/player.png")
	player := objects.NewPlayer(0, 0, spriteSheet)
	myGame.player = player

	go c.ReceiveMessage(myGame.messageChan)

	if err := ebiten.RunGame(myGame); err != nil {
		log.Fatal(err)
	}
}
