package main

import (
	"github.com/jonggulee/go-blockchain/cli"
	"github.com/jonggulee/go-blockchain/db"
	"github.com/jonggulee/go-blockchain/wallet"
)

func main() {
	defer db.Close()
	cli.Start()
	wallet.Wallet()
}
