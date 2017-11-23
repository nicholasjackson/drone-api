package main

import (
	"log"

	"github.com/nats-io/nats"
	messages "github.com/nicholasjackson/drone-messages"
)

var nc *nats.Conn

func main() {
	log.Println("Starting Drone API")

	var err error
	nc, err = nats.Connect("nats://192.168.1.113:4222")
	if err != nil {
		log.Fatal("Unable to connect to nats")
	}

	logMessages(nc)
	logTweets(nc)
	logImages(nc)

	startServer()
}

func logMessages(nc *nats.Conn) {
	nc.Subscribe("log", func(m *nats.Msg) {
		log.Println(string(m.Data))
	})
}

func logImages(nc *nats.Conn) {
	nc.Subscribe(messages.MessageDroneImage, func(m *nats.Msg) {
		di := messages.DroneImage{}
		di.DecodeMessage(m.Data)
		di.SaveDataToFile("./latest.jpg")

		log.Println("Written latest image to ./latest.jpg")
	})
}

func logTweets(nc *nats.Conn) {
	nc.Subscribe("tweet", func(m *nats.Msg) {
		log.Println(string(m.Data))
	})
}
