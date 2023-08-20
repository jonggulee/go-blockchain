package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/jonggulee/go-blockchain/blockchain"
	"github.com/jonggulee/go-blockchain/utils"
)

type MessageKind int

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m)
}

func sendNewestBlock(p *peer) {
	fmt.Printf("Sending newest block to %s\n", p.key)
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func requstAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksRequest, nil)
	p.inbox <- m
}

func sendAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksResponse, blockchain.Blocks(blockchain.Blockchain()))
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	// fmt.Printf("Peer: %s, Sent a message with kind of %d\n", p.key, m.Kind)
	switch m.Kind {
	case MessageNewestBlock:
		fmt.Printf("Recieved the newest block from %s\n", p.key)
		var payload blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)
		if payload.Height >= b.Height {
			fmt.Printf("Requesting all blocks from %s\n", p.key)
			requstAllBlocks(p)
		} else {
			fmt.Printf("Sending newest block to %s\n", p.key)
			sendNewestBlock(p)
		}
	case MessageAllBlocksRequest:
		fmt.Printf("%s wants all the blocks\n", p.key)
		sendAllBlocks(p)
	case MessageAllBlocksResponse:
		fmt.Printf("Recieved all the blocks from %s\n", p.key)
		var payload []*blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
	}
}
