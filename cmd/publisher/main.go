package main

import (
	"context"
	"log"

	example "github.com/cloudamqp/go-amqp-example"
)

func main() {
	ctx := context.Background()
	c, err := example.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	if err := example.Publish(ctx, c, "hello", "Hello CloudAMQP!"); err != nil {
		log.Fatal(err)
	}
	log.Println(" [x] Sent 'Hello CloudAMQP!'")
}
