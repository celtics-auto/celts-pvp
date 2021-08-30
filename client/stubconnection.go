package client

import (
	"log"
)

type StubConnection struct {
	server chan *UpdateJson
}

func NewStubConnection() *StubConnection {
	return &StubConnection{
		server: make(chan *UpdateJson),
	}
}

func (s *StubConnection) Read(u *UpdateJson) {}

func (s *StubConnection) Write(u *UpdateJson) {}

func (s *StubConnection) Close() error {
	log.Println("closing connection")
	return nil
}
