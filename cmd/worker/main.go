package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	example "github.com/cloudamqp/go-amqp-example"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	c, err := example.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	_, err = example.Consume(ctx, c, "hello", func(message string) error {
		log.Printf(" [x] Received '%s'", message)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-ctx.Done()
}
