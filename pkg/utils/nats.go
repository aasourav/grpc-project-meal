package utils

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

var NatsConn *nats.Conn

func init() {
	natsClient, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("nats not connected")
	}
	NatsConn = natsClient
}
