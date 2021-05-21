package model

import "time"

type Message struct {
	Type      string `json:type`
	Message   string `json:message`
	CreatedAt int64  `json:createdAt`
}

func NewMessage(messageType string, message string) *Message {
	return &Message{
		Type:      messageType,
		Message:   message,
		CreatedAt: time.Now().Unix(),
	}
}

type Messages []Message
