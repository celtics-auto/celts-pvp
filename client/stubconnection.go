package client

import (
	"log"

	"github.com/celtics-auto/ebiten-chat/objects"
)

type StubConnection struct {
	server chan *objects.Player
}

func NewStubConnection() *StubConnection {
	return &StubConnection{
		server: make(chan *objects.Player),
	}
}

func (s *StubConnection) Read(u *objects.Player) {}

func (s *StubConnection) Write(u *objects.Player) {}

func (s *StubConnection) Close() error {
	log.Println("closing connection")
	return nil
}
