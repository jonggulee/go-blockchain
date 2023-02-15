package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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

func (b *Block) toBytes() []byte {
	var blockBuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockBuffer)
	err := encoder.Encode(b)
	utils.HandleErr(err)
	return blockBuffer.Bytes()
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
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
	block.persist()
	return &block

}
