package main

import (
	"github.com/btceasypay/bitcoinpay/core/blockchain"
	"github.com/btceasypay/bitcoinpay/database"
)

type INode interface {
	BlockChain() *blockchain.BlockChain
	DB() database.DB
}
