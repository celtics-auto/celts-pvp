package client

import (
	"github.com/celtics-auto/ebiten-chat/config"
)

type Vector struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Player struct {
	Position  Vector `json:"position"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Animation int    `json:"animation"`
	Face      string `json:"face"`
}

type Message struct {
	Address string `json:"address"`
	Text    []byte `json:"text"`
}

type UpdateJson struct {
	Message *Message `json:"message"`
	Player  *Player  `json:"player"`
}

type Client struct {
	Conn     Connection
	Receiver chan *UpdateJson
	Sender   chan *UpdateJson
}

func New(cfg *config.Client, env string) *Client {
	receiver := make(chan *UpdateJson)
	sender := make(chan *UpdateJson)

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
		u := &UpdateJson{}
		c.Conn.Read(u)
		c.Receiver <- u
	}
}

func (c *Client) SendUpdates() {
	for {
		uJson := <-c.Sender
		c.Conn.Write(uJson)
	}
}
