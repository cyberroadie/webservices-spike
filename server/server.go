package server

import uuid "github.com/satori/go.uuid"

// Server server
type Server struct {
	path     string
	messages []*Message
	clients  map[uuid.UUID]*Client
	addC     chan *Client
	delC     chan *Client
	msgC     chan *Message
	finC     chan struct{}
	errC     chan error
}

// NewServer create new server
func NewServer(path string) *Server {
	return &Server{
		path,
		[]*Message{},
		make(map[uuid.UUID]*Client),
		make(chan *Client),
		make(chan *Client),
		make(chan *Message),
		make(chan struct{}),
		make(chan error),
	}
}
