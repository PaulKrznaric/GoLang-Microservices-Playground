package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {

	conn := ConnectToConnection()
	defer conn.Close()

	fmt.Println("Sucessfully connected to RabitMQ Instance")

	ch := ConnectToChannel(conn)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	CheckForError(err)

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)
	CheckForError(err)

	fmt.Println("Sucessfully published message to queue")
}
