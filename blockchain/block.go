package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"Hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"Height"`
}

func createBlock(data, prevHash string, hight int) *Block {
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   hight,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	return &block

}
