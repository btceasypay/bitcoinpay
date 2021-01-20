// Copyright (c) 2020-2021 The bitcoinpay developers
package node

import (
	"github.com/btceasypay/bitcoinpay/core/blockchain"
	"github.com/btceasypay/bitcoinpay/database"
	"github.com/btceasypay/bitcoinpay/engine/txscript"
	"github.com/btceasypay/bitcoinpay/node/notify"
	"github.com/btceasypay/bitcoinpay/p2p/peerserver"
	"github.com/btceasypay/bitcoinpay/params"
	"github.com/btceasypay/bitcoinpay/rpc"
	"github.com/btceasypay/bitcoinpay/services/acct"
	"github.com/btceasypay/bitcoinpay/services/address"
	"github.com/btceasypay/bitcoinpay/services/blkmgr"
	"github.com/btceasypay/bitcoinpay/services/common"
	"github.com/btceasypay/bitcoinpay/services/index"
	"github.com/btceasypay/bitcoinpay/services/mempool"
	"github.com/btceasypay/bitcoinpay/services/miner"
	"github.com/btceasypay/bitcoinpay/services/mining"
	"github.com/btceasypay/bitcoinpay/services/notifymgr"
	"github.com/btceasypay/bitcoinpay/services/tx"
)

// BitcoinpayFull implements the Bitcoinpay full node service.
type BitcoinpayFull struct {
	// under node
	node *Node
	// msg notifier
	nfManager notify.Notify
	// database
	db database.DB
	// account/wallet service
	acctmanager *acct.AccountManager
	// block manager handles all incoming blocks.
	blockManager *blkmgr.BlockManager
	// tx manager
	txManager *tx.TxManager

	// miner service
	cpuMiner *miner.CPUMiner

	// address service
	addressApi *address.AddressApi

	// clock time service
	timeSource blockchain.MedianTimeSource
	// signature cache
	sigCache *txscript.SigCache
}

func (qm *BitcoinpayFull) Start(server *peerserver.PeerServer) error {
	log.Debug("Starting Bitcoinpay full node service")

	// Start the CPU miner if generation is enabled.
	if qm.node.Config.Generate {
		qm.cpuMiner.Start()
	}

	qm.blockManager.Start()
	qm.txManager.Start()
	return nil
}

func (qm *BitcoinpayFull) Stop() error {
	log.Debug("Stopping Bitcoinpay full node service")

	log.Info("try stop bm")

	qm.blockManager.Stop()
	qm.blockManager.WaitForStop()

	qm.txManager.Stop()

	log.Info("try stop cpu miner")
	// Stop the CPU miner if needed.
	if qm.node.Config.Generate && qm.cpuMiner != nil {
		qm.cpuMiner.Stop()
	}

	return nil
}

func (qm *BitcoinpayFull) APIs() []rpc.API {
	apis := qm.acctmanager.APIs()
	apis = append(apis, qm.addressApi.APIs()...)
	apis = append(apis, qm.cpuMiner.APIs()...)
	apis = append(apis, qm.blockManager.API())
	apis = append(apis, qm.txManager.APIs()...)
	apis = append(apis, qm.apis()...)
	return apis
}
func newBitcoinpayFullNode(node *Node) (*BitcoinpayFull, error) {

	// account manager
	acctmgr, err := acct.New()
	if err != nil {
		return nil, err
	}
	qm := BitcoinpayFull{
		node:        node,
		db:          node.DB,
		acctmanager: acctmgr,
		timeSource:  blockchain.NewMedianTime(),
		sigCache:    txscript.NewSigCache(node.Config.SigCacheMaxSize),
	}
	// Create the transaction and address indexes if needed.
	var indexes []index.Indexer
	cfg := node.Config

	var txIndex *index.TxIndex
	var addrIndex *index.AddrIndex
	log.Info("Transaction index is enabled")
	txIndex = index.NewTxIndex(qm.db)
	indexes = append(indexes, txIndex)
	if cfg.AddrIndex {
		log.Info("Address index is enabled")
		addrIndex = index.NewAddrIndex(qm.db, node.Params)
		indexes = append(indexes, addrIndex)
	}
	// index-manager
	var indexManager blockchain.IndexManager
	if len(indexes) > 0 {
		indexManager = index.NewManager(qm.db, indexes, node.Params)
	}

	qm.nfManager = &notifymgr.NotifyMgr{Server: node.peerServer, RpcServer: node.rpcServer}

	// block-manager
	bm, err := blkmgr.NewBlockManager(qm.nfManager, indexManager, node.DB, qm.timeSource, qm.sigCache, node.Config, node.Params,
		mining.BlockVersion(node.Params.Net), node.quit)
	if err != nil {
		return nil, err
	}
	qm.blockManager = bm

	// txmanager
	tm, err := tx.NewTxManager(bm, txIndex, addrIndex, cfg, qm.nfManager, qm.sigCache, node.DB)
	if err != nil {
		return nil, err
	}
	qm.txManager = tm
	bm.SetTxManager(tm)
	// prepare peerServer
	node.peerServer.BlockManager = bm
	node.peerServer.TimeSource = qm.timeSource
	node.peerServer.TxMemPool = qm.txManager.MemPool().(*mempool.TxPool)

	// Cpu Miner
	// Create the mining policy based on the configuration options.
	// NOTE: The CPU miner relies on the mempool, so the mempool has to be
	// created before calling the function to create the CPU miner.
	policy := mining.Policy{
		BlockMinSize:      cfg.BlockMinSize,
		BlockMaxSize:      cfg.BlockMaxSize,
		BlockPrioritySize: cfg.BlockPrioritySize,
		TxMinFreeFee:      cfg.MinTxFee, //TODO, duplicated config item with mem-pool
		StandardVerifyFlags: func() (txscript.ScriptFlags, error) {
			return common.StandardScriptVerifyFlags()
		}, //TODO, duplicated config item with mem-pool
	}
	// defaultNumWorkers is the default number of workers to use for mining
	// and is based on the number of processor cores.  This helps ensure the
	// system stays reasonably responsive under heavy load.
	defaultNumWorkers := uint32(params.CPUMinerThreads) //TODO, move to config

	qm.cpuMiner = miner.NewCPUMiner(cfg, node.Params, &policy, qm.sigCache,
		qm.txManager.MemPool().(*mempool.TxPool), qm.timeSource, qm.blockManager, defaultNumWorkers)
	// init address api
	qm.addressApi = address.NewAddressApi(cfg, node.Params)
	return &qm, nil
}

// return block manager
func (qm *BitcoinpayFull) GetBlockManager() *blkmgr.BlockManager {
	return qm.blockManager
}

// return cpu miner
func (qm *BitcoinpayFull) GetCpuMiner() *miner.CPUMiner {
	return qm.cpuMiner
}

// return address api
func (qm *BitcoinpayFull) GetAddressApi() *address.AddressApi {
	return qm.addressApi
}

// return peer server
func (qm *BitcoinpayFull) GetPeerServer() *peerserver.PeerServer {
	return qm.node.peerServer
}
