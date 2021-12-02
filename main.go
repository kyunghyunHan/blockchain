package main

import (
	"github.com/kyunghyun/blockchain/cli"
	"github.com/kyunghyun/blockchain/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
