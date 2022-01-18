package main

import (
	"fmt"

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

func CheckForError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
}
