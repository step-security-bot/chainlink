// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_malicious_consumer_v2_5

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

var VRFMaliciousConsumerV25MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"name\":\"OnlyOwnerOrCoordinator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"createSubscriptionAndFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_gasAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_randomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_requestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"}],\"name\":\"setCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"name\":\"updateSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001201380380620012018339810160408190526200003491620001cc565b8133806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000103565b5050600280546001600160a01b03199081166001600160a01b0394851617909155600580548216958416959095179094555060068054909316911617905562000204565b6001600160a01b0381163314156200015e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0381168114620001c757600080fd5b919050565b60008060408385031215620001e057600080fd5b620001eb83620001af565b9150620001fb60208401620001af565b90509250929050565b610fed80620002146000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638ea9811711610081578063f08c5daa1161005b578063f08c5daa1461017b578063f2fde38b14610184578063f6eaffc81461019757600080fd5b80638ea981171461014c578063cf62c8ab1461015f578063e89e106a1461017257600080fd5b80635e3b709f116100b25780635e3b709f146100f657806379ba50971461011c5780638da5cb5b1461012457600080fd5b80631fe543e3146100ce57806336bfffed146100e3575b600080fd5b6100e16100dc366004610cc1565b6101aa565b005b6100e16100f1366004610bc9565b610230565b610109610104366004610c8f565b610368565b6040519081526020015b60405180910390f35b6100e161045e565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610113565b6100e161015a366004610bae565b61055b565b6100e161016d366004610d65565b610666565b61010960045481565b61010960075481565b6100e1610192366004610bae565b61086c565b6101096101a5366004610c8f565b610880565b60025473ffffffffffffffffffffffffffffffffffffffff163314610222576002546040517f1cf993f400000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911660248201526044015b60405180910390fd5b61022c82826108a1565b5050565b600854610299576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f7375624944206e6f7420736574000000000000000000000000000000000000006044820152606401610219565b60005b815181101561022c57600554600854835173ffffffffffffffffffffffffffffffffffffffff9092169163bec4c08c91908590859081106102df576102df610f82565b60200260200101516040518363ffffffff1660e01b815260040161032392919091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b600060405180830381600087803b15801561033d57600080fd5b505af1158015610351573d6000803e3d6000fd5b50505050808061036090610f22565b91505061029c565b60098190556040805160c08101825282815260085460208083019190915260018284018190526207a1206060840152608083015282519081018352600080825260a083019190915260055492517f9b1c385e000000000000000000000000000000000000000000000000000000008152909273ffffffffffffffffffffffffffffffffffffffff1690639b1c385e90610405908490600401610e4a565b602060405180830381600087803b15801561041f57600080fd5b505af1158015610433573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104579190610ca8565b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146104df576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610219565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60005473ffffffffffffffffffffffffffffffffffffffff16331480159061059b575060025473ffffffffffffffffffffffffffffffffffffffff163314155b1561061f57336105c060005473ffffffffffffffffffffffffffffffffffffffff1690565b6002546040517f061db9c100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff93841660048201529183166024830152919091166044820152606401610219565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60085461079e57600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a21a23e46040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156106d757600080fd5b505af11580156106eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070f9190610ca8565b60088190556005546040517fbec4c08c000000000000000000000000000000000000000000000000000000008152600481019290925230602483015273ffffffffffffffffffffffffffffffffffffffff169063bec4c08c90604401600060405180830381600087803b15801561078557600080fd5b505af1158015610799573d6000803e3d6000fd5b505050505b6006546005546008546040805160208082019390935281518082039093018352808201918290527f4000aea00000000000000000000000000000000000000000000000000000000090915273ffffffffffffffffffffffffffffffffffffffff93841693634000aea09361081a93911691869190604401610dfe565b602060405180830381600087803b15801561083457600080fd5b505af1158015610848573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022c9190610c6d565b6108746109ac565b61087d81610a2f565b50565b6003818154811061089057600080fd5b600091825260209091200154905081565b5a60075580516108b8906003906020840190610b25565b5060048281556040805160c0810182526009548152600854602080830191909152600182840181905262030d4060608401526080830152825190810183526000815260a082015260055491517f9b1c385e000000000000000000000000000000000000000000000000000000008152909273ffffffffffffffffffffffffffffffffffffffff90921691639b1c385e9161095491859101610e4a565b602060405180830381600087803b15801561096e57600080fd5b505af1158015610982573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a69190610ca8565b50505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610a2d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610219565b565b73ffffffffffffffffffffffffffffffffffffffff8116331415610aaf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610219565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215610b60579160200282015b82811115610b60578251825591602001919060010190610b45565b50610b6c929150610b70565b5090565b5b80821115610b6c5760008155600101610b71565b803573ffffffffffffffffffffffffffffffffffffffff81168114610ba957600080fd5b919050565b600060208284031215610bc057600080fd5b61045782610b85565b60006020808385031215610bdc57600080fd5b823567ffffffffffffffff811115610bf357600080fd5b8301601f81018513610c0457600080fd5b8035610c17610c1282610efe565b610eaf565b80828252848201915084840188868560051b8701011115610c3757600080fd5b600094505b83851015610c6157610c4d81610b85565b835260019490940193918501918501610c3c565b50979650505050505050565b600060208284031215610c7f57600080fd5b8151801515811461045757600080fd5b600060208284031215610ca157600080fd5b5035919050565b600060208284031215610cba57600080fd5b5051919050565b60008060408385031215610cd457600080fd5b8235915060208084013567ffffffffffffffff811115610cf357600080fd5b8401601f81018613610d0457600080fd5b8035610d12610c1282610efe565b80828252848201915084840189868560051b8701011115610d3257600080fd5b600094505b83851015610d55578035835260019490940193918501918501610d37565b5080955050505050509250929050565b600060208284031215610d7757600080fd5b81356bffffffffffffffffffffffff8116811461045757600080fd5b6000815180845260005b81811015610db957602081850181015186830182015201610d9d565b81811115610dcb576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff841681526bffffffffffffffffffffffff83166020820152606060408201526000610e416060830184610d93565b95945050505050565b60208152815160208201526020820151604082015261ffff60408301511660608201526000606083015163ffffffff80821660808501528060808601511660a0850152505060a083015160c080840152610ea760e0840182610d93565b949350505050565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610ef657610ef6610fb1565b604052919050565b600067ffffffffffffffff821115610f1857610f18610fb1565b5060051b60200190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610f7b577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var VRFMaliciousConsumerV25ABI = VRFMaliciousConsumerV25MetaData.ABI

