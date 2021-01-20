package notify

import (
	"github.com/btceasypay/bitcoinpay/core/message"
	"github.com/btceasypay/bitcoinpay/core/types"
)

// Notify interface manage message announce & relay & notification between mempool, websocket, gbt long pull
// and rpc server.
type Notify interface {
	AnnounceNewTransactions(newTxs []*types.TxDesc)
	RelayInventory(invVect *message.InvVect, data interface{})
	BroadcastMessage(msg message.Message)
}
