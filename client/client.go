package client

import (
	"github.com/celtics-auto/ebiten-chat/config"
	"github.com/celtics-auto/ebiten-chat/objects"
)

type Message struct {
	Address string
	Text    []byte
}

type Client struct {
	Conn     Connection
	Receiver chan *objects.Player
	Sender   chan *objects.Player
}

func New(cfg *config.Client, env string) *Client {
	receiver := make(chan *objects.Player)
	sender := make(chan *objects.Player)

	c := &Client{
		Receiver: receiver,
		Sender:   sender,
	}

	switch env {
	case "development":
		c.Conn = NewStubConnection()
	default:
		c.Conn, _ = NewWsConnection(cfg.Host, cfg.Path)
	}

	return c
}

func (c *Client) ReceiveUpdates() {
	for {
		p := &objects.Player{}
		c.Conn.Read(p)
		c.Receiver <- p
	}
}

func (c *Client) SendUpdates() {
	for {
		p := <-c.Sender
		c.Conn.Write(p)
	}
}
