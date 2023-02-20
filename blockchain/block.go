package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/jonggu/jakecoin/db"
	"github.com/jonggu/jakecoin/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"Hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"Height"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func createBlock(data, prevHash string, hight int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   hight,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return block
}
