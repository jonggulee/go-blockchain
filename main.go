package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	// defer db.Close()
	// blockchain.Blockchain()
	// cli.Start()

	difficutly := 5
	target := strings.Repeat("0", difficutly)
	nonce := 1
	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("Hello"+fmt.Sprint(nonce))))
		fmt.Printf("Hash: %s\nTarget: %s\nNonce: %d\n\n", hash, target, nonce)
		if strings.HasPrefix(hash, target) {
			fmt.Println(nonce)
			return
		} else {
			nonce++
		}
	}
}
