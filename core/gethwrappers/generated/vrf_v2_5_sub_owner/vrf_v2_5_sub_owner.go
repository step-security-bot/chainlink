// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_v2_5_sub_owner

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

var VRFV25ExternalSubOwnerExampleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"name\":\"OnlyOwnerOrCoordinator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"nativePayment\",\"type\":\"bool\"}],\"name\":\"requestRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_randomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_requestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"}],\"name\":\"setCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610d25380380610d2583398101604081905261002f916101d0565b8133806000816100865760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100b6576100b68161010a565b5050600280546001600160a01b039384166001600160a01b03199182161790915560038054958416958216959095179094555060048054929091169183169190911790556007805490911633179055610203565b6001600160a01b0381163314156101635760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007d565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b03811681146101cb57600080fd5b919050565b600080604083850312156101e357600080fd5b6101ec836101b4565b91506101fa602084016101b4565b90509250929050565b610b13806102126000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638ea981171161005b5780638ea98117146100ea578063e89e106a146100fd578063f2fde38b14610114578063f6eaffc81461012757600080fd5b80631fe543e31461008d5780635b6c5de8146100a257806379ba5097146100b55780638da5cb5b146100bd575b600080fd5b6100a061009b3660046108b0565b61013a565b005b6100a06100b036600461099f565b6101c0565b6100a06102d3565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100a06100f8366004610841565b6103d0565b61010660065481565b6040519081526020016100e1565b6100a0610122366004610841565b6104db565b61010661013536600461087e565b6104ef565b60025473ffffffffffffffffffffffffffffffffffffffff1633146101b2576002546040517f1cf993f400000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911660248201526044015b60405180910390fd5b6101bc8282610510565b5050565b6101c8610593565b60006040518060c001604052808481526020018881526020018661ffff1681526020018763ffffffff1681526020018563ffffffff16815260200161021c6040518060200160405280861515815250610616565b90526003546040517f9b1c385e00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff1690639b1c385e90610275908490600401610a17565b602060405180830381600087803b15801561028f57600080fd5b505af11580156102a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c79190610897565b60065550505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610354576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016101a9565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590610410575060025473ffffffffffffffffffffffffffffffffffffffff163314155b15610494573361043560005473ffffffffffffffffffffffffffffffffffffffff1690565b6002546040517f061db9c100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff938416600482015291831660248301529190911660448201526064016101a9565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6104e3610593565b6104ec816106d2565b50565b600581815481106104ff57600080fd5b600091825260209091200154905081565b600654821461057b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265717565737420494420697320696e636f727265637400000000000000000060448201526064016101a9565b805161058e9060059060208401906107c8565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610614576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016101a9565b565b60607f92fd13387c7fe7befbc38d303d6468778fb9731bc4583f17d92989c6fcfdeaaa8260405160240161064f91511515815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915292915050565b73ffffffffffffffffffffffffffffffffffffffff8116331415610752576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016101a9565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215610803579160200282015b828111156108035782518255916020019190600101906107e8565b5061080f929150610813565b5090565b5b8082111561080f5760008155600101610814565b803563ffffffff8116811461083c57600080fd5b919050565b60006020828403121561085357600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461087757600080fd5b9392505050565b60006020828403121561089057600080fd5b5035919050565b6000602082840312156108a957600080fd5b5051919050565b600080604083850312156108c357600080fd5b8235915060208084013567ffffffffffffffff808211156108e357600080fd5b818601915086601f8301126108f757600080fd5b81358181111561090957610909610ad7565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f8301168101818110858211171561094c5761094c610ad7565b604052828152858101935084860182860187018b101561096b57600080fd5b600095505b8386101561098e578035855260019590950194938601938601610970565b508096505050505050509250929050565b60008060008060008060c087890312156109b857600080fd5b863595506109c860208801610828565b9450604087013561ffff811681146109df57600080fd5b93506109ed60608801610828565b92506080870135915060a08701358015158114610a0957600080fd5b809150509295509295509295565b6000602080835283518184015280840151604084015261ffff6040850151166060840152606084015163ffffffff80821660808601528060808701511660a0860152505060a084015160c08085015280518060e086015260005b81811015610a8e5782810184015186820161010001528301610a71565b81811115610aa157600061010083880101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169390930161010001949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var VRFV25ExternalSubOwnerExampleABI = VRFV25ExternalSubOwnerExampleMetaData.ABI