var VRFMaliciousConsumerV25Bin = VRFMaliciousConsumerV25MetaData.Bin

func DeployVRFMaliciousConsumerV25(auth *bind.TransactOpts, backend bind.ContractBackend, vrfCoordinator common.Address, link common.Address) (common.Address, *types.Transaction, *VRFMaliciousConsumerV25, error) {
	parsed, err := VRFMaliciousConsumerV25MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFMaliciousConsumerV25Bin), backend, vrfCoordinator, link)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFMaliciousConsumerV25{VRFMaliciousConsumerV25Caller: VRFMaliciousConsumerV25Caller{contract: contract}, VRFMaliciousConsumerV25Transactor: VRFMaliciousConsumerV25Transactor{contract: contract}, VRFMaliciousConsumerV25Filterer: VRFMaliciousConsumerV25Filterer{contract: contract}}, nil
}

type VRFMaliciousConsumerV25 struct {
	address common.Address
	abi     abi.ABI
	VRFMaliciousConsumerV25Caller
	VRFMaliciousConsumerV25Transactor
	VRFMaliciousConsumerV25Filterer
}

type VRFMaliciousConsumerV25Caller struct {
	contract *bind.BoundContract
}

type VRFMaliciousConsumerV25Transactor struct {
	contract *bind.BoundContract
}

type VRFMaliciousConsumerV25Filterer struct {
	contract *bind.BoundContract
}

