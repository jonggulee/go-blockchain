package main

import (
	"github.com/jonggu/jakecoin/blockchain"
	"github.com/jonggu/jakecoin/cli"
)

func main() {
	blockchain.Blockchain()
	// blockchain.Blockchain().AddBlock("First")
	// blockchain.Blockchain().AddBlock("Second")
	// blockchain.Blockchain().AddBlock("Third")
	cli.Start()
}
