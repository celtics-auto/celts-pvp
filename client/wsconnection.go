package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/celtics-auto/ebiten-chat/objects"
	"github.com/celtics-auto/ebiten-chat/utils"
	"github.com/gorilla/websocket"
)

type Connection interface {
	Write(p *objects.Player)
	Read(p *objects.Player)
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

func (w *WsConnection) Read(p *objects.Player) {
	msgType, msg, err := w.Connection.ReadMessage()
	if err != nil || msgType != websocket.BinaryMessage {
		log.Println("failed to read websocket:", err)
	}

	v := utils.Vector{}
	buf := bytes.NewReader(msg)
	log.Println("read", msg)
	if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
		fmt.Printf("failed to decode byte array: %s\n", err.Error())
		return
	}

	// log.Printf("x: %d - y: %d\n", v.X, v.Y)
	p.Position = &v
}

func (w *WsConnection) Write(p *objects.Player) {
	posX := uint16(p.Position.X)
	posY := uint16(p.Position.Y)
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint16(buf[0:], posX)
	binary.LittleEndian.PutUint16(buf[2:], posY)
	log.Println("write", buf)
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
