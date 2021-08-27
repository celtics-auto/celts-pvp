package chat

import (
	"fmt"
	"image/color"

	"github.com/celtics-auto/ebiten-chat/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Message struct {
	Address string
	Text    string
}

type Chat struct {
	input   string
	History []Message
	Fonts   *config.Fonts
}

func (ch *Chat) ReceiveMessages(address string, text []byte) {
	msgString := string(text[:])
	ch.History = append(ch.History, Message{
		Address: address,
		Text:    msgString,
	})
}

func (ch *Chat) Update() bool {
	ch.input += string(ebiten.InputChars())

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		ch.History = append(ch.History, Message{
			Text: ch.input,
		})
		ch.input = ""

		return true
	}

	return false
}

func (ch *Chat) Draw(screen *ebiten.Image, screenHeight int) {
	if len(ch.History) > 0 {
		lineHeight := 60
		for i := len(ch.History) - 1; i >= 0; i-- {
			txt := fmt.Sprintf("%s: %s", ch.History[i].Address, ch.History[i].Text)
			text.Draw(screen, txt, ch.Fonts.MplusNormal, 10, screenHeight-lineHeight, color.White)
			lineHeight += 30
		}
	}
	text.Draw(screen, ch.input, ch.Fonts.MplusNormal, 10, screenHeight-20, color.White)
}
