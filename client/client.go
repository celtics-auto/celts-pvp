package client

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
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
	Conn *websocket.Conn
}

func (c *Client) connect() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/connection"}
	log.Printf("connecting to %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	return conn
}

func (c *Client) ReceiveUpdates(receiver chan UpdateJson) {
	log.Println("Waiting messages...")
	for {
		u := UpdateJson{}
		err := c.Conn.ReadJSON(&u)
		if err != nil {
			log.Println("read:", err)
		}
		receiver <- u
	}
}

func (c *Client) SendUpdates(sender chan UpdateJson) {
	for {
		uJson := <-sender

		err := c.Conn.WriteJSON(uJson)
		if err != nil {
			log.Println("write:", err)
		}

	}
}

func (c *Client) CloseConnection() error {
	err := c.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}
	// FIXME: tá ruim
	time.Sleep(1 * time.Second)
	return nil
}

func New(env string) *Client {
	c := &Client{}

	if env != "development" {
		conn := c.connect()
		c.Conn = conn
	}

	return c
}
