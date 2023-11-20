package main

import (
	"log"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/api"
	cr "github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
)

const (
	ListenAddress = ":8000"
	// TODO: add further configuration parameters here ...
)

func main() {
	storage := persistence.NewInMemoryStorage()
	signer := cr.NewSigner()
	server := api.NewServer(ListenAddress, storage, signer)

	if err := server.Run(); err != nil {
		log.Fatal("Could not start server on ", ListenAddress)
	}
}
