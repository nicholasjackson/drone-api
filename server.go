package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Message is the json message passed to the server
type Message struct {
	Command string
	Value   string
}

func startServer() {
	log.Println("Starting Control Server")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	})

	http.ListenAndServe(":8088", nil)
}

func apiHandler(rw http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		log.Println(err)
	}

	switch message.Command {
	case "LAUNCH":
		log.Println("Launching")
		nc.Publish("drone", []byte("LAUNCH"))

	case "LAND":
		log.Println("Landing")
		nc.Publish("drone", []byte("LAND"))

	case "PICTURE":
		log.Println("Taking picture")
		nc.Publish("drone", []byte("TAKEPICTURE"))
	}
}
