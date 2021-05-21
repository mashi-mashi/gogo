package model

import "time"

type Message struct {
	Type      string `json:type`
	Message   string `json:message`
	Timestamp int64  `json:timestamp`
}

func NewMessage(messageType string, message string) *Message {
	return &Message{
		Type:      messageType,
		Message:   message,
		Timestamp: time.Now().Unix(),
	}
}

type Messages []Message
