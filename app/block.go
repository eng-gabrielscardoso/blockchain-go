package app

import "time"

// This struct represents a block from a blockchain
type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	proofOfWork  int
}