type VRFMaliciousConsumerV25Session struct {
	Contract     *VRFMaliciousConsumerV25
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFMaliciousConsumerV25CallerSession struct {
	Contract *VRFMaliciousConsumerV25Caller
	CallOpts bind.CallOpts
}

type VRFMaliciousConsumerV25TransactorSession struct {
	Contract     *VRFMaliciousConsumerV25Transactor
	TransactOpts bind.TransactOpts
}

type VRFMaliciousConsumerV25Raw struct {
	Contract *VRFMaliciousConsumerV25
}

type VRFMaliciousConsumerV25CallerRaw struct {
	Contract *VRFMaliciousConsumerV25Caller
}

type VRFMaliciousConsumerV25TransactorRaw struct {
	Contract *VRFMaliciousConsumerV25Transactor
}

func NewVRFMaliciousConsumerV25(address common.Address, backend bind.ContractBackend) (*VRFMaliciousConsumerV25, error) {
	abi, err := abi.JSON(strings.NewReader(VRFMaliciousConsumerV25ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFMaliciousConsumerV25(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFMaliciousConsumerV25{address: address, abi: abi, VRFMaliciousConsumerV25Caller: VRFMaliciousConsumerV25Caller{contract: contract}, VRFMaliciousConsumerV25Transactor: VRFMaliciousConsumerV25Transactor{contract: contract}, VRFMaliciousConsumerV25Filterer: VRFMaliciousConsumerV25Filterer{contract: contract}}, nil
}

func NewVRFMaliciousConsumerV25Caller(address common.Address, caller bind.ContractCaller) (*VRFMaliciousConsumerV25Caller, error) {
	contract, err := bindVRFMaliciousConsumerV25(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFMaliciousConsumerV25Caller{contract: contract}, nil
}

func NewVRFMaliciousConsumerV25Transactor(address common.Address, transactor bind.ContractTransactor) (*VRFMaliciousConsumerV25Transactor, error) {
	contract, err := bindVRFMaliciousConsumerV25(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFMaliciousConsumerV25Transactor{contract: contract}, nil
}

func NewVRFMaliciousConsumerV25Filterer(address common.Address, filterer bind.ContractFilterer) (*VRFMaliciousConsumerV25Filterer, error) {
	contract, err := bindVRFMaliciousConsumerV25(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFMaliciousConsumerV25Filterer{contract: contract}, nil
}

func bindVRFMaliciousConsumerV25(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFMaliciousConsumerV25MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFMaliciousConsumerV25.Contract.VRFMaliciousConsumerV25Caller.contract.Call(opts, result, method, params...)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.VRFMaliciousConsumerV25Transactor.contract.Transfer(opts)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.VRFMaliciousConsumerV25Transactor.contract.Transact(opts, method, params...)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFMaliciousConsumerV25.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.contract.Transfer(opts)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.contract.Transact(opts, method, params...)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFMaliciousConsumerV25.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) Owner() (common.Address, error) {
	return _VRFMaliciousConsumerV25.Contract.Owner(&_VRFMaliciousConsumerV25.CallOpts)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25CallerSession) Owner() (common.Address, error) {
	return _VRFMaliciousConsumerV25.Contract.Owner(&_VRFMaliciousConsumerV25.CallOpts)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Caller) SGasAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFMaliciousConsumerV25.contract.Call(opts, &out, "s_gasAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) SGasAvailable() (*big.Int, error) {
	return _VRFMaliciousConsumerV25.Contract.SGasAvailable(&_VRFMaliciousConsumerV25.CallOpts)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25CallerSession) SGasAvailable() (*big.Int, error) {
	return _VRFMaliciousConsumerV25.Contract.SGasAvailable(&_VRFMaliciousConsumerV25.CallOpts)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Caller) SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VRFMaliciousConsumerV25.contract.Call(opts, &out, "s_randomWords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _VRFMaliciousConsumerV25.Contract.SRandomWords(&_VRFMaliciousConsumerV25.CallOpts, arg0)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25CallerSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _VRFMaliciousConsumerV25.Contract.SRandomWords(&_VRFMaliciousConsumerV25.CallOpts, arg0)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Caller) SRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFMaliciousConsumerV25.contract.Call(opts, &out, "s_requestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) SRequestId() (*big.Int, error) {
	return _VRFMaliciousConsumerV25.Contract.SRequestId(&_VRFMaliciousConsumerV25.CallOpts)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25CallerSession) SRequestId() (*big.Int, error) {
	return _VRFMaliciousConsumerV25.Contract.SRequestId(&_VRFMaliciousConsumerV25.CallOpts)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.contract.Transact(opts, "acceptOwnership")
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) AcceptOwnership() (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.AcceptOwnership(&_VRFMaliciousConsumerV25.TransactOpts)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.AcceptOwnership(&_VRFMaliciousConsumerV25.TransactOpts)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Transactor) CreateSubscriptionAndFund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.contract.Transact(opts, "createSubscriptionAndFund", amount)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) CreateSubscriptionAndFund(amount *big.Int) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.CreateSubscriptionAndFund(&_VRFMaliciousConsumerV25.TransactOpts, amount)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25TransactorSession) CreateSubscriptionAndFund(amount *big.Int) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.CreateSubscriptionAndFund(&_VRFMaliciousConsumerV25.TransactOpts, amount)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Transactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.RawFulfillRandomWords(&_VRFMaliciousConsumerV25.TransactOpts, requestId, randomWords)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25TransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.RawFulfillRandomWords(&_VRFMaliciousConsumerV25.TransactOpts, requestId, randomWords)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Transactor) RequestRandomness(opts *bind.TransactOpts, keyHash [32]byte) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.contract.Transact(opts, "requestRandomness", keyHash)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) RequestRandomness(keyHash [32]byte) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.RequestRandomness(&_VRFMaliciousConsumerV25.TransactOpts, keyHash)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25TransactorSession) RequestRandomness(keyHash [32]byte) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.RequestRandomness(&_VRFMaliciousConsumerV25.TransactOpts, keyHash)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Transactor) SetCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.contract.Transact(opts, "setCoordinator", _vrfCoordinator)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.SetCoordinator(&_VRFMaliciousConsumerV25.TransactOpts, _vrfCoordinator)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25TransactorSession) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.SetCoordinator(&_VRFMaliciousConsumerV25.TransactOpts, _vrfCoordinator)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.TransferOwnership(&_VRFMaliciousConsumerV25.TransactOpts, to)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.TransferOwnership(&_VRFMaliciousConsumerV25.TransactOpts, to)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Transactor) UpdateSubscription(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.contract.Transact(opts, "updateSubscription", consumers)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Session) UpdateSubscription(consumers []common.Address) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.UpdateSubscription(&_VRFMaliciousConsumerV25.TransactOpts, consumers)
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25TransactorSession) UpdateSubscription(consumers []common.Address) (*types.Transaction, error) {
	return _VRFMaliciousConsumerV25.Contract.UpdateSubscription(&_VRFMaliciousConsumerV25.TransactOpts, consumers)
}

