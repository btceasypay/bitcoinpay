package cryptonight

import (
	"github.com/btceasypay/bitcoinpay/crypto/cryptonight/groestl"
	"github.com/btceasypay/bitcoinpay/crypto/cryptonight/jh"
	"hash"
	"sync"
	"unsafe"

	"github.com/aead/skein"
	"github.com/dchest/blake256"
)

var hashPool = [...]*sync.Pool{
	{New: func() interface{} { return blake256.New() }},
	{New: func() interface{} { return groestl.New256() }},
	{New: func() interface{} { return jh.New256() }},
	{New: func() interface{} { return skein.New256(nil) }},
}

func (cc *cache) finalHash() []byte {
	hp := hashPool[cc.finalState[0]&0x03]
	h := hp.Get().(hash.Hash)
	h.Reset()
	h.Write((*[200]byte)(unsafe.Pointer(&cc.finalState))[:])
	sum := h.Sum(nil)
	hp.Put(h)

	return sum
}
