package client

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct{}

func (c *Client) Connect() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return conn
}

func New() *Client {
	return &Client{}
}
