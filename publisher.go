package example

import (
	"context"

	amqp "github.com/cloudamqp/amqp-client.go"
)

// Publish publishes a persistent message and returns once the broker has confirmed it.
func Publish(ctx context.Context, c *amqp.Client, queue, message string) error {
	q, err := c.Queue(ctx, queue, nil)
	if err != nil {
		return err
	}
	return q.Publish(ctx, amqp.Message{Body: message})
}
