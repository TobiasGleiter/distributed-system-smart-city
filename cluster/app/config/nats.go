package config

import "github.com/nats-io/nats.go"


func ConnectToNatsClient()( *nats.Conn, error) {
    nc, err := nats.Connect("demo.nats.io")
    return nc, err
}