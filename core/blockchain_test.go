package core

import (
	"testing"
)

func TestBlockchainValidity(t *testing.T) {
	blockchain := CreateBlockchain(2)

	blockchain.AddBlock("Alice", "Bob", 5.0)
	blockchain.AddBlock("John", "Bob", 2.0)

	if !blockchain.IsValid() {
		t.Error("The blockchain must be valid using the correct data and mining process, but doesn't")
	}
}

func TestBlockchainInvalidity(t *testing.T) {
	blockchain := CreateBlockchain(2)

	blockchain.AddBlock("Alice", "Bob", 5.0)
	blockchain.AddBlock("John", "Bob", 2.0)
	blockchain.chain[2].hash = "INVALID_HASH"

	if blockchain.IsValid() {
		t.Error("The blockchain must be invalid for the given hash, but doesn't")
	}
}
