package chat

import "github.com/celtics-auto/ebiten-chat/client"

type Message struct {
	address string
	text    string
}

type Chat struct {
	text        string
	history     []Message
	messageChan chan *client.MessageJson
}
