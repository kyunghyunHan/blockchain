package main

import (
	"fmt"

	blockchain "github.com/kyunghyun/blockchain/block"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third  Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Hash:%s\n", block.Hash)
		fmt.Printf("Prev Hash:%s\n", block.PrevHash)
	}
}
