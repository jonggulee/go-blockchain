package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jonggulee/go-blockchain/utils"
)

type Block struct {
	Hash         string `json:"Hash"`
	PrevHash     string `json:"prevHash,omitempty"`
	Height       int    `json:"Height"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestemp    int    `json:"timestemp"`
	Transactions []*Tx  `json:"transactions"`
}

func persistBlock(b *Block) {
	dbStorage.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := dbStorage.FindBlock(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestemp = int(time.Now().Unix())
		hash := utils.Hash(b)
		fmt.Printf("\n\n\nTarget:%s\nHash:%s\nNonce:%d\n\n\n", target, hash, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, hight, diff int) *Block {
	block := &Block{
		Hash:       "",
		PrevHash:   prevHash,
		Height:     hight,
		Difficulty: diff,
		Nonce:      0,
	}
	block.Transactions = Mempool().TxToConfirm()
	block.mine()
	persistBlock(block)
	return block
}
