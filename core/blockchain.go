package core

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

// PRIVATE METHODS

func (block Block) calculateHash() string {
	data, _ := json.Marshal(block.data)
	blockData := block.previousHash + string(data) + block.timestamp.String() + strconv.Itoa(block.proofOfWork)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (block *Block) mine(difficulty int) {
	for !strings.HasPrefix(block.hash, strings.Repeat("0", difficulty)) {
		block.proofOfWork++
		block.hash = block.calculateHash()
	}
}

// PUBLIC METHODS

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (blockchain *Blockchain) AddBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := blockchain.chain[len(blockchain.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(blockchain.difficulty)
	blockchain.chain = append(blockchain.chain, newBlock)
}

func (blockchain Blockchain) IsValid() bool {
	for i := range blockchain.chain[1:] {
		previousBlock := blockchain.chain[i]
		currentBlock := blockchain.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}
