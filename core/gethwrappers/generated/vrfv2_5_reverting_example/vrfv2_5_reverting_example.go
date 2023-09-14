// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrfv2_5_reverting_example

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

var VRFV25RevertingExampleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"name\":\"OnlyOwnerOrCoordinator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"createSubscriptionAndFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"minReqConfs\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_gasAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_randomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_requestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_subId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"}],\"name\":\"setCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"topUpSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"name\":\"updateSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620011dd380380620011dd8339810160408190526200003491620001cc565b8133806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000103565b5050600280546001600160a01b03199081166001600160a01b0394851617909155600580548216958416959095179094555060068054909316911617905562000204565b6001600160a01b0381163314156200015e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0381168114620001c757600080fd5b919050565b60008060408385031215620001e057600080fd5b620001eb83620001af565b9150620001fb60208401620001af565b90509250929050565b610fc980620002146000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80638da5cb5b1161008c578063e89e106a11610066578063e89e106a146101a4578063f08c5daa146101ad578063f2fde38b146101b6578063f6eaffc8146101c957600080fd5b80638da5cb5b146101565780638ea981171461017e578063cf62c8ab1461019157600080fd5b806336bfffed116100bd57806336bfffed14610132578063706da1ca1461014557806379ba50971461014e57600080fd5b80631fe543e3146100e45780632e75964e146100f95780632fa4e4421461011f575b600080fd5b6100f76100f2366004610c9d565b6101dc565b005b61010c610107366004610c0b565b610262565b6040519081526020015b60405180910390f35b6100f761012d366004610d41565b61035f565b6100f7610140366004610b45565b610481565b61010c60075481565b6100f76105b9565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610116565b6100f761018c366004610b23565b6106b6565b6100f761019f366004610d41565b6107c1565b61010c60045481565b61010c60085481565b6100f76101c4366004610b23565b610938565b61010c6101d7366004610c6b565b61094c565b60025473ffffffffffffffffffffffffffffffffffffffff163314610254576002546040517f1cf993f400000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911660248201526044015b60405180910390fd5b61025e8282600080fd5b5050565b6040805160c081018252868152602080820187905261ffff86168284015263ffffffff80861660608401528416608083015282519081018352600080825260a083019190915260055492517f9b1c385e000000000000000000000000000000000000000000000000000000008152909273ffffffffffffffffffffffffffffffffffffffff1690639b1c385e906102fd908490600401610e26565b602060405180830381600087803b15801561031757600080fd5b505af115801561032b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034f9190610c84565b6004819055979650505050505050565b6007546103c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f737562206e6f7420736574000000000000000000000000000000000000000000604482015260640161024b565b60065460055460075460408051602081019290925273ffffffffffffffffffffffffffffffffffffffff93841693634000aea09316918591015b6040516020818303038152906040526040518463ffffffff1660e01b815260040161042f93929190610dda565b602060405180830381600087803b15801561044957600080fd5b505af115801561045d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061025e9190610be9565b6007546104ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f7375624944206e6f742073657400000000000000000000000000000000000000604482015260640161024b565b60005b815181101561025e57600554600754835173ffffffffffffffffffffffffffffffffffffffff9092169163bec4c08c919085908590811061053057610530610f5e565b60200260200101516040518363ffffffff1660e01b815260040161057492919091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b600060405180830381600087803b15801561058e57600080fd5b505af11580156105a2573d6000803e3d6000fd5b5050505080806105b190610efe565b9150506104ed565b60015473ffffffffffffffffffffffffffffffffffffffff16331461063a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161024b565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60005473ffffffffffffffffffffffffffffffffffffffff1633148015906106f6575060025473ffffffffffffffffffffffffffffffffffffffff163314155b1561077a573361071b60005473ffffffffffffffffffffffffffffffffffffffff1690565b6002546040517f061db9c100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9384166004820152918316602483015291909116604482015260640161024b565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6007546103c857600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a21a23e46040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561083257600080fd5b505af1158015610846573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086a9190610c84565b60078190556005546040517fbec4c08c000000000000000000000000000000000000000000000000000000008152600481019290925230602483015273ffffffffffffffffffffffffffffffffffffffff169063bec4c08c90604401600060405180830381600087803b1580156108e057600080fd5b505af11580156108f4573d6000803e3d6000fd5b5050505060065460055460075460405173ffffffffffffffffffffffffffffffffffffffff93841693634000aea09316918591610402919060200190815260200190565b61094061096d565b610949816109f0565b50565b6003818154811061095c57600080fd5b600091825260209091200154905081565b60005473ffffffffffffffffffffffffffffffffffffffff1633146109ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161024b565b565b73ffffffffffffffffffffffffffffffffffffffff8116331415610a70576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161024b565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b0a57600080fd5b919050565b803563ffffffff81168114610b0a57600080fd5b600060208284031215610b3557600080fd5b610b3e82610ae6565b9392505050565b60006020808385031215610b5857600080fd5b823567ffffffffffffffff811115610b6f57600080fd5b8301601f81018513610b8057600080fd5b8035610b93610b8e82610eda565b610e8b565b80828252848201915084840188868560051b8701011115610bb357600080fd5b600094505b83851015610bdd57610bc981610ae6565b835260019490940193918501918501610bb8565b50979650505050505050565b600060208284031215610bfb57600080fd5b81518015158114610b3e57600080fd5b600080600080600060a08688031215610c2357600080fd5b8535945060208601359350604086013561ffff81168114610c4357600080fd5b9250610c5160608701610b0f565b9150610c5f60808701610b0f565b90509295509295909350565b600060208284031215610c7d57600080fd5b5035919050565b600060208284031215610c9657600080fd5b5051919050565b60008060408385031215610cb057600080fd5b8235915060208084013567ffffffffffffffff811115610ccf57600080fd5b8401601f81018613610ce057600080fd5b8035610cee610b8e82610eda565b80828252848201915084840189868560051b8701011115610d0e57600080fd5b600094505b83851015610d31578035835260019490940193918501918501610d13565b5080955050505050509250929050565b600060208284031215610d5357600080fd5b81356bffffffffffffffffffffffff81168114610b3e57600080fd5b6000815180845260005b81811015610d9557602081850181015186830182015201610d79565b81811115610da7576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff841681526bffffffffffffffffffffffff83166020820152606060408201526000610e1d6060830184610d6f565b95945050505050565b60208152815160208201526020820151604082015261ffff60408301511660608201526000606083015163ffffffff80821660808501528060808601511660a0850152505060a083015160c080840152610e8360e0840182610d6f565b949350505050565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610ed257610ed2610f8d565b604052919050565b600067ffffffffffffffff821115610ef457610ef4610f8d565b5060051b60200190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610f57577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var VRFV25RevertingExampleABI = VRFV25RevertingExampleMetaData.ABI

