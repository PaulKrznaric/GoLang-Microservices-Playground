package main

import "fmt"

func checkForError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
}
