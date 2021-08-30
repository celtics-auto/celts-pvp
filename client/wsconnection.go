package client

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Connection interface {
	Write(u *UpdateJson)
	Read(u *UpdateJson)
	Close() error
}

type WsConnection struct {
	Connection *websocket.Conn
}

func NewWsConnection(host, path string) (*WsConnection, error) {
	u := url.URL{Scheme: "ws", Host: host, Path: path}

	log.Printf("connecting to %s", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		return nil, err
	}

	return &WsConnection{
		Connection: conn,
	}, nil
}

func (w *WsConnection) Read(u *UpdateJson) {
	log.Println("Waiting messages...")
	err := w.Connection.ReadJSON(&u)
	if err != nil {
		log.Println("read:", err)
	}
}

func (w *WsConnection) Write(u *UpdateJson) {
	err := w.Connection.WriteJSON(u)
	if err != nil {
		log.Println("write:", err)
	}
}

func (w *WsConnection) Close() error {
	err := w.Connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}
	// FIXME: tá ruim
	time.Sleep(1 * time.Second)
	return nil
}
