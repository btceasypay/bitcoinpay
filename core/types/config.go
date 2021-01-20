// Copyright (c) 2020-2021 The bitcoinpay developers

package types

import (
	"math/big"
)

type Config struct {
	Id *big.Int `json:"Id"  required:"true" min:"0"`
}

type configJSON struct {
	Id *UInt256
}