var VRFV25ExternalSubOwnerExampleBin = VRFV25ExternalSubOwnerExampleMetaData.Bin

func DeployVRFV25ExternalSubOwnerExample(auth *bind.TransactOpts, backend bind.ContractBackend, vrfCoordinator common.Address, link common.Address) (common.Address, *types.Transaction, *VRFV25ExternalSubOwnerExample, error) {
	parsed, err := VRFV25ExternalSubOwnerExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFV25ExternalSubOwnerExampleBin), backend, vrfCoordinator, link)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFV25ExternalSubOwnerExample{VRFV25ExternalSubOwnerExampleCaller: VRFV25ExternalSubOwnerExampleCaller{contract: contract}, VRFV25ExternalSubOwnerExampleTransactor: VRFV25ExternalSubOwnerExampleTransactor{contract: contract}, VRFV25ExternalSubOwnerExampleFilterer: VRFV25ExternalSubOwnerExampleFilterer{contract: contract}}, nil
}

type VRFV25ExternalSubOwnerExample struct {
	address common.Address
	abi     abi.ABI
	VRFV25ExternalSubOwnerExampleCaller
	VRFV25ExternalSubOwnerExampleTransactor
	VRFV25ExternalSubOwnerExampleFilterer
}

type VRFV25ExternalSubOwnerExampleCaller struct {
	contract *bind.BoundContract
}

type VRFV25ExternalSubOwnerExampleTransactor struct {
	contract *bind.BoundContract
}

type VRFV25ExternalSubOwnerExampleFilterer struct {
	contract *bind.BoundContract
}

