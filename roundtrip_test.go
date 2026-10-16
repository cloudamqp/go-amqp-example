package example

import (
	"context"
	"crypto/rand"
	"testing"
	"time"

	amqp "github.com/cloudamqp/amqp-client.go"
)

func TestWorkerReceivesPublishedMessage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	queue := "roundtrip-" + rand.Text()

	c, err := Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	received := make(chan string, 1)
	sub, err := Consume(ctx, c, queue, func(message string) error {
		received <- message
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Cancel(ctx)

	if err := Publish(ctx, c, queue, "Hello CloudAMQP!"); err != nil {
		t.Fatal(err)
	}

	select {
	case got := <-received:
		if got != "Hello CloudAMQP!" {
			t.Errorf("got %q, want %q", got, "Hello CloudAMQP!")
		}
	case <-ctx.Done():
		t.Fatal("no message received")
	}

	q, err := c.Queue(ctx, queue, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := q.Delete(ctx, amqp.QueueDeleteOptions{}); err != nil {
		t.Fatal(err)
	}
}
