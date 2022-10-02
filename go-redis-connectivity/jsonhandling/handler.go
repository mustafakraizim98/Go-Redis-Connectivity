package jsonhandling

import (
	"encoding/json"
	"time"

	"github/go-redis-connectivity/db"
)

type Message struct {
	Sender      string    `json:"sender,omitempty" binding:"required"`
	Receiver    string    `json:"receiver,omitempty" binding:"required"`
	MessageBody string    `json:"message,omitempty" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func JsonMarshal(msg Message) []byte {
	json, err := json.Marshal(msg)
	db.Must(err)
	return json
}

func JsonUnmarshal(data string, msg Message) Message {
	err := json.Unmarshal([]byte(data), &msg)
	db.Must(err)
	return msg
}
