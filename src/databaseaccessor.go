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

var db *gorm.DB
var err error

func setup() {
	router, db := ConnectToDatabase()
	defer db.Close()

	router.HandleFunc("/messages", GetMessages).Methods("GET")
	router.HandleFunc("/messages/{id}", GetMessage).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/messages/{id}", DeleteMessage).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

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
