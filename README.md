# Altica DHT Bootstrap Node

A lightweight and efficient DHT (Distributed Hash Table) bootstrap node implementation using libp2p. This node serves as a bootstrapping point for other nodes in a libp2p network, helping new peers discover and connect to the network.

## Features

- Implements a Kademlia DHT in server mode
- Persistent peer identity across restarts
- IPv4 and IPv6 support
- Docker support
- Systemd service configuration
- Graceful shutdown handling

## Prerequisites

- Go 1.23.8 or later
- Docker (optional, for containerized deployment)

## Installation

Clone the repository:

```bash
git clone https://github.com/infrade/altica_dht.git
cd altica_dht
```

### Building from Source

```bash
make build
```

### Docker Build

```bash
make docker-build
```

## Running the Node

### Local Execution

```bash
make run
```

### Docker Container

```bash
make docker-run
```

The node will listen on port 4001 for both TCP and UDP connections.

### Systemd Service

1. Copy the binary to the system:

    ```bash
    sudo cp altica_dht /usr/local/bin/
    ```

1. Copy the systemd service file:

    ```bash
    sudo cp systemd/altica_dht.service /etc/systemd/system/
    ```

1. Create the altica_dht user and directory:

    ```bash
    sudo useradd -r altica_dht
    sudo mkdir -p /opt/altica_dht
    sudo chown altica_dht:altica_dht /opt/altica_dht
    ```

1. Enable and start the service:

    ```bash
    sudo systemctl daemon-reload
    sudo systemctl enable altica_dht
    sudo systemctl start altica_dht
    ```

## Configuration

The node automatically generates a peer identity on first run and stores it in `peer.key`. This identity is persistent across restarts. To skip this step, you can provide your own `peer.key` file.

Default listen addresses:

- `/ip4/0.0.0.0/tcp/4001`
- `/ip6/::/tcp/4001`

## Building

- `make build`: Build the binary
- `make docker-build`: Build Docker image
- `make docker-run`: Run in Docker container
- `make run`: Run locally

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
