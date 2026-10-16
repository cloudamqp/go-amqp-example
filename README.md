# Use CloudAMQP from Go

A minimal example of publishing and consuming messages with
[amqp-client.go](https://github.com/cloudamqp/amqp-client.go), against
[CloudAMQP](https://www.cloudamqp.com) or any AMQP 0-9-1 broker such as
[LavinMQ](https://lavinmq.com) or RabbitMQ.

- [`publisher.go`](publisher.go) publishes a persistent message and returns once the broker has confirmed it.
- [`worker.go`](worker.go) consumes with a prefetch of 10, acking each message when the handler returns nil.
- [`main.go`](main.go) connects using the `CLOUDAMQP_URL` environment variable, reconnects automatically, and runs the publisher or worker.

## Usage

Tool versions are in `mise.toml`; install them with [mise](https://mise.jdx.dev),
or use Go 1.24+.

```sh
export CLOUDAMQP_URL="amqps://user:password@host/vhost"
go run . worker    # in one terminal
go run . publish   # in another
```

Without `CLOUDAMQP_URL` it connects to `amqp://guest:guest@localhost`.

## Tests

`go test ./...` publishes and consumes a message through a real broker on
localhost. CI runs it against LavinMQ.
