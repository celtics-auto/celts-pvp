package main

import (
	"fmt"

	"log"

	"github.com/celtics-auto/ebiten-chat/chat"
	"github.com/celtics-auto/ebiten-chat/client"
	"github.com/celtics-auto/ebiten-chat/config"
	devutils "github.com/celtics-auto/ebiten-chat/dev"
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

type Game struct {
	config   *config.Config
	client   *client.Client
	receiver chan client.UpdateJson
	sender   chan client.UpdateJson
	player   *objects.Player
	chat     *chat.Chat
	devMode  bool
	count    int
}

func newGame(cfg *config.Config, c *client.Client, player *objects.Player, chat *chat.Chat) *Game {
	receiver := make(chan client.UpdateJson)
	sender := make(chan client.UpdateJson)

	return &Game{
		config:   cfg,
		client:   c,
		receiver: receiver,
		sender:   sender,
		player:   player,
		chat:     chat,
		devMode:  cfg.Env == "development",
	}
}

func (g *Game) Update() error {
	// TODO: use game state logic to check for changes

	select {
	case uJson := <-g.receiver:
		if uJson.Message != nil {
			g.chat.ReceiveMessages(uJson.Message.Address, uJson.Message.Text)
		}
		if uJson.Player != nil {
			log.Println("Test")
		}

	default:
	}

	// TODO: if backspace was pressed
	//   delete one char from g.text

	g.chat.Update(g.sender, g.devMode)
	g.player.Update(g.sender, g.devMode)

	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.chat.Draw(screen, g.config.Screen.Height)
	g.player.Draw(screen, g.count)

	if g.devMode {
		devutils.Draw(screen, g.player)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.config.Screen.Width, g.config.Screen.Height
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(cfg.Screen.Width, cfg.Screen.Height)
	ebiten.SetWindowTitle("CHAT")

	c := client.New(cfg.Env)
	spriteSheet, _ := utils.NewSpriteSheet("./images/genericPlayer_50x50.png", 50, 50)
	player := objects.NewPlayer(0, 0, spriteSheet)
	ch := &chat.Chat{
		Fonts: &cfg.Fonts,
	}
	myGame := newGame(cfg, c, player, ch)

	if !myGame.devMode {
		go c.ReceiveUpdates(myGame.receiver)
		go c.SendUpdates(myGame.sender)
	}

	if err := ebiten.RunGame(myGame); err != nil {
		log.Fatal(err)
	}

	closeErr := c.CloseConnection()
	if closeErr != nil {
		log.Fatal(closeErr)
	}
}
