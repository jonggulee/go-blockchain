package main

import (
	"github.com/jonggu/jakecoin/explorer"
	"github.com/jonggu/jakecoin/rest"
)

func main() {
	go rest.Start(4000)
	explorer.Start(3000)
}
