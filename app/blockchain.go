package app

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

////////////////////////////////////////////////////////////////////////////////
// STRUCTS
////////////////////////////////////////////////////////////////////////////////

// This struct represents a blockchain
type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS
////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS
////////////////////////////////////////////////////////////////////////////////

// This function creates a new blockchain
func CreateBlockchain(ctx *cli.Context) {
	difficulty, error := strconv.Atoi(ctx.Args().Get(0))

	if error != nil {
		log.Fatal(error)
	}

	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}

	blockchain := Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}

	fmt.Println("Blockchain created with difficulty: ", blockchain.difficulty)
}

// This function mines a new block into the blockchain
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

// This function evaluate if a blockchain is or continues valid
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

////////////////////////////////////////////////////////////////////////////////
// COMMANDS
////////////////////////////////////////////////////////////////////////////////

// Create a new blockchain command using Urfave CLI
func CreateBlockchainCommand() cli.Command {
	return cli.Command{
		Name:     "create-blockchain",
		Aliases:  []string{"cb"},
		Usage:    "This command creates a new blockchain",
		Category: "blockchain",
		Action:   CreateBlockchain,
	}
}
