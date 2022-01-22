package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

var routeName string = "amqp://guest:guest@localhost:5672/"

func ConnectToConnection() *amqp.Connection {
	connection, err := amqp.Dial(routeName)
	CheckForError(err)
	return connection
}

func ConnectToChannel(conn *amqp.Connection) (channel *amqp.Channel) {
	channel, err := conn.Channel()
	CheckForError(err)
	return channel
}

func CreateQueue(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	CheckForError(err)
	fmt.Print(q)
}

func ConnectToDatabase() (*mux.Router router, *gorm.db db){
	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=localhost port=54320 user=go_user dbname=gotest1 sslmode=disable password=pw")
	CheckForError(err)

	return router, db
}

func CheckForError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
}
