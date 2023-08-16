package main

import (
	"github.com/jonggulee/go-coin/cli"
	"github.com/jonggulee/go-coin/db"
	"github.com/jonggulee/go-coin/wallet"
)

func main() {
	defer db.Close()
	cli.Start()
	wallet.Wallet()
}
