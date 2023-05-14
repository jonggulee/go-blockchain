package main

import (
	"github.com/jonggu/jakecoin/cli"
	"github.com/jonggu/jakecoin/db"
	"github.com/jonggu/jakecoin/wallet"
)

func main() {
	defer db.Close()
	cli.Start()
	wallet.Wallet()
}
