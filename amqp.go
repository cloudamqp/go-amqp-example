// Package example publishes and consumes messages with amqp-client.go.
package example

import (
	"context"
	"os"

	amqp "github.com/cloudamqp/amqp-client.go"
)

// Connect connects to CLOUDAMQP_URL, and reconnects if the connection is lost.
func Connect(ctx context.Context) (*amqp.Client, error) {
	url := os.Getenv("CLOUDAMQP_URL")
	if url == "" {
		url = "amqp://guest:guest@localhost"
	}
	return amqp.NewClient(ctx, url, nil)
}
