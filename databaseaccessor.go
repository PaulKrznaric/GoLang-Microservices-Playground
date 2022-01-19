package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Message struct {
	Content string
	id      int
}

type User struct {
	Received []Message
	Sent     []Message
	Name     string
}

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=db port=5432 user=paulkrznaric dbname=gotest sslmode=disable password=pw")
	CheckForError(err)
	defer db.Close()

	addData(db)

	router.HandleFunc("/messages", GetMessages).Methods("GET")
	router.HandleFunc("/messages/{id}", GetMessage).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/messages/{id}", DeleteMessage).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

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

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	db.Find(&messages)
	json.NewEncoder(w).Encode(&messages)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var message Message
	db.First(&message, params["id"])
	json.NewEncoder(w).Encode(&message)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	var messages []Message
	db.First(&user, params["id"])
	db.Model(&user).Related(&messages)
	user.Sent = messages
	json.NewEncoder(w).Encode(&user)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var message Message
	db.First(&message, params["id"])
	db.Delete(&message)

	var messages []Message
	db.Find(&messages)
	json.NewEncoder(w).Encode(&messages)
}
