package server

import (
	"fmt"
	"io"
	"log"

	"github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

// Client client structure
type Client struct {
	uc     uuid.UUID
	ws     *websocket.Conn
	server *Server
	msgCh  chan *Message
	doneCh chan struct{}
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
	msgCh := make(chan *Message)
	doneCh := make(chan struct{})

	return &Client{
		uc,
		ws,
		server,
		msgCh,
		doneCh,
	}

}

// Conn websocket connection to client
func (c *Client) Conn() *websocket.Conn {
	return c.ws
}

func (c *Client) Write(msg *Message) {
	select {
	case c.msgCh <- msg:
	default:
		c.server.Del(c)
		err := fmt.Errorf("client %s is disconnected", c.uc.String())
		c.server.Err(err)
	}
}

// Done terminate client
func (c *Client) Done() {
	c.doneCh <- struct{}{}
}

// Listen Write and Read request via chanel
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {

		// send message to the client
		case msg := <-c.msgCh:
			log.Println("Send:", msg)
			websocket.JSON.Send(c.ws, msg)

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- struct{}{} // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (c *Client) listenRead() {
	log.Println("Listening read from client")
	for {
		select {

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- struct{}{} // for listenWrite method
			return

		// read data from websocket connection
		default:
			var msg Message
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				c.doneCh <- struct{}{}
			} else if err != nil {
				c.server.Err(err)
			} else {
				c.server.SendAll(&msg)
			}
		}
	}
}
