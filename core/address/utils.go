// Copyright (c) 2020-2021 The bitcoinpay developers

package address

import (
	"github.com/btceasypay/bitcoinpay/core/types"
	"github.com/btceasypay/bitcoinpay/params"
)

// IsForNetwork returns whether or not the address is associated with the
// passed network.
//TODO, other addr type and ec type check
func IsForNetwork(addr types.Address, p *params.Params) bool {
	switch addr := addr.(type) {
	case *PubKeyHashAddress:
		return addr.netID == p.PubKeyHashAddrID
	}
	return false
}
