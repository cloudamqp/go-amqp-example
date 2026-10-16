package main

import (
	"context"
	"log"

	amqp "github.com/cloudamqp/amqp-client.go"
)

// runPublisher publishes a persistent message and returns once the broker has confirmed it.
func runPublisher(ctx context.Context, c *amqp.Client, queue, message string) error {
	q, err := c.Queue(ctx, queue, nil)
	if err != nil {
		return err
	}
	if err := q.Publish(ctx, amqp.Message{Body: message}); err != nil {
		return err
	}
	log.Printf(" [x] Sent '%s'", message)
	return nil
}
