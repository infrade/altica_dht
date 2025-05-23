package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/crypto"
	identify "github.com/libp2p/go-libp2p/p2p/protocol/identify"
)

func loadPrivKey(path string) (crypto.PrivKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return crypto.UnmarshalPrivateKey(data)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load identity or generate a new one
	var priv crypto.PrivKey
	var err error

	if _, err := os.Stat("peer.key"); os.IsNotExist(err) {
		// Generate a new key if it doesn't exist
		priv, err = generate_key()
		if err != nil {
			log.Fatalf("Failed to generate identity: %v", err)
		}
		log.Printf("Generated new identity")
	} else {
		priv, err = loadPrivKey("peer.key")
		if err != nil {
			log.Fatalf("Failed to load identity: %v", err)
		}
	}
	// Create host
	h, err := libp2p.New(
		libp2p.Identity(priv),
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/4001",
			"/ip6/::/tcp/4001",
		),
	)
	if err != nil {
		log.Fatalf("Failed to create host: %v", err)
	}

	log.Printf("✅ Bootstrap node started: %s", h.ID())
	for _, addr := range h.Addrs() {
		log.Printf(" - %s/p2p/%s", addr, h.ID())
	}

	// Ensure Identify service is active
	identify.NewIDService(h)

	// Start DHT in server mode
	kademliaDHT, err := dht.New(ctx, h, dht.Mode(dht.ModeServer), dht.ProtocolPrefix("/altica"))
	if err != nil {
		log.Fatalf("Failed to start DHT: %v", err)
	}
	if err := kademliaDHT.Bootstrap(ctx); err != nil {
		log.Fatalf("DHT bootstrap error: %v", err)
	}

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down...")
	if err := h.Close(); err != nil {
		log.Printf("Host shutdown error: %v", err)
	}
}
