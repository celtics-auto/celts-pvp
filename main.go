package main

import (
	"fmt"

	"log"

	"github.com/celtics-auto/ebiten-chat/chat"
	"github.com/celtics-auto/ebiten-chat/client"
	"github.com/celtics-auto/ebiten-chat/config"
	"github.com/celtics-auto/ebiten-chat/game"
	"github.com/celtics-auto/ebiten-chat/objects"
	"github.com/celtics-auto/ebiten-chat/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Message struct {
	address string
	text    string
}

func (m *Message) String() string {
	return fmt.Sprintf("%s: %s", m.address, m.text)
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(cfg.Screen.Width, cfg.Screen.Height)
	ebiten.SetWindowTitle("CHAT")

	c := client.New(&cfg.Client, cfg.Env)
	spriteSheet, _ := utils.NewSpriteSheet("./images/genericPlayer_50x50.png", 50, 50)
	player := objects.NewPlayer(0, 0, spriteSheet)
	ch := &chat.Chat{
		Fonts: &cfg.Fonts,
	}
	myGame := game.New(cfg, c, player, ch)

	go c.ReceiveUpdates()
	go c.SendUpdates()

	if err := ebiten.RunGame(myGame); err != nil {
		log.Fatal(err)
	}

	closeErr := c.Conn.Close()
	if closeErr != nil {
		log.Fatal(closeErr)
	}
}
