package main

import (
	"github.com/jonggulee/go-blockchain/cli"
	"github.com/jonggulee/go-blockchain/db"
)

func main() {
	defer db.Close()
	db.InitDB()
	cli.Start()
}
