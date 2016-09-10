package server

import (
	"github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

// Client client structure
type Client struct {
	uc     uuid.UUID
	ws     *websocket.Conn
	server *Server
	ch     chan *Message
	done   chan struct{}
}

// NewClient construct Client
func NewClient(ws *websocket.Conn, server *Server) *Client {
	if ws == nil {
		panic("websocket can not be nil")
	}

	if server == nil {
		panic("server can not be nil")
	}

	uc := uuid.NewV1()
	ch := make(chan *Message)
	done := make(chan struct{})

	return &Client{
		uc,
		ws,
		server,
		ch,
		done,
	}

}
