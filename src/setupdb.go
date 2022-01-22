package main

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	_, db := ConnectToDatabase()
	defer db.Close()

	addData(db)
}

func addData(db *gorm.DB) {

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Message{})

	users := []User{
		{Name: "Paul"},
		{Name: "Brody"},
		{Name: "Raj"},
	}

	messages := []Message{
		{Content: "Hello", id: 0},
		{Content: "Hey", id: 1},
		{Content: "How are you?", id: 2},
		{Content: "Swell. You?", id: 3},
		{Content: "Fantastic.", id: 4},
		{Content: "Are you there?", id: 5},
	}

	for index := range users {
		db.Create(&users[index])
	}

	for index := range messages {
		db.Create(&messages[index])
	}
}
