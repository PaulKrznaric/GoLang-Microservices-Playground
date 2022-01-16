package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkForError(err)
	defer conn.Close()

	fmt.Println("Sucessfully connected to RabitMQ Instance")

	ch, err := conn.Channel()
	checkForError(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	checkForError(err)

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
	checkForError(err)

	fmt.Println("Sucessfully published message to queue")
}

func checkForError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
}