type VRFV25ExternalSubOwnerExampleSession struct {
	Contract     *VRFV25ExternalSubOwnerExample
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFV25ExternalSubOwnerExampleCallerSession struct {
	Contract *VRFV25ExternalSubOwnerExampleCaller
	CallOpts bind.CallOpts
}

type VRFV25ExternalSubOwnerExampleTransactorSession struct {
	Contract     *VRFV25ExternalSubOwnerExampleTransactor
	TransactOpts bind.TransactOpts
}

type VRFV25ExternalSubOwnerExampleRaw struct {
	Contract *VRFV25ExternalSubOwnerExample
}

type VRFV25ExternalSubOwnerExampleCallerRaw struct {
	Contract *VRFV25ExternalSubOwnerExampleCaller
}

type VRFV25ExternalSubOwnerExampleTransactorRaw struct {
	Contract *VRFV25ExternalSubOwnerExampleTransactor
}

func NewVRFV25ExternalSubOwnerExample(address common.Address, backend bind.ContractBackend) (*VRFV25ExternalSubOwnerExample, error) {
	abi, err := abi.JSON(strings.NewReader(VRFV25ExternalSubOwnerExampleABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFV25ExternalSubOwnerExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFV25ExternalSubOwnerExample{address: address, abi: abi, VRFV25ExternalSubOwnerExampleCaller: VRFV25ExternalSubOwnerExampleCaller{contract: contract}, VRFV25ExternalSubOwnerExampleTransactor: VRFV25ExternalSubOwnerExampleTransactor{contract: contract}, VRFV25ExternalSubOwnerExampleFilterer: VRFV25ExternalSubOwnerExampleFilterer{contract: contract}}, nil
}

func NewVRFV25ExternalSubOwnerExampleCaller(address common.Address, caller bind.ContractCaller) (*VRFV25ExternalSubOwnerExampleCaller, error) {
	contract, err := bindVRFV25ExternalSubOwnerExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV25ExternalSubOwnerExampleCaller{contract: contract}, nil
}

func NewVRFV25ExternalSubOwnerExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFV25ExternalSubOwnerExampleTransactor, error) {
	contract, err := bindVRFV25ExternalSubOwnerExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV25ExternalSubOwnerExampleTransactor{contract: contract}, nil
}

func NewVRFV25ExternalSubOwnerExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFV25ExternalSubOwnerExampleFilterer, error) {
	contract, err := bindVRFV25ExternalSubOwnerExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFV25ExternalSubOwnerExampleFilterer{contract: contract}, nil
}

func bindVRFV25ExternalSubOwnerExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFV25ExternalSubOwnerExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV25ExternalSubOwnerExample.Contract.VRFV25ExternalSubOwnerExampleCaller.contract.Call(opts, result, method, params...)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.VRFV25ExternalSubOwnerExampleTransactor.contract.Transfer(opts)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.VRFV25ExternalSubOwnerExampleTransactor.contract.Transact(opts, method, params...)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV25ExternalSubOwnerExample.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.contract.Transfer(opts)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.contract.Transact(opts, method, params...)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV25ExternalSubOwnerExample.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleSession) Owner() (common.Address, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.Owner(&_VRFV25ExternalSubOwnerExample.CallOpts)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleCallerSession) Owner() (common.Address, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.Owner(&_VRFV25ExternalSubOwnerExample.CallOpts)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleCaller) SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VRFV25ExternalSubOwnerExample.contract.Call(opts, &out, "s_randomWords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.SRandomWords(&_VRFV25ExternalSubOwnerExample.CallOpts, arg0)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleCallerSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.SRandomWords(&_VRFV25ExternalSubOwnerExample.CallOpts, arg0)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleCaller) SRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV25ExternalSubOwnerExample.contract.Call(opts, &out, "s_requestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleSession) SRequestId() (*big.Int, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.SRequestId(&_VRFV25ExternalSubOwnerExample.CallOpts)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleCallerSession) SRequestId() (*big.Int, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.SRequestId(&_VRFV25ExternalSubOwnerExample.CallOpts)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.contract.Transact(opts, "acceptOwnership")
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.AcceptOwnership(&_VRFV25ExternalSubOwnerExample.TransactOpts)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.AcceptOwnership(&_VRFV25ExternalSubOwnerExample.TransactOpts)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.RawFulfillRandomWords(&_VRFV25ExternalSubOwnerExample.TransactOpts, requestId, randomWords)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.RawFulfillRandomWords(&_VRFV25ExternalSubOwnerExample.TransactOpts, requestId, randomWords)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactor) RequestRandomWords(opts *bind.TransactOpts, subId *big.Int, callbackGasLimit uint32, requestConfirmations uint16, numWords uint32, keyHash [32]byte, nativePayment bool) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.contract.Transact(opts, "requestRandomWords", subId, callbackGasLimit, requestConfirmations, numWords, keyHash, nativePayment)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleSession) RequestRandomWords(subId *big.Int, callbackGasLimit uint32, requestConfirmations uint16, numWords uint32, keyHash [32]byte, nativePayment bool) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.RequestRandomWords(&_VRFV25ExternalSubOwnerExample.TransactOpts, subId, callbackGasLimit, requestConfirmations, numWords, keyHash, nativePayment)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactorSession) RequestRandomWords(subId *big.Int, callbackGasLimit uint32, requestConfirmations uint16, numWords uint32, keyHash [32]byte, nativePayment bool) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.RequestRandomWords(&_VRFV25ExternalSubOwnerExample.TransactOpts, subId, callbackGasLimit, requestConfirmations, numWords, keyHash, nativePayment)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactor) SetCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.contract.Transact(opts, "setCoordinator", _vrfCoordinator)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleSession) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.SetCoordinator(&_VRFV25ExternalSubOwnerExample.TransactOpts, _vrfCoordinator)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactorSession) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.SetCoordinator(&_VRFV25ExternalSubOwnerExample.TransactOpts, _vrfCoordinator)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.TransferOwnership(&_VRFV25ExternalSubOwnerExample.TransactOpts, to)
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV25ExternalSubOwnerExample.Contract.TransferOwnership(&_VRFV25ExternalSubOwnerExample.TransactOpts, to)
}

