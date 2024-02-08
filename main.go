package main

import (
	"communications/telegram"
	"communications/slack"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/telegram", telegram.Send).Methods("POST")
	router.HandleFunc("/slack", slack.Send).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}