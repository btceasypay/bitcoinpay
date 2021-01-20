// Copyright (c) 2020-2021 The bitcoinpay developers

package blockchain

import (
	"math/rand"
	"sort"
	"testing"
)

func Test_SortOrphanBlockSlice(t *testing.T) {
	obs := orphanBlockSlice{}

	for i := uint(0); i < 5; i++ {
		obs = append(obs, &orphanBlock{height: uint64(rand.Intn(100))})
	}
	if len(obs) >= 2 {
		sort.Sort(obs)
	}

}
