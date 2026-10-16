package example

import (
	"context"

	amqp "github.com/cloudamqp/amqp-client.go"
)

// Consume calls handler for each message, with a prefetch of 10. The message
// is acked when handler returns nil, and requeued when it returns an error.
func Consume(ctx context.Context, c *amqp.Client, queue string, handler func(message string) error) (*amqp.Subscription, error) {
	q, err := c.Queue(ctx, queue, nil)
	if err != nil {
		return nil, err
	}
	return q.Subscribe(ctx, func(ctx context.Context, d *amqp.Delivery) error {
		return handler(string(d.Body))
	}, &amqp.SubscribeOptions{Prefetch: 10})
}
