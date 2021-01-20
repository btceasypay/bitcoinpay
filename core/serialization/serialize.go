// Copyright (c) 2020-2021 The bitcoinpay developers

package serialization

import "io"

// TODO, redefine the protocol version and storage

type Serializable interface {
	Serialize(w io.Writer) error

	Deserialize(r io.Reader) error
}
