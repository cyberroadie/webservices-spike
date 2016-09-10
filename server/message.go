package server

import (
	"github.com/satori/go.uuid"
)

// Message unique id and text
type Message struct {
	Mid  uuid.UUID `json:"uuid"`
	Text string    `json:"text"`
}

func (m *Message) String() string {
	return m.Mid.String() + " : '" + m.Text + "'"
}
