// Copyright (c) 2020-2021 The bitcoinpay developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package bx

import (
	"fmt"
	"github.com/btceasypay/bitcoinpay/crypto/bip32"
	"github.com/btceasypay/bitcoinpay/params"
)

var (
	BitcoinpayMainnetBip32Version = bip32.Bip32Version{PrivKeyVersion: params.MainNetParams.HDPrivateKeyID[:], PubKeyVersion: params.MainNetParams.HDPublicKeyID[:]}
	BitcoinpayTestnetBip32Version = bip32.Bip32Version{PrivKeyVersion: params.TestNetParams.HDPrivateKeyID[:], PubKeyVersion: params.TestNetParams.HDPublicKeyID[:]}
	BitcoinpayPrivnetBip32Version = bip32.Bip32Version{PrivKeyVersion: params.PrivNetParams.HDPrivateKeyID[:], PubKeyVersion: params.PrivNetParams.HDPublicKeyID[:]}
	BitcoinpayMixnetBip32Version  = bip32.Bip32Version{PrivKeyVersion: params.MixNetParam.HDPrivateKeyID[:], PubKeyVersion: params.MixNetParam.HDPublicKeyID[:]}
)

type Bip32VersionFlag struct {
	Version bip32.Bip32Version
	flag    string
}

func (v *Bip32VersionFlag) String() string {
	return v.flag
}

func (v *Bip32VersionFlag) Set(versionFlag string) error {
	var version bip32.Bip32Version
	switch versionFlag {
	case "bip32", "btc":
		version = bip32.DefaultBip32Version
	case "mainnet":
		version = BitcoinpayMainnetBip32Version
	case "testnet":
		version = BitcoinpayTestnetBip32Version
	case "privnet":
		version = BitcoinpayPrivnetBip32Version
	case "mixnet":
		version = BitcoinpayMixnetBip32Version
	default:
		return fmt.Errorf("unknown bip32 version flag %s", versionFlag)
	}
	v.Version = version
	v.flag = versionFlag
	return nil
}

func GetBip32NetworkInfo(rawVersionByte []byte) string {
	if BitcoinpayMainnetBip32Version.IsPrivkeyVersion(rawVersionByte) || BitcoinpayMainnetBip32Version.IsPubkeyVersion(rawVersionByte) {
		return "bx mainet"
	} else if BitcoinpayTestnetBip32Version.IsPrivkeyVersion(rawVersionByte) || BitcoinpayTestnetBip32Version.IsPubkeyVersion(rawVersionByte) {
		return "bx testnet"
	} else if BitcoinpayPrivnetBip32Version.IsPrivkeyVersion(rawVersionByte) || BitcoinpayPrivnetBip32Version.IsPubkeyVersion(rawVersionByte) {
		return "bx privnet"
	} else if BitcoinpayMixnetBip32Version.IsPrivkeyVersion(rawVersionByte) || BitcoinpayMixnetBip32Version.IsPubkeyVersion(rawVersionByte) {
		return "bx mixnet"
	} else if bip32.DefaultBip32Version.IsPrivkeyVersion(rawVersionByte) || bip32.DefaultBip32Version.IsPubkeyVersion(rawVersionByte) {
		return "btc mainnet"
	} else {
		return "unknown"
	}
}
