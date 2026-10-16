package main

import (
	"context"
	"log"

	amqp "github.com/cloudamqp/amqp-client.go"
)

// runWorker calls handler for each message until ctx is done, with a prefetch of 10.
// The message is acked when handler returns nil, and requeued when it returns an error.
func runWorker(ctx context.Context, c *amqp.Client, queue string, handler func(message string) error) error {
	q, err := c.Queue(ctx, queue, nil)
	if err != nil {
		return err
	}
	_, err = q.Subscribe(ctx, func(ctx context.Context, d *amqp.Delivery) error {
		return handler(string(d.Body))
	}, &amqp.SubscribeOptions{Prefetch: 10})
	if err != nil {
		return err
	}
	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-ctx.Done()
	return nil
}
