package main

import (
	"log"

	"github.com/nats-io/nats"
)

func main() {
	log.Println("Starting Drone API")

	nc, err := nats.Connect("nats://192.168.1.113:4222")
	if err != nil {
		log.Fatal("Unable to connect to nats")
	}

	logMessages(nc)

	// do not exit
	for {
	}
}

func logMessages(nc *nats.Conn) {
	nc.Subscribe("log", func(m *nats.Msg) {
		log.Println(string(m.Data))
	})
}
