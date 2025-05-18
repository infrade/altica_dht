package main

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
)

func generate_key() (crypto.PrivKey, error) {
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, -1, rand.Reader)
	if err != nil {
		panic(err)
	}
	id, _ := crypto.MarshalPrivateKey(priv)
	_ = os.WriteFile("peer.key", id, 0600)

	fmt.Printf("Private key written to peer.key\n")
	return priv, nil
}
