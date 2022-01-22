package main

type Message struct {
	Content string
	id      int
}

type User struct {
	Received []Message
	Sent     []Message
	Name     string
}
