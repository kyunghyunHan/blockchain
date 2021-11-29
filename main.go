package main

import (
	blockchain "github.com/kyunghyun/blockchain/block"
	"github.com/kyunghyun/blockchain/cli"
)

func main() {

	blockchain.Blockchain()
	cli.Start()
}
