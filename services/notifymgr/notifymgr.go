package notifymgr

import (
	"github.com/btceasypay/bitcoinpay/core/message"
	"github.com/btceasypay/bitcoinpay/core/types"
	"github.com/btceasypay/bitcoinpay/p2p/peerserver"
	"github.com/btceasypay/bitcoinpay/rpc"
)

// NotifyMgr manage message announce & relay & notification between mempool, websocket, gbt long pull
// and rpc server.
type NotifyMgr struct {
	Server    *peerserver.PeerServer
	RpcServer *rpc.RpcServer
}

// AnnounceNewTransactions generates and relays inventory vectors and notifies
// both websocket and getblocktemplate long poll clients of the passed
// transactions.  This function should be called whenever new transactions
// are added to the mempool.
func (ntmgr *NotifyMgr) AnnounceNewTransactions(newTxs []*types.TxDesc) {
	// Generate and relay inventory vectors for all newly accepted
	// transactions into the memory pool due to the original being
	// accepted.
	for _, tx := range newTxs {
		// Generate the inventory vector and relay it.
		iv := message.NewInvVect(message.InvTypeTx, tx.Tx.Hash())
		// reply to p2p
		ntmgr.RelayInventory(iv, tx)
		// reply to rpc

	}
}

// RelayInventory relays the passed inventory vector to all connected peers
// that are not already known to have it.
func (ntmgr *NotifyMgr) RelayInventory(invVect *message.InvVect, data interface{}) {
	ntmgr.Server.RelayInventory(invVect, data)
}

func (ntmgr *NotifyMgr) BroadcastMessage(msg message.Message) {
	ntmgr.Server.BroadcastMessage(msg)
}
