package main

import (
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}
	// Simple Async Subscriber
	if _, err := nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println(string(msg.Data))
	}); err != nil {
		log.Fatalf("Subscription failed: %v", err)
	}
	// Keep the connection alive
	runtime.Goexit()
}
