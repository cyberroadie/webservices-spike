package server

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

// Server server
type Server struct {
	path     string
	messages []*Message
	clients  map[uuid.UUID]*Client
	addCh    chan *Client
	delCh    chan *Client
	msgCh    chan *Message
	finCh    chan struct{}
	errCh    chan error
	doneCh   chan struct{}
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
		make(chan struct{}),
	}
}

// Add add client to server
func (s *Server) Add(c *Client) {
	s.addCh <- c
}

// Del delete client from server
func (s *Server) Del(c *Client) {
	s.delCh <- c
}

// Err send to error channel
func (s *Server) Err(e error) {
	s.errCh <- e
}

// Done stops server
func (s *Server) Done() {
	s.doneCh <- struct{}{}
}

// SendAll send message to all clients
func (s *Server) SendAll(msg *Message) {
	s.msgCh <- msg
}

func (s *Server) sendPastMessages(c *Client) {
	for _, msg := range s.messages {
		c.Write(msg)
	}
}

func (s *Server) sendAll(msg *Message) {
	for _, c := range s.clients {
		c.Write(msg)
	}
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {

	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s)
		s.Add(client)
		client.Listen()
	}
	http.Handle(s.path, websocket.Handler(onConnected))
	log.Println("Created handler")

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client")
			s.clients[c.uc] = c
			log.Println("Now", len(s.clients), "clients connected.")
			s.sendPastMessages(c)

		// del a client
		case c := <-s.delCh:
			log.Println("Delete client")
			delete(s.clients, c.uc)

		// broadcast message for all clients
		case msg := <-s.msgCh:
			log.Println("Send all:", msg)
			s.messages = append(s.messages, msg)
			s.sendAll(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}
