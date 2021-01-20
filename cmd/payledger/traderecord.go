// Copyright (c) 2020-2021 The bitcoinpay developers

package main

import "github.com/btceasypay/bitcoinpay/common/hash"

type TradeRecord struct {
	blockHash    *hash.Hash
	blockId      uint
	blockOrder   uint
	blockConfirm uint
	blockStatus  byte
	blockBlue    int // 0:not blue;  1：blue  2：Cannot confirm
	blockHeight  uint
	txHash       *hash.Hash
	txFullHash   *hash.Hash
	txUIndex     int
	txValid      bool
	txIsIn       bool
	amount       uint64
	isCoinbase   bool
}
