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
	Receiver chan client.UpdateJson
	Sender   chan client.UpdateJson
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
	case uJson := <-g.client.Receiver:
		if uJson.Message != nil {
			g.chat.ReceiveMessages(uJson.Message.Address, uJson.Message.Text)
		}
		if uJson.Player != nil {
			log.Println("Test")
		}

	default:
	}

	uJson := client.UpdateJson{}
	sendUpdate := false
	if g.player.Update() {
		uJson.Player = &client.Player{
			Position: client.Vector{
				X: g.player.Position.X,
				Y: g.player.Position.Y,
			},
			Width:     g.player.Width,
			Height:    g.player.Height,
			Animation: g.player.Animation,
			Face:      g.player.Face,
		}
		sendUpdate = true
	}

	if g.chat.Update() {
		message := g.chat.History[len(g.chat.History)-1]
		mString := []byte(message.Text)

		// FIXME: pegar endereço do client para colocar junto a mensagem
		uJson.Message = &client.Message{
			Text: mString,
		}
		sendUpdate = true
	}

	if sendUpdate {
		g.client.Sender <- &uJson
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
