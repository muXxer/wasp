// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package jsonrpc

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/parameters"
	"github.com/iotaledger/wasp/packages/state"
)

type ChainBackend interface {
	EVMSendTransaction(tx *types.Transaction) error
	EVMEstimateGas(callMsg ethereum.CallMsg) (uint64, error)
	ISCCallView(chainState state.State, scName string, funName string, args dict.Dict) (dict.Dict, error)
	ISCLatestState() state.State
	ISCStateByBlockIndex(blockIndex uint32) (state.State, error)
	EVMGasPrice() *big.Int
	BaseToken() *parameters.BaseToken
}
