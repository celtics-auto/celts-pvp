package game

import (
	"log"

	"github.com/celtics-auto/ebiten-chat/chat"
	"github.com/celtics-auto/ebiten-chat/client"
	"github.com/celtics-auto/ebiten-chat/config"
	"github.com/celtics-auto/ebiten-chat/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	config   *config.Config
	client   *client.Client
	Receiver chan objects.Player
	Sender   chan objects.Player
	player   *objects.Player
	chat     *chat.Chat
	count    int
}

func New(cfg *config.Config, c *client.Client, player *objects.Player, chat *chat.Chat) *Game {
	return &Game{
		config: cfg,
		client: c,
		player: player,
		chat:   chat,
	}
}

func (g *Game) Update() error {
	select {
	case p := <-g.client.Receiver:
		log.Printf("x: %d - y: %d\n", p.Position.X, p.Position.Y)
	default:
	}

	if g.player.Update() {
		g.client.Sender <- g.player
	}
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.chat.Draw(screen, g.config.Screen.Height)
	g.player.Draw(screen, g.count)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.config.Screen.Width, g.config.Screen.Height
}
