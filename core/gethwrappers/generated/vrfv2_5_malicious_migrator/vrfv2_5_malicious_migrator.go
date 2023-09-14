// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrfv2_5_malicious_migrator

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

var VRFV25MaliciousMigratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516102e03803806102e083398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b61024d806100936000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80638ea9811714610030575b600080fd5b61004361003e36600461012a565b610045565b005b600080546040805160c081018252838152602080820185905281830185905260608201859052608082018590528251908101835293845260a0810193909352517f9b1c385e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911691639b1c385e916100d49190600401610180565b602060405180830381600087803b1580156100ee57600080fd5b505af1158015610102573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101269190610167565b5050565b60006020828403121561013c57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461016057600080fd5b9392505050565b60006020828403121561017957600080fd5b5051919050565b6000602080835283518184015280840151604084015261ffff6040850151166060840152606084015163ffffffff80821660808601528060808701511660a0860152505060a084015160c08085015280518060e086015260005b818110156101f757828101840151868201610100015283016101da565b8181111561020a57600061010083880101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016939093016101000194935050505056fea164736f6c6343000806000a",
}

var VRFV25MaliciousMigratorABI = VRFV25MaliciousMigratorMetaData.ABI

var VRFV25MaliciousMigratorBin = VRFV25MaliciousMigratorMetaData.Bin

func DeployVRFV25MaliciousMigrator(auth *bind.TransactOpts, backend bind.ContractBackend, _vrfCoordinator common.Address) (common.Address, *types.Transaction, *VRFV25MaliciousMigrator, error) {
	parsed, err := VRFV25MaliciousMigratorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFV25MaliciousMigratorBin), backend, _vrfCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFV25MaliciousMigrator{VRFV25MaliciousMigratorCaller: VRFV25MaliciousMigratorCaller{contract: contract}, VRFV25MaliciousMigratorTransactor: VRFV25MaliciousMigratorTransactor{contract: contract}, VRFV25MaliciousMigratorFilterer: VRFV25MaliciousMigratorFilterer{contract: contract}}, nil
}

type VRFV25MaliciousMigrator struct {
	address common.Address
	abi     abi.ABI
	VRFV25MaliciousMigratorCaller
	VRFV25MaliciousMigratorTransactor
	VRFV25MaliciousMigratorFilterer
}

type VRFV25MaliciousMigratorCaller struct {
	contract *bind.BoundContract
}

type VRFV25MaliciousMigratorTransactor struct {
	contract *bind.BoundContract
}

type VRFV25MaliciousMigratorFilterer struct {
	contract *bind.BoundContract
}

type VRFV25MaliciousMigratorSession struct {
	Contract     *VRFV25MaliciousMigrator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFV25MaliciousMigratorCallerSession struct {
	Contract *VRFV25MaliciousMigratorCaller
	CallOpts bind.CallOpts
}

type VRFV25MaliciousMigratorTransactorSession struct {
	Contract     *VRFV25MaliciousMigratorTransactor
	TransactOpts bind.TransactOpts
}

type VRFV25MaliciousMigratorRaw struct {
	Contract *VRFV25MaliciousMigrator
}

type VRFV25MaliciousMigratorCallerRaw struct {
	Contract *VRFV25MaliciousMigratorCaller
}

type VRFV25MaliciousMigratorTransactorRaw struct {
	Contract *VRFV25MaliciousMigratorTransactor
}

func NewVRFV25MaliciousMigrator(address common.Address, backend bind.ContractBackend) (*VRFV25MaliciousMigrator, error) {
	abi, err := abi.JSON(strings.NewReader(VRFV25MaliciousMigratorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFV25MaliciousMigrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFV25MaliciousMigrator{address: address, abi: abi, VRFV25MaliciousMigratorCaller: VRFV25MaliciousMigratorCaller{contract: contract}, VRFV25MaliciousMigratorTransactor: VRFV25MaliciousMigratorTransactor{contract: contract}, VRFV25MaliciousMigratorFilterer: VRFV25MaliciousMigratorFilterer{contract: contract}}, nil
}

func NewVRFV25MaliciousMigratorCaller(address common.Address, caller bind.ContractCaller) (*VRFV25MaliciousMigratorCaller, error) {
	contract, err := bindVRFV25MaliciousMigrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV25MaliciousMigratorCaller{contract: contract}, nil
}

func NewVRFV25MaliciousMigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFV25MaliciousMigratorTransactor, error) {
	contract, err := bindVRFV25MaliciousMigrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV25MaliciousMigratorTransactor{contract: contract}, nil
}

func NewVRFV25MaliciousMigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFV25MaliciousMigratorFilterer, error) {
	contract, err := bindVRFV25MaliciousMigrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFV25MaliciousMigratorFilterer{contract: contract}, nil
}

func bindVRFV25MaliciousMigrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFV25MaliciousMigratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFV25MaliciousMigrator *VRFV25MaliciousMigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV25MaliciousMigrator.Contract.VRFV25MaliciousMigratorCaller.contract.Call(opts, result, method, params...)
}

func (_VRFV25MaliciousMigrator *VRFV25MaliciousMigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25MaliciousMigrator.Contract.VRFV25MaliciousMigratorTransactor.contract.Transfer(opts)
}

func (_VRFV25MaliciousMigrator *VRFV25MaliciousMigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV25MaliciousMigrator.Contract.VRFV25MaliciousMigratorTransactor.contract.Transact(opts, method, params...)
}

func (_VRFV25MaliciousMigrator *VRFV25MaliciousMigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV25MaliciousMigrator.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFV25MaliciousMigrator *VRFV25MaliciousMigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25MaliciousMigrator.Contract.contract.Transfer(opts)
}

func (_VRFV25MaliciousMigrator *VRFV25MaliciousMigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV25MaliciousMigrator.Contract.contract.Transact(opts, method, params...)
}

func (_VRFV25MaliciousMigrator *VRFV25MaliciousMigratorTransactor) SetCoordinator(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _VRFV25MaliciousMigrator.contract.Transact(opts, "setCoordinator", arg0)
}

func (_VRFV25MaliciousMigrator *VRFV25MaliciousMigratorSession) SetCoordinator(arg0 common.Address) (*types.Transaction, error) {
	return _VRFV25MaliciousMigrator.Contract.SetCoordinator(&_VRFV25MaliciousMigrator.TransactOpts, arg0)
}

func (_VRFV25MaliciousMigrator *VRFV25MaliciousMigratorTransactorSession) SetCoordinator(arg0 common.Address) (*types.Transaction, error) {
	return _VRFV25MaliciousMigrator.Contract.SetCoordinator(&_VRFV25MaliciousMigrator.TransactOpts, arg0)
}

func (_VRFV25MaliciousMigrator *VRFV25MaliciousMigrator) Address() common.Address {
	return _VRFV25MaliciousMigrator.address
}

type VRFV25MaliciousMigratorInterface interface {
	SetCoordinator(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error)

	Address() common.Address
}
