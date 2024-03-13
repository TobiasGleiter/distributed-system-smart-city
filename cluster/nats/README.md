# NATS Cluster

## Overview

- Nats Clients connect to `nats://0.0.0.0:4222`
- Nats Server connection on port `<other-server-ip>:6222`

## Steps

1. Download and Install `nats-server` on each Server (and test)
2. Setup configuration files
3. Run all

### 1. Download and Install Nats Server

NATS server is highly optimized and its binary is very compact (less than 20 MB) and perfect for a Raspberry Pi.

[Release Build with curl](https://docs.nats.io/running-a-nats-service/introduction/installation#downloading-a-release-build)

### 2. Setup configuration files

Server 1, `nats-server1.conf`:

```
cluster {
  name: server-1

  listen: localhost:4244

  routes = [
    nats://127.0.0.1:4245
    nats://127.0.0.1:4246
  ]
}
```

Server 2, `nats-server2.conf`:

```
cluster {
  name: server-2

  listen: localhost:4245

  routes = [
    nats://127.0.0.1:4244
    nats://127.0.0.1:4246
  ]
}
```

Server 3, `nats-server3.conf`:

```
cluster {
  name: server-3

  listen: localhost:4246

  routes = [
    nats://127.0.0.1:4245
    nats://127.0.0.1:4244
  ]
}
```

### 3. Run all

Start on each server the cluster: `nats-server -config ./nats-serverX.conf -D`