type VRFMaliciousConsumerV25OwnershipTransferRequestedIterator struct {
	Event *VRFMaliciousConsumerV25OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFMaliciousConsumerV25OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFMaliciousConsumerV25OwnershipTransferRequested)
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
		it.Event = new(VRFMaliciousConsumerV25OwnershipTransferRequested)
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

func (it *VRFMaliciousConsumerV25OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFMaliciousConsumerV25OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFMaliciousConsumerV25OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFMaliciousConsumerV25OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFMaliciousConsumerV25.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFMaliciousConsumerV25OwnershipTransferRequestedIterator{contract: _VRFMaliciousConsumerV25.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFMaliciousConsumerV25OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFMaliciousConsumerV25.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFMaliciousConsumerV25OwnershipTransferRequested)
				if err := _VRFMaliciousConsumerV25.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Filterer) ParseOwnershipTransferRequested(log types.Log) (*VRFMaliciousConsumerV25OwnershipTransferRequested, error) {
	event := new(VRFMaliciousConsumerV25OwnershipTransferRequested)
	if err := _VRFMaliciousConsumerV25.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFMaliciousConsumerV25OwnershipTransferredIterator struct {
	Event *VRFMaliciousConsumerV25OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFMaliciousConsumerV25OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFMaliciousConsumerV25OwnershipTransferred)
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
		it.Event = new(VRFMaliciousConsumerV25OwnershipTransferred)
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

func (it *VRFMaliciousConsumerV25OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFMaliciousConsumerV25OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFMaliciousConsumerV25OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFMaliciousConsumerV25OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFMaliciousConsumerV25.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFMaliciousConsumerV25OwnershipTransferredIterator{contract: _VRFMaliciousConsumerV25.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFMaliciousConsumerV25OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFMaliciousConsumerV25.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFMaliciousConsumerV25OwnershipTransferred)
				if err := _VRFMaliciousConsumerV25.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25Filterer) ParseOwnershipTransferred(log types.Log) (*VRFMaliciousConsumerV25OwnershipTransferred, error) {
	event := new(VRFMaliciousConsumerV25OwnershipTransferred)
	if err := _VRFMaliciousConsumerV25.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFMaliciousConsumerV25.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFMaliciousConsumerV25.ParseOwnershipTransferRequested(log)
	case _VRFMaliciousConsumerV25.abi.Events["OwnershipTransferred"].ID:
		return _VRFMaliciousConsumerV25.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFMaliciousConsumerV25OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFMaliciousConsumerV25OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_VRFMaliciousConsumerV25 *VRFMaliciousConsumerV25) Address() common.Address {
	return _VRFMaliciousConsumerV25.address
}

type VRFMaliciousConsumerV25Interface interface {
	Owner(opts *bind.CallOpts) (common.Address, error)

	SGasAvailable(opts *bind.CallOpts) (*big.Int, error)

	SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)

	SRequestId(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CreateSubscriptionAndFund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error)

	RequestRandomness(opts *bind.TransactOpts, keyHash [32]byte) (*types.Transaction, error)

	SetCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateSubscription(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFMaliciousConsumerV25OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFMaliciousConsumerV25OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFMaliciousConsumerV25OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFMaliciousConsumerV25OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFMaliciousConsumerV25OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFMaliciousConsumerV25OwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
