package client

import (
	"fmt"
	"log"
	"net/url"

	"github.com/celtics-auto/ebiten-chat/objects"
	"github.com/gorilla/websocket"
)

type vector struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type player struct {
	Position vector `json:"position"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

type MessageJson struct {
	Address string `json:"address"`
	Message []byte `json:"message"`
}

type Client struct {
	Conn    *websocket.Conn
	Message *MessageJson
}

func (c *Client) connect() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/connection"}
	log.Printf("connecting to %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	return conn
}

func (c *Client) ReceiveMessage(msgChan chan *MessageJson) {
	fmt.Println("Waiting messages...")
	for {
		err := c.Conn.ReadJSON(c.Message)
		if err != nil {
			fmt.Println("read:", err)
		}
		msgChan <- c.Message
	}
}

func (c *Client) SendMessage(p *objects.Player) error {
	playerJson := player{
		Position: vector(*p.Position),
		Width:    p.Width,
		Height:   p.Height,
	}
	err := c.Conn.WriteJSON(playerJson)

	return err
}

func New() *Client {
	m := &MessageJson{}
	c := &Client{
		Message: m,
	}
	conn := c.connect()
	c.Conn = conn

	return c
}
