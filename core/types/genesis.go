// Copyright (c) 2020-2021 The bitcoinpay developers

package types

type Genesis struct {
	Config *Config `json:"config" required:"true"`
	Nonce  uint64  `json:"nonce"  required:"true" min:"1"`
}

type genesisJSON struct {
	Config *Config
	Nonce  UInt64
}
