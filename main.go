package main

import (
	"log"
	"messageServer/handlers"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	handler := handlers.NewMessageHandler()
	server.HandleFunc(handler.NewGetMessageHandler())
	server.HandleFunc(handler.NewPostMessageHandler())

	err := http.ListenAndServe("localhost:80", server)
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("leaving")
}