type VRFV25ExternalSubOwnerExampleOwnershipTransferRequestedIterator struct {
	Event *VRFV25ExternalSubOwnerExampleOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV25ExternalSubOwnerExampleOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV25ExternalSubOwnerExampleOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFV25ExternalSubOwnerExampleOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFV25ExternalSubOwnerExampleOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFV25ExternalSubOwnerExampleOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV25ExternalSubOwnerExampleOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25ExternalSubOwnerExampleOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25ExternalSubOwnerExample.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV25ExternalSubOwnerExampleOwnershipTransferRequestedIterator{contract: _VRFV25ExternalSubOwnerExample.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV25ExternalSubOwnerExampleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25ExternalSubOwnerExample.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV25ExternalSubOwnerExampleOwnershipTransferRequested)
				if err := _VRFV25ExternalSubOwnerExample.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFV25ExternalSubOwnerExampleOwnershipTransferRequested, error) {
	event := new(VRFV25ExternalSubOwnerExampleOwnershipTransferRequested)
	if err := _VRFV25ExternalSubOwnerExample.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV25ExternalSubOwnerExampleOwnershipTransferredIterator struct {
	Event *VRFV25ExternalSubOwnerExampleOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV25ExternalSubOwnerExampleOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV25ExternalSubOwnerExampleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFV25ExternalSubOwnerExampleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFV25ExternalSubOwnerExampleOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFV25ExternalSubOwnerExampleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV25ExternalSubOwnerExampleOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25ExternalSubOwnerExampleOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25ExternalSubOwnerExample.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV25ExternalSubOwnerExampleOwnershipTransferredIterator{contract: _VRFV25ExternalSubOwnerExample.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV25ExternalSubOwnerExampleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25ExternalSubOwnerExample.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV25ExternalSubOwnerExampleOwnershipTransferred)
				if err := _VRFV25ExternalSubOwnerExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExampleFilterer) ParseOwnershipTransferred(log types.Log) (*VRFV25ExternalSubOwnerExampleOwnershipTransferred, error) {
	event := new(VRFV25ExternalSubOwnerExampleOwnershipTransferred)
	if err := _VRFV25ExternalSubOwnerExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExample) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFV25ExternalSubOwnerExample.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFV25ExternalSubOwnerExample.ParseOwnershipTransferRequested(log)
	case _VRFV25ExternalSubOwnerExample.abi.Events["OwnershipTransferred"].ID:
		return _VRFV25ExternalSubOwnerExample.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFV25ExternalSubOwnerExampleOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFV25ExternalSubOwnerExampleOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_VRFV25ExternalSubOwnerExample *VRFV25ExternalSubOwnerExample) Address() common.Address {
	return _VRFV25ExternalSubOwnerExample.address
}

type VRFV25ExternalSubOwnerExampleInterface interface {
	Owner(opts *bind.CallOpts) (common.Address, error)

	SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)

	SRequestId(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error)

	RequestRandomWords(opts *bind.TransactOpts, subId *big.Int, callbackGasLimit uint32, requestConfirmations uint16, numWords uint32, keyHash [32]byte, nativePayment bool) (*types.Transaction, error)

	SetCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25ExternalSubOwnerExampleOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV25ExternalSubOwnerExampleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFV25ExternalSubOwnerExampleOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25ExternalSubOwnerExampleOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV25ExternalSubOwnerExampleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFV25ExternalSubOwnerExampleOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
