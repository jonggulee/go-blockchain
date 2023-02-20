package main

import (
	"github.com/jonggu/jakecoin/blockchain"
	"github.com/jonggu/jakecoin/cli"
	"github.com/jonggu/jakecoin/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()
}
