// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrfv2_5_client

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

var VRFV25ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EXTRA_ARGS_V1_TAG\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f7514ab4146038575b600080fd5b605e7f92fd13387c7fe7befbc38d303d6468778fb9731bc4583f17d92989c6fcfdeaaa81565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200160405180910390f3fea164736f6c6343000806000a",
}

var VRFV25ClientABI = VRFV25ClientMetaData.ABI

var VRFV25ClientBin = VRFV25ClientMetaData.Bin

func DeployVRFV25Client(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFV25Client, error) {
	parsed, err := VRFV25ClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFV25ClientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFV25Client{VRFV25ClientCaller: VRFV25ClientCaller{contract: contract}, VRFV25ClientTransactor: VRFV25ClientTransactor{contract: contract}, VRFV25ClientFilterer: VRFV25ClientFilterer{contract: contract}}, nil
}

type VRFV25Client struct {
	address common.Address
	abi     abi.ABI
	VRFV25ClientCaller
	VRFV25ClientTransactor
	VRFV25ClientFilterer
}

type VRFV25ClientCaller struct {
	contract *bind.BoundContract
}

type VRFV25ClientTransactor struct {
	contract *bind.BoundContract
}

type VRFV25ClientFilterer struct {
	contract *bind.BoundContract
}

type VRFV25ClientSession struct {
	Contract     *VRFV25Client
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFV25ClientCallerSession struct {
	Contract *VRFV25ClientCaller
	CallOpts bind.CallOpts
}

type VRFV25ClientTransactorSession struct {
	Contract     *VRFV25ClientTransactor
	TransactOpts bind.TransactOpts
}

type VRFV25ClientRaw struct {
	Contract *VRFV25Client
}

type VRFV25ClientCallerRaw struct {
	Contract *VRFV25ClientCaller
}

type VRFV25ClientTransactorRaw struct {
	Contract *VRFV25ClientTransactor
}

func NewVRFV25Client(address common.Address, backend bind.ContractBackend) (*VRFV25Client, error) {
	abi, err := abi.JSON(strings.NewReader(VRFV25ClientABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFV25Client(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFV25Client{address: address, abi: abi, VRFV25ClientCaller: VRFV25ClientCaller{contract: contract}, VRFV25ClientTransactor: VRFV25ClientTransactor{contract: contract}, VRFV25ClientFilterer: VRFV25ClientFilterer{contract: contract}}, nil
}

func NewVRFV25ClientCaller(address common.Address, caller bind.ContractCaller) (*VRFV25ClientCaller, error) {
	contract, err := bindVRFV25Client(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV25ClientCaller{contract: contract}, nil
}

func NewVRFV25ClientTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFV25ClientTransactor, error) {
	contract, err := bindVRFV25Client(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV25ClientTransactor{contract: contract}, nil
}

func NewVRFV25ClientFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFV25ClientFilterer, error) {
	contract, err := bindVRFV25Client(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFV25ClientFilterer{contract: contract}, nil
}

func bindVRFV25Client(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFV25ClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFV25Client *VRFV25ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV25Client.Contract.VRFV25ClientCaller.contract.Call(opts, result, method, params...)
}

func (_VRFV25Client *VRFV25ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25Client.Contract.VRFV25ClientTransactor.contract.Transfer(opts)
}

func (_VRFV25Client *VRFV25ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV25Client.Contract.VRFV25ClientTransactor.contract.Transact(opts, method, params...)
}

func (_VRFV25Client *VRFV25ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV25Client.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFV25Client *VRFV25ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25Client.Contract.contract.Transfer(opts)
}

func (_VRFV25Client *VRFV25ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV25Client.Contract.contract.Transact(opts, method, params...)
}

func (_VRFV25Client *VRFV25ClientCaller) EXTRAARGSV1TAG(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _VRFV25Client.contract.Call(opts, &out, "EXTRA_ARGS_V1_TAG")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_VRFV25Client *VRFV25ClientSession) EXTRAARGSV1TAG() ([4]byte, error) {
	return _VRFV25Client.Contract.EXTRAARGSV1TAG(&_VRFV25Client.CallOpts)
}

func (_VRFV25Client *VRFV25ClientCallerSession) EXTRAARGSV1TAG() ([4]byte, error) {
	return _VRFV25Client.Contract.EXTRAARGSV1TAG(&_VRFV25Client.CallOpts)
}

func (_VRFV25Client *VRFV25Client) Address() common.Address {
	return _VRFV25Client.address
}

type VRFV25ClientInterface interface {
	EXTRAARGSV1TAG(opts *bind.CallOpts) ([4]byte, error)

	Address() common.Address
}
