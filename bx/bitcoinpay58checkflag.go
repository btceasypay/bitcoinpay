// Copyright (c) 2020-2021 The bitcoinpay developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package bx

import (
	"encoding/hex"
	"github.com/btceasypay/bitcoinpay/params"
)

type BitcoinpayBase58checkVersionFlag struct {
	Ver  []byte
	flag string
}

func (n *BitcoinpayBase58checkVersionFlag) Set(s string) error {
	n.Ver = []byte{}
	switch s {
	case "mainnet":
		n.Ver = append(n.Ver, params.MainNetParams.PubKeyHashAddrID[0:]...)
	case "privnet":
		n.Ver = append(n.Ver, params.PrivNetParams.PubKeyHashAddrID[0:]...)
	case "testnet":
		n.Ver = append(n.Ver, params.TestNetParams.PubKeyHashAddrID[0:]...)
	case "mixnet":
		n.Ver = append(n.Ver, params.MixNetParams.PubKeyHashAddrID[0:]...)
	default:
		v, err := hex.DecodeString(s)
		if err != nil {
			return err
		}
		n.Ver = append(n.Ver, v...)
	}
	n.flag = s
	return nil
}

func (n *BitcoinpayBase58checkVersionFlag) String() string {
	return n.flag
}
