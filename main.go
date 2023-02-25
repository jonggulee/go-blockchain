package main

import (
	"github.com/jonggu/jakecoin/cli"
	"github.com/jonggu/jakecoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
