package main

import (
	"fmt"
	"time"

	"github/go-redis-connectivity/db"
	"github/go-redis-connectivity/jsonhandling"
)

var (
	RedisAddr = "localhost:6379"

	message = jsonhandling.Message{
		Sender:      "Bob",
		Receiver:    "Alice",
		MessageBody: "Welcome to Redis in-memory caching.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
)

func main() {

	database, err := db.NewClient(RedisAddr)
	db.Must(err)

	db.Set(database, "creator", "Mustafa", 0)
	fmt.Println(db.Get(database, "creator"))

	db.Set(database, "message", jsonhandling.JsonMarshal(message), 0)

	var msg jsonhandling.Message
	data := db.Get(database, "message")
	current := jsonhandling.JsonUnmarshal(data, msg)
	fmt.Println(string(current.MessageBody))

	//db.Del(database, "message")

}
