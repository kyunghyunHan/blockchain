package main

import blockchain "github.com/kyunghyun/blockchain/block"

func main() {

	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
}
