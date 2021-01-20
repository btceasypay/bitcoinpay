// Copyright (c) 2020-2021 The bitcoinpay developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package bx

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Base64Encode(input string) {
	data, err := hex.DecodeString(input)
	if err != nil {
		ErrExit(err)
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("%s\n", encoded)
}

func Base64Decode(input string) {
	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		ErrExit(err)
	}
	fmt.Printf("%x\n", data)
}
