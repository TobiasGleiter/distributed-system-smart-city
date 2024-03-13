package config

import (
    "strings"

    "github.com/nats-io/nats.go"
)

func ConnectToNatsServer()( *nats.Conn, error) {
    servers := []string{"nats://0.0.0.0:4244", "nats://0.0.0.0:4245", "nats://0.0.0.0:4246"}
    nc, err := nats.Connect(strings.Join(servers, ","))
    return nc, err
}