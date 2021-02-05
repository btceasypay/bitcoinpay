// Copyright (c) 2020-2021 The bitcoinpay developers
// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package params

import (
	"github.com/btceasypay/bitcoinpay/common"
	"github.com/btceasypay/bitcoinpay/core/protocol"
	"github.com/btceasypay/bitcoinpay/core/types/pow"
	"math/big"
	"time"
)

// testNetPowLimit is the highest proof of work value a block can
// have for the test network. It is the value 2^208 - 1.
var testNetPowLimit = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 255), common.Big1)

// target time per block unit second(s)
const testTargetTimePerBlock = 60

// Difficulty check interval is about 60*60 = 60 mins
const testWorkDiffWindowSize = 60

// TestNetParams defines the network parameters for the test network.
var TestNetParams = Params{
	Name:        "testnet",
	Net:         protocol.TestNet,
	DefaultPort: "19130",
	DNSSeeds: []DNSSeed{
		{"seed.biteasypay.xyz", true},
	},

	// Chain parameters
	GenesisBlock: &testNetGenesisBlock,
	GenesisHash:  &testNetGenesisHash,
	PowConfig: &pow.PowConfig{
		Blake2bdPowLimit:                testNetPowLimit,
		Blake2bdPowLimitBits:            0x1b7fffff, // compact from of testNetPowLimit (2^215-1)
		X16rv3PowLimit:                  testNetPowLimit,
		X16rv3PowLimitBits:              0x1b7fffff, // compact from of testNetPowLimit (2^215-1)
		X8r16PowLimit:                   testNetPowLimit,
		X8r16PowLimitBits:               0x1b7fffff, // compact from of testNetPowLimit (2^215-1)
		BitcoinpayKeccak256PowLimit:     testNetPowLimit,
		BitcoinpayKeccak256PowLimitBits: 0x207fffff, // compact from of testNetPowLimit (2^208-1) 453050367
		//hash ffffffffffffffff000000000000000000000000000000000000000000000000 corresponding difficulty is 48 for edge bits 24
		// Uniform field type uint64 value is 48 . bigToCompact the uint32 value
		// 24 edge_bits only need hash 1*4 times use for privnet if GPS is 2. need 50 /2 * 4 = 1min find once
		CuckarooMinDifficulty:  0x2018000, // 96 * 4 = 384
		CuckatooMinDifficulty:  0x2074000, // 1856
		CuckaroomMinDifficulty: 0x1300000, // 48

		Percent: []pow.Percent{
			{
				Blake2bDPercent:            0,
				X16rv3Percent:              0,
				BitcoinpayKeccak256Percent: 100,
				CuckaroomPercent:           0,
				CuckatooPercent:            0,
				MainHeight:                 0,
			},
			{
				Blake2bDPercent:            0,
				X16rv3Percent:              0,
				BitcoinpayKeccak256Percent: 0,
				CuckaroomPercent:           100,
				CuckatooPercent:            0,
				// | time	| timestamp	| mainHeight |
				// | ---| --- | --- |
				// | 2020-08-30 10:31:46 | 1598754706 | 192266
				// | 2020-09-15 12:00 | 1600142400 | 238522
				// The soft forking mainHeight was calculated according to the average time of 30s
				// In other words, BTP will be produced by the pow of BitcoinpayKeccak256 only after mainHeight arrived 238522
				MainHeight: 100000,
			},
		},
		// after this height the big graph will be the main pow graph
		AdjustmentStartMainHeight: 365 * 1440 * 60 / testTargetTimePerBlock,
	},
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0, // Does not apply since ReduceMinDifficulty false
	GenerateSupported:        true,
	WorkDiffAlpha:            1,
	WorkDiffWindowSize:       testWorkDiffWindowSize,
	WorkDiffWindows:          20,
	MaximumBlockSizes:        []int{1310720},
	MaxTxSize:                1000000,
	TargetTimePerBlock:       time.Second * testTargetTimePerBlock,
	TargetTimespan:           time.Second * testTargetTimePerBlock * testWorkDiffWindowSize, // TimePerBlock * WindowSize
	RetargetAdjustmentFactor: 2,

	// Subsidy parameters.
	BaseSubsidy:              5000000000, // 50 Coin , daily supply is 50*60*24 = 72000 ~ 72000 * 2 (DAG factor)
	MulSubsidy:               100,
	DivSubsidy:               200,
	SubsidyReductionInterval: 2100000, // >=210 million ->  >=130 year
	WorkRewardProportion:  10,
	StakeRewardProportion: 0,
	BlockTaxProportion:    0,

	// Maturity
	CoinbaseMaturity: 720, // coinbase required 720 * 60 = 12 hours before repent

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{},

	// Consensus rule change deployments.
	//
	Deployments: map[uint32][]ConsensusDeployment{},

	// Address encoding magics
	NetworkAddressPrefix: "T",
	PubKeyAddrID:         [2]byte{0x0f, 0x0f}, // starts with Tk
	PubKeyHashAddrID:     [2]byte{0x0f, 0x12}, // starts with Tm
	PKHEdwardsAddrID:     [2]byte{0x0f, 0x01}, // starts with Te
	PKHSchnorrAddrID:     [2]byte{0x0f, 0x1e}, // starts with Tr
	ScriptHashAddrID:     [2]byte{0x0e, 0xe2}, // starts with TS
	PrivateKeyID:         [2]byte{0x0c, 0xe2}, // starts with Pt

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x97}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xd1}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 223,

	//OrganizationPkScript:  hexMustDecode("76a914868b9b6bc7e4a9c804ad3d3d7a2a6be27476941e88ac"),
}
