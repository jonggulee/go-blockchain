package blockchain

import (
	"errors"
	"time"

	"github.com/jonggu/jakecoin/utils"
)

const (
	minerReward int = 50
)

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

type Tx struct {
	Id        string   `json:"id"`
	Timestemp int      `json:"timestemp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

func (t *Tx) getID() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	Owner  string
	Amount int
}

type TxOut struct {
	Owner  string
	Amount int
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		Id:        "",
		Timestemp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getID()
	return &tx
}

func makeTx(from, to string, amout int) (*Tx, error) {
	if Blockchain().BalanceByAddress(from) < amout {
		return nil, errors.New("not enough money")
	}
}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("jake", to, amount)
	// utils.HandleErr(err)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}
