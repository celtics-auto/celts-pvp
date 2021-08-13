package chat

import (
	"fmt"
	"image/color"

	"github.com/celtics-auto/ebiten-chat/client"
	"github.com/celtics-auto/ebiten-chat/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Message struct {
	address string
	text    string
}

type Chat struct {
	input   string
	history []Message
	Fonts   *config.Fonts
}

func (ch *Chat) ReceiveMessages(address string, text []byte) {
	msgString := string(text[:])
	ch.history = append(ch.history, Message{
		address: address,
		text:    msgString,
	})
}

func (ch *Chat) Update(sender chan client.UpdateJson, env string) {
	ch.input += string(ebiten.InputChars())

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		message := []byte(ch.input)
		ch.input = ""

		// FIXME: pegar endereço do client para colocar junto a mensagem
		uJson := &client.UpdateJson{
			Message: &client.Message{
				Text: message,
			},
		}

		if env != "development" {
			sender <- *uJson
		}
	}
}

func (ch *Chat) Draw(screen *ebiten.Image, screenHeight int) {
	if len(ch.history) > 0 {
		lineHeight := 60
		for i := len(ch.history) - 1; i >= 0; i-- {
			txt := fmt.Sprintf("%s: %s", ch.history[i].address, ch.history[i].text)
			text.Draw(screen, txt, ch.Fonts.MplusNormal, 10, screenHeight-lineHeight, color.White)
			lineHeight += 30
		}
	}
	text.Draw(screen, ch.input, ch.Fonts.MplusNormal, 10, screenHeight-20, color.White)
}
