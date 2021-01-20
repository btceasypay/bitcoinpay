// Copyright (c) 2020-2021 The bitcoinpay developers

package main

import (
	"fmt"
	"github.com/btceasypay/bitcoinpay/core/blockchain"
	"github.com/btceasypay/bitcoinpay/database"
	"github.com/btceasypay/bitcoinpay/log"
	"github.com/btceasypay/bitcoinpay/params"
	"github.com/btceasypay/bitcoinpay/services/mining"
	"path"
)

type SrcNode struct {
	name string
	bc   *blockchain.BlockChain
	db   database.DB
	cfg  *Config
}

func (node *SrcNode) init(cfg *Config) error {
	node.cfg = cfg
	// Load the block database.
	srcDataDir := cfg.SrcDataDir
	if cfg.Last {
		srcDataDir = cfg.DataDir
	}
	db, err := LoadBlockDB(cfg.DbType, srcDataDir, false)
	if err != nil {
		log.Error("load block database", "error", err)
		return err
	}
	defer func() {
		// Ensure the database is sync'd and closed on shutdown.

	}()
	node.db = db
	//

	bc, err := blockchain.New(&blockchain.Config{
		DB:           db,
		ChainParams:  params.ActiveNetParams.Params,
		TimeSource:   blockchain.NewMedianTime(),
		DAGType:      cfg.DAGType,
		BlockVersion: mining.BlockVersion(params.ActiveNetParams.Params.Net),
	})
	if err != nil {
		log.Error(err.Error())
		return err
	}
	node.bc = bc
	node.name = path.Base(srcDataDir)

	log.Info(fmt.Sprintf("Load Src Data:%s", srcDataDir))
	return nil
}

func (node *SrcNode) exit() {
	if node.db != nil {
		log.Info(fmt.Sprintf("Gracefully shutting down the database:%s", node.name))
		node.db.Close()
	}
}

func (node *SrcNode) BlockChain() *blockchain.BlockChain {
	return node.bc
}

func (node *SrcNode) DB() database.DB {
	return node.db
}
