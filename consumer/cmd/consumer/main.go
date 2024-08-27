package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

// const ntsUrl = "127.0.0.1:4222"

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe
	sub1, err := nc.Subscribe("updates", func(m *nats.Msg) {
		log.Printf("sub1 Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set limits of 1000 messages or 5MB, whichever comes first
	sub1.SetPendingLimits(1000, 5*1024*1024)

	// Subscribe
	sub2, err := nc.Subscribe("updates", func(m *nats.Msg) {
		log.Printf("sub2 Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set no limits for this subscription
	sub2.SetPendingLimits(-1, -1)

	// Close the connection
	nc.Close()

	select {}
}
