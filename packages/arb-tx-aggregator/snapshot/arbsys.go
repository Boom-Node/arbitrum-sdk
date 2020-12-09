/*
 * Copyright 2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package snapshot

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/pkg/errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var (
	txCountABI        abi.Method
	withdrawEthABI    abi.Method
	withdrawERC20ABI  abi.Method
	withdrawERC721ABI abi.Method
	getStorageAtABI   abi.Method

	ethWithdrawal    ethcommon.Hash
	erc20Withdrawal  ethcommon.Hash
	erc721Withdrawal ethcommon.Hash

	arbsysConn *bind.BoundContract
)

func init() {
	arbsys, err := abi.JSON(strings.NewReader(arboscontracts.ArbSysABI))
	if err != nil {
		panic(err)
	}

	txCountABI = arbsys.Methods["getTransactionCount"]
	withdrawEthABI = arbsys.Methods["withdrawEth"]
	withdrawERC20ABI = arbsys.Methods["withdrawERC20"]
	withdrawERC721ABI = arbsys.Methods["withdrawERC721"]
	getStorageAtABI = arbsys.Methods["getStorageAt"]

	ethWithdrawal = arbsys.Events["EthWithdrawal"].ID
	erc20Withdrawal = arbsys.Events["ERC20Withdrawal"].ID
	erc721Withdrawal = arbsys.Events["ERC721Withdrawal"].ID

	arbsysConn = bind.NewBoundContract(arbos.ARB_SYS_ADDRESS, arbsys, nil, nil, nil)
}

func TransactionCountData(address common.Address) []byte {
	return makeFuncData(txCountABI, address)
}

func ParseTransactionCountResult(res *evm.TxResult) (*big.Int, error) {
	vals, err := txCountABI.Outputs.UnpackValues(res.ReturnData)
	if err != nil {
		return nil, err
	}
	val, ok := vals[0].(*big.Int)
	if !ok {
		return nil, errors.New("unexpected tx result")
	}
	return val, nil
}

func WithdrawEthData(address common.Address) []byte {
	return makeFuncData(withdrawEthABI, address)
}

func WithdrawERC20Data(address common.Address, amount *big.Int) []byte {
	return makeFuncData(withdrawERC20ABI, address, amount)
}

func WithdrawERC721Data(address common.Address, id *big.Int) []byte {
	return makeFuncData(withdrawERC721ABI, address, id)
}

func encodeLog(log evm.Log) types.Log {
	return types.Log{
		Address: log.Address.ToEthAddress(),
		Topics:  common.NewEthHashesFromHashes(log.Topics),
		Data:    log.Data,
	}
}

func ParseEthWithdrawalEvent(log evm.Log) (*arboscontracts.ArbSysEthWithdrawal, error) {
	if log.Topics[0].ToEthHash() != ethWithdrawal {
		return nil, errors.New("wrong event type")
	}
	event := new(arboscontracts.ArbSysEthWithdrawal)
	if err := arbsysConn.UnpackLog(event, "EthWithdrawal", encodeLog(log)); err != nil {
		return nil, err
	}
	return event, nil
}

func ParseERC20WithdrawalEvent(log evm.Log) (*arboscontracts.ArbSysERC20Withdrawal, error) {
	if log.Topics[0].ToEthHash() != erc20Withdrawal {
		return nil, errors.New("wrong event type")
	}
	event := new(arboscontracts.ArbSysERC20Withdrawal)
	if err := arbsysConn.UnpackLog(event, "ERC20Withdrawal", encodeLog(log)); err != nil {
		return nil, err
	}
	return event, nil
}

func ParseERC721WithdrawalEvent(log evm.Log) (*arboscontracts.ArbSysERC721Withdrawal, error) {
	if log.Topics[0].ToEthHash() != erc721Withdrawal {
		return nil, errors.New("wrong event type")
	}
	event := new(arboscontracts.ArbSysERC721Withdrawal)
	if err := arbsysConn.UnpackLog(event, "ERC721Withdrawal", encodeLog(log)); err != nil {
		return nil, err
	}
	return event, nil
}

func StorageAtData(address common.Address, index *big.Int) []byte {
	return makeFuncData(getStorageAtABI, address, index)
}

func parseGetStorageAtResult(res *evm.TxResult) (*big.Int, error) {
	vals, err := getStorageAtABI.Outputs.UnpackValues(res.ReturnData)
	if err != nil {
		return nil, err
	}
	val, ok := vals[0].(*big.Int)
	if !ok {
		return nil, errors.New("unexpected tx result")
	}
	return val, nil
}

func makeFuncData(funcABI abi.Method, params ...interface{}) []byte {
	txData, err := funcABI.Inputs.Pack(params...)
	if err != nil {
		panic(err)
	}
	return append(funcABI.ID, txData...)
}
