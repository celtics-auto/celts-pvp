package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	msgType, msg, err := w.Connection.ReadMessage()
	if err != nil || msgType != websocket.BinaryMessage {
		log.Println("read:", err)
	}

	v := Vector{}
	buf := bytes.NewReader(msg)
	if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
		fmt.Printf("failed to decode byte array: %v", err)
		return
	}

	u.Player = &Player{
		Position: v,
	}
	fmt.Printf("x: %d - y: %d", u.Player.Position.X, u.Player.Position.Y)
}

func (w *WsConnection) Write(u *UpdateJson) {
	posX := u.Player.Position.X
	posY := u.Player.Position.Y
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint16(buf[0:], posX)
	binary.LittleEndian.PutUint16(buf[2:], posY)

	err := w.Connection.WriteMessage(websocket.BinaryMessage, buf)
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
