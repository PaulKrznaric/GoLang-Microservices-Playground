package main

import (
	"fmt"
)

func main() {
	conn := ConnectToConnection()
	defer conn.Close()

	ch := ConnectToChannel(conn)
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	CheckForError(err)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
		}
	}()

	fmt.Printf("Sucessfully connected to our RabitMQ instance")
	fmt.Println(" [*] - waiting for message")
	<-forever
}
