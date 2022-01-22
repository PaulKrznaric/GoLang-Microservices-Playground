package main

import (
	"fmt"
)

func main() {

	conn := ConnectToConnection()
	defer conn.Close()

	fmt.Println("Sucessfully connected to RabitMQ Instance")

	ch := ConnectToChannel(conn)
	defer ch.Close()

	CreateQueue(ch)
}
