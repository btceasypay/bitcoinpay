// Copyright (c) 2020-2021 The bitcoinpay developers
package node

import (
	"github.com/btceasypay/bitcoinpay/config"
	"github.com/btceasypay/bitcoinpay/database"
	"github.com/btceasypay/bitcoinpay/p2p/peerserver"
	"github.com/btceasypay/bitcoinpay/rpc"
)

// BitcoinpayLight implements the bitcoinpay light node service.
type BitcoinpayLight struct {
	// database
	db     database.DB
	config *config.Config
}

func (light *BitcoinpayLight) Start(server *peerserver.PeerServer) error {
	log.Debug("Starting bitcoinpay light node service")
	return nil
}

func (light *BitcoinpayLight) Stop() error {
	log.Debug("Stopping bitcoinpay light node service")
	return nil
}

func (light *BitcoinpayLight) APIs() []rpc.API {
	return []rpc.API{}
}

func newBitcoinpayLight(n *Node) (*BitcoinpayLight, error) {
	light := BitcoinpayLight{
		config: n.Config,
		db:     n.DB,
	}
	return &light, nil
}
