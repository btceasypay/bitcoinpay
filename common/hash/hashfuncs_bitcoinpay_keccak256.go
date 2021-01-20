// Copyright (c) 2020-2021 The bitcoinpay developers
package hash

import (
	bitcoinpaySha3 "github.com/btceasypay/bitcoinpay/crypto/sha3"
)

// Bitcoinpay Keccak256 calculates hash(b) and returns the resulting bytes as a Hash.
func HashBitcoinpayKeccak256(b []byte) Hash {
	h := bitcoinpaySha3.NewBitcoinpayKeccak256()
	h.Write(b)
	r := h.Sum(nil)
	hashR := [32]byte{}
	copy(hashR[:32], r[:32])
	return Hash(hashR)
}
