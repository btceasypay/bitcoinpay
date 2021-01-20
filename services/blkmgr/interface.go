// Copyright (c) 2020-2021 The bitcoinpay developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package blkmgr

import (
	"github.com/btceasypay/bitcoinpay/common/hash"
	"github.com/btceasypay/bitcoinpay/core/blockchain"
	"github.com/btceasypay/bitcoinpay/core/types"
)

type TxManager interface {
	MemPool() TxPool
}

type TxPool interface {
	AddTransaction(utxoView *blockchain.UtxoViewpoint,
		tx *types.Tx, height uint64, fee int64)

	RemoveTransaction(tx *types.Tx, removeRedeemers bool)

	RemoveDoubleSpends(tx *types.Tx)

	RemoveOrphan(txHash *hash.Hash)

	ProcessOrphans(hash *hash.Hash) []*types.TxDesc

	MaybeAcceptTransaction(tx *types.Tx, isNew, rateLimit bool) ([]*hash.Hash, error)

	HaveTransaction(hash *hash.Hash) bool

	PruneExpiredTx()

	ProcessTransaction(tx *types.Tx, allowOrphan, rateLimit, allowHighFees bool) ([]*types.TxDesc, error)
}
