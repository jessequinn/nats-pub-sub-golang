package main

import (
	"fmt"
	"log"

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
	// Simple Publisher
	for i := 0; i < 1000000; i++ {
		if err := nc.Publish("foo", []byte(fmt.Sprintf("[%d] Hello World", i))); err == nil {
			log.Println("Message published")
		}
	}
	defer nc.Close()
}