var VRFV25RevertingExampleBin = VRFV25RevertingExampleMetaData.Bin

func DeployVRFV25RevertingExample(auth *bind.TransactOpts, backend bind.ContractBackend, vrfCoordinator common.Address, link common.Address) (common.Address, *types.Transaction, *VRFV25RevertingExample, error) {
	parsed, err := VRFV25RevertingExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFV25RevertingExampleBin), backend, vrfCoordinator, link)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFV25RevertingExample{VRFV25RevertingExampleCaller: VRFV25RevertingExampleCaller{contract: contract}, VRFV25RevertingExampleTransactor: VRFV25RevertingExampleTransactor{contract: contract}, VRFV25RevertingExampleFilterer: VRFV25RevertingExampleFilterer{contract: contract}}, nil
}

type VRFV25RevertingExample struct {
	address common.Address
	abi     abi.ABI
	VRFV25RevertingExampleCaller
	VRFV25RevertingExampleTransactor
	VRFV25RevertingExampleFilterer
}

type VRFV25RevertingExampleCaller struct {
	contract *bind.BoundContract
}

type VRFV25RevertingExampleTransactor struct {
	contract *bind.BoundContract
}

type VRFV25RevertingExampleFilterer struct {
	contract *bind.BoundContract
}

type VRFV25RevertingExampleSession struct {
	Contract     *VRFV25RevertingExample
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFV25RevertingExampleCallerSession struct {
	Contract *VRFV25RevertingExampleCaller
	CallOpts bind.CallOpts
}

type VRFV25RevertingExampleTransactorSession struct {
	Contract     *VRFV25RevertingExampleTransactor
	TransactOpts bind.TransactOpts
}

type VRFV25RevertingExampleRaw struct {
	Contract *VRFV25RevertingExample
}

type VRFV25RevertingExampleCallerRaw struct {
	Contract *VRFV25RevertingExampleCaller
}

type VRFV25RevertingExampleTransactorRaw struct {
	Contract *VRFV25RevertingExampleTransactor
}

func NewVRFV25RevertingExample(address common.Address, backend bind.ContractBackend) (*VRFV25RevertingExample, error) {
	abi, err := abi.JSON(strings.NewReader(VRFV25RevertingExampleABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFV25RevertingExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFV25RevertingExample{address: address, abi: abi, VRFV25RevertingExampleCaller: VRFV25RevertingExampleCaller{contract: contract}, VRFV25RevertingExampleTransactor: VRFV25RevertingExampleTransactor{contract: contract}, VRFV25RevertingExampleFilterer: VRFV25RevertingExampleFilterer{contract: contract}}, nil
}

func NewVRFV25RevertingExampleCaller(address common.Address, caller bind.ContractCaller) (*VRFV25RevertingExampleCaller, error) {
	contract, err := bindVRFV25RevertingExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV25RevertingExampleCaller{contract: contract}, nil
}

func NewVRFV25RevertingExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFV25RevertingExampleTransactor, error) {
	contract, err := bindVRFV25RevertingExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV25RevertingExampleTransactor{contract: contract}, nil
}

func NewVRFV25RevertingExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFV25RevertingExampleFilterer, error) {
	contract, err := bindVRFV25RevertingExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFV25RevertingExampleFilterer{contract: contract}, nil
}

func bindVRFV25RevertingExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFV25RevertingExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV25RevertingExample.Contract.VRFV25RevertingExampleCaller.contract.Call(opts, result, method, params...)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.VRFV25RevertingExampleTransactor.contract.Transfer(opts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.VRFV25RevertingExampleTransactor.contract.Transact(opts, method, params...)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV25RevertingExample.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.contract.Transfer(opts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.contract.Transact(opts, method, params...)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV25RevertingExample.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) Owner() (common.Address, error) {
	return _VRFV25RevertingExample.Contract.Owner(&_VRFV25RevertingExample.CallOpts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCallerSession) Owner() (common.Address, error) {
	return _VRFV25RevertingExample.Contract.Owner(&_VRFV25RevertingExample.CallOpts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCaller) SGasAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV25RevertingExample.contract.Call(opts, &out, "s_gasAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) SGasAvailable() (*big.Int, error) {
	return _VRFV25RevertingExample.Contract.SGasAvailable(&_VRFV25RevertingExample.CallOpts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCallerSession) SGasAvailable() (*big.Int, error) {
	return _VRFV25RevertingExample.Contract.SGasAvailable(&_VRFV25RevertingExample.CallOpts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCaller) SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VRFV25RevertingExample.contract.Call(opts, &out, "s_randomWords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _VRFV25RevertingExample.Contract.SRandomWords(&_VRFV25RevertingExample.CallOpts, arg0)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCallerSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _VRFV25RevertingExample.Contract.SRandomWords(&_VRFV25RevertingExample.CallOpts, arg0)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCaller) SRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV25RevertingExample.contract.Call(opts, &out, "s_requestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) SRequestId() (*big.Int, error) {
	return _VRFV25RevertingExample.Contract.SRequestId(&_VRFV25RevertingExample.CallOpts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCallerSession) SRequestId() (*big.Int, error) {
	return _VRFV25RevertingExample.Contract.SRequestId(&_VRFV25RevertingExample.CallOpts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCaller) SSubId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV25RevertingExample.contract.Call(opts, &out, "s_subId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) SSubId() (*big.Int, error) {
	return _VRFV25RevertingExample.Contract.SSubId(&_VRFV25RevertingExample.CallOpts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleCallerSession) SSubId() (*big.Int, error) {
	return _VRFV25RevertingExample.Contract.SSubId(&_VRFV25RevertingExample.CallOpts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25RevertingExample.contract.Transact(opts, "acceptOwnership")
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.AcceptOwnership(&_VRFV25RevertingExample.TransactOpts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.AcceptOwnership(&_VRFV25RevertingExample.TransactOpts)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactor) CreateSubscriptionAndFund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VRFV25RevertingExample.contract.Transact(opts, "createSubscriptionAndFund", amount)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) CreateSubscriptionAndFund(amount *big.Int) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.CreateSubscriptionAndFund(&_VRFV25RevertingExample.TransactOpts, amount)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactorSession) CreateSubscriptionAndFund(amount *big.Int) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.CreateSubscriptionAndFund(&_VRFV25RevertingExample.TransactOpts, amount)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV25RevertingExample.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.RawFulfillRandomWords(&_VRFV25RevertingExample.TransactOpts, requestId, randomWords)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.RawFulfillRandomWords(&_VRFV25RevertingExample.TransactOpts, requestId, randomWords)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactor) RequestRandomness(opts *bind.TransactOpts, keyHash [32]byte, subId *big.Int, minReqConfs uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFV25RevertingExample.contract.Transact(opts, "requestRandomness", keyHash, subId, minReqConfs, callbackGasLimit, numWords)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) RequestRandomness(keyHash [32]byte, subId *big.Int, minReqConfs uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.RequestRandomness(&_VRFV25RevertingExample.TransactOpts, keyHash, subId, minReqConfs, callbackGasLimit, numWords)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactorSession) RequestRandomness(keyHash [32]byte, subId *big.Int, minReqConfs uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.RequestRandomness(&_VRFV25RevertingExample.TransactOpts, keyHash, subId, minReqConfs, callbackGasLimit, numWords)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactor) SetCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV25RevertingExample.contract.Transact(opts, "setCoordinator", _vrfCoordinator)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.SetCoordinator(&_VRFV25RevertingExample.TransactOpts, _vrfCoordinator)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactorSession) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.SetCoordinator(&_VRFV25RevertingExample.TransactOpts, _vrfCoordinator)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactor) TopUpSubscription(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VRFV25RevertingExample.contract.Transact(opts, "topUpSubscription", amount)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) TopUpSubscription(amount *big.Int) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.TopUpSubscription(&_VRFV25RevertingExample.TransactOpts, amount)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactorSession) TopUpSubscription(amount *big.Int) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.TopUpSubscription(&_VRFV25RevertingExample.TransactOpts, amount)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFV25RevertingExample.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.TransferOwnership(&_VRFV25RevertingExample.TransactOpts, to)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.TransferOwnership(&_VRFV25RevertingExample.TransactOpts, to)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactor) UpdateSubscription(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error) {
	return _VRFV25RevertingExample.contract.Transact(opts, "updateSubscription", consumers)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleSession) UpdateSubscription(consumers []common.Address) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.UpdateSubscription(&_VRFV25RevertingExample.TransactOpts, consumers)
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleTransactorSession) UpdateSubscription(consumers []common.Address) (*types.Transaction, error) {
	return _VRFV25RevertingExample.Contract.UpdateSubscription(&_VRFV25RevertingExample.TransactOpts, consumers)
}

type VRFV25RevertingExampleOwnershipTransferRequestedIterator struct {
	Event *VRFV25RevertingExampleOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV25RevertingExampleOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV25RevertingExampleOwnershipTransferRequested)
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
		it.Event = new(VRFV25RevertingExampleOwnershipTransferRequested)
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

func (it *VRFV25RevertingExampleOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFV25RevertingExampleOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV25RevertingExampleOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25RevertingExampleOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25RevertingExample.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV25RevertingExampleOwnershipTransferRequestedIterator{contract: _VRFV25RevertingExample.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV25RevertingExampleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25RevertingExample.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV25RevertingExampleOwnershipTransferRequested)
				if err := _VRFV25RevertingExample.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFV25RevertingExample *VRFV25RevertingExampleFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFV25RevertingExampleOwnershipTransferRequested, error) {
	event := new(VRFV25RevertingExampleOwnershipTransferRequested)
	if err := _VRFV25RevertingExample.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV25RevertingExampleOwnershipTransferredIterator struct {
	Event *VRFV25RevertingExampleOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV25RevertingExampleOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV25RevertingExampleOwnershipTransferred)
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
		it.Event = new(VRFV25RevertingExampleOwnershipTransferred)
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

func (it *VRFV25RevertingExampleOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFV25RevertingExampleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV25RevertingExampleOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25RevertingExampleOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25RevertingExample.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV25RevertingExampleOwnershipTransferredIterator{contract: _VRFV25RevertingExample.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFV25RevertingExample *VRFV25RevertingExampleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV25RevertingExampleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25RevertingExample.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV25RevertingExampleOwnershipTransferred)
				if err := _VRFV25RevertingExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFV25RevertingExample *VRFV25RevertingExampleFilterer) ParseOwnershipTransferred(log types.Log) (*VRFV25RevertingExampleOwnershipTransferred, error) {
	event := new(VRFV25RevertingExampleOwnershipTransferred)
	if err := _VRFV25RevertingExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_VRFV25RevertingExample *VRFV25RevertingExample) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFV25RevertingExample.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFV25RevertingExample.ParseOwnershipTransferRequested(log)
	case _VRFV25RevertingExample.abi.Events["OwnershipTransferred"].ID:
		return _VRFV25RevertingExample.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFV25RevertingExampleOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFV25RevertingExampleOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_VRFV25RevertingExample *VRFV25RevertingExample) Address() common.Address {
	return _VRFV25RevertingExample.address
}

type VRFV25RevertingExampleInterface interface {
	Owner(opts *bind.CallOpts) (common.Address, error)

	SGasAvailable(opts *bind.CallOpts) (*big.Int, error)

	SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)

	SRequestId(opts *bind.CallOpts) (*big.Int, error)

	SSubId(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CreateSubscriptionAndFund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error)

	RequestRandomness(opts *bind.TransactOpts, keyHash [32]byte, subId *big.Int, minReqConfs uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error)

	SetCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error)

	TopUpSubscription(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateSubscription(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25RevertingExampleOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV25RevertingExampleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFV25RevertingExampleOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25RevertingExampleOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV25RevertingExampleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFV25RevertingExampleOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
