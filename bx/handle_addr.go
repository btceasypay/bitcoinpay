package bx

import (
	"encoding/hex"
	"fmt"
	"github.com/btceasypay/bitcoinpay/common/encode/base58"
	"github.com/btceasypay/bitcoinpay/common/hash"
	"github.com/btceasypay/bitcoinpay/params"
)

func EcPubKeyToAddress(version string, pubkey string) (string, error) {
	ver := []byte{}

	switch version {
	case "mainnet":
		ver = append(ver, params.MainNetParams.PubKeyHashAddrID[0:]...)
	case "privnet":
		ver = append(ver, params.PrivNetParams.PubKeyHashAddrID[0:]...)
	case "testnet":
		ver = append(ver, params.TestNetParams.PubKeyHashAddrID[0:]...)
	case "mixnet":
		ver = append(ver, params.MixNetParam.PubKeyHashAddrID[0:]...)
	default:
		v, err := hex.DecodeString(version)
		if err != nil {
			return "", err
		}
		ver = append(ver, v...)
	}

	data, err := hex.DecodeString(pubkey)
	if err != nil {
		return "", err
	}
	h := hash.Hash160(data)

	address := base58.BitcoinpayCheckEncode(h, ver[:])
	return address, nil
}

func EcScriptKeyToAddress(version string, pubkey string) (string, error) {
	ver := []byte{}

	switch version {
	case "mainnet":
		ver = append(ver, params.MainNetParams.ScriptHashAddrID[0:]...)
	case "privnet":
		ver = append(ver, params.PrivNetParams.ScriptHashAddrID[0:]...)
	case "testnet":
		ver = append(ver, params.TestNetParams.ScriptHashAddrID[0:]...)
	case "mixnet":
		ver = append(ver, params.MixNetParam.ScriptHashAddrID[0:]...)
	default:
		v, err := hex.DecodeString(version)
		if err != nil {
			return "", err
		}
		ver = append(ver, v...)
	}

	data, err := hex.DecodeString(pubkey)
	if err != nil {
		return "", err
	}
	h := hash.Hash160(data)

	address := base58.BitcoinpayCheckEncode(h, ver[:])
	return address, nil
}

func EcPubKeyToAddressSTDO(version []byte, pubkey string) {
	data, err := hex.DecodeString(pubkey)
	if err != nil {
		ErrExit(err)
	}
	h := hash.Hash160(data)

	address := base58.BitcoinpayCheckEncode(h, version[:])
	fmt.Printf("%s\n", address)
}
