package main

import (
	"github.com/kyunghyun/blockchain/explorer"
	"github.com/kyunghyun/blockchain/rest"
)

func main() {

	go explorer.Start(3000)
	rest.Start(4000)
}
