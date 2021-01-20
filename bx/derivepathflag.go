// Copyright (c) 2020-2021 The bitcoinpay developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package bx

import "github.com/btceasypay/bitcoinpay/wallet"

type DerivePathFlag struct {
	Path wallet.DerivationPath
}

func (d *DerivePathFlag) Set(s string) error {
	path, err := wallet.ParseDerivationPath(s)
	if err != nil {
		return err
	}
	d.Path = path
	return nil
}

func (d *DerivePathFlag) String() string {
	return d.Path.String()
}
