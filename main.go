package main

import (
	"eng-gabrielscardoso/blockchain-go/core"
	"fmt"
)

func main() {
	blockchain := core.CreateBlockchain(2)

	blockchain.AddBlock("Alice", "Bob", 5)
	blockchain.AddBlock("John", "Bob", 2)

	fmt.Println(blockchain.IsValid())
}
