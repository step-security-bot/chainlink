// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_consumer_v2_5_upgradeable_example

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

var VRFConsumerV25UpgradeableExampleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COORDINATOR\",\"outputs\":[{\"internalType\":\"contractIVRFCoordinatorV2_5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINKTOKEN\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"createSubscriptionAndFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"minReqConfs\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_gasAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_randomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_requestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_subId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"topUpSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"name\":\"updateSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506110c1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c806355380dfb11610081578063e89e106a1161005b578063e89e106a146101ce578063f08c5daa146101d7578063f6eaffc8146101e057600080fd5b806355380dfb14610192578063706da1ca146101b2578063cf62c8ab146101bb57600080fd5b806336bfffed116100b257806336bfffed146101275780633b2bcbf11461013a578063485cc9551461017f57600080fd5b80631fe543e3146100d95780632e75964e146100ee5780632fa4e44214610114575b600080fd5b6100ec6100e7366004610d95565b6101f3565b005b6101016100fc366004610d03565b610284565b6040519081526020015b60405180910390f35b6100ec610122366004610e39565b610381565b6100ec610135366004610c36565b6104a3565b60345461015a9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010b565b6100ec61018d366004610c03565b6105db565b60355461015a9073ffffffffffffffffffffffffffffffffffffffff1681565b61010160365481565b6100ec6101c9366004610e39565b6107c5565b61010160335481565b61010160375481565b6101016101ee366004610d63565b61093c565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff163314610276576000546040517f1cf993f40000000000000000000000000000000000000000000000000000000081523360048201526201000090910473ffffffffffffffffffffffffffffffffffffffff1660248201526044015b60405180910390fd5b610280828261095d565b5050565b6040805160c081018252868152602080820187905261ffff86168284015263ffffffff80861660608401528416608083015282519081018352600080825260a083019190915260345492517f9b1c385e000000000000000000000000000000000000000000000000000000008152909273ffffffffffffffffffffffffffffffffffffffff1690639b1c385e9061031f908490600401610f1e565b602060405180830381600087803b15801561033957600080fd5b505af115801561034d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103719190610d7c565b6033819055979650505050505050565b6036546103ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f737562206e6f7420736574000000000000000000000000000000000000000000604482015260640161026d565b60355460345460365460408051602081019290925273ffffffffffffffffffffffffffffffffffffffff93841693634000aea09316918591015b6040516020818303038152906040526040518463ffffffff1660e01b815260040161045193929190610ed2565b602060405180830381600087803b15801561046b57600080fd5b505af115801561047f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102809190610cda565b60365461050c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f7375624944206e6f742073657400000000000000000000000000000000000000604482015260640161026d565b60005b815181101561028057603454603654835173ffffffffffffffffffffffffffffffffffffffff9092169163bec4c08c919085908590811061055257610552611056565b60200260200101516040518363ffffffff1660e01b815260040161059692919091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b600060405180830381600087803b1580156105b057600080fd5b505af11580156105c4573d6000803e3d6000fd5b5050505080806105d390610ff6565b91505061050f565b600054610100900460ff16158080156105fb5750600054600160ff909116105b806106155750303b158015610615575060005460ff166001145b6106a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161026d565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156106ff57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610708836109df565b6034805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255603580549285169290911691909117905580156107c057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6036546103ea57603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a21a23e46040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561083657600080fd5b505af115801561084a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086e9190610d7c565b60368190556034546040517fbec4c08c000000000000000000000000000000000000000000000000000000008152600481019290925230602483015273ffffffffffffffffffffffffffffffffffffffff169063bec4c08c90604401600060405180830381600087803b1580156108e457600080fd5b505af11580156108f8573d6000803e3d6000fd5b5050505060355460345460365460405173ffffffffffffffffffffffffffffffffffffffff93841693634000aea09316918591610424919060200190815260200190565b6032818154811061094c57600080fd5b600091825260209091200154905081565b60335482146109c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265717565737420494420697320696e636f7272656374000000000000000000604482015260640161026d565b5a60375580516107c0906032906020840190610b66565b600054610100900460ff16610a76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161026d565b73ffffffffffffffffffffffffffffffffffffffff8116610b19576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f6d75737420676976652076616c696420636f6f7264696e61746f72206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015260840161026d565b6000805473ffffffffffffffffffffffffffffffffffffffff90921662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff909216919091179055565b828054828255906000526020600020908101928215610ba1579160200282015b82811115610ba1578251825591602001919060010190610b86565b50610bad929150610bb1565b5090565b5b80821115610bad5760008155600101610bb2565b803573ffffffffffffffffffffffffffffffffffffffff81168114610bea57600080fd5b919050565b803563ffffffff81168114610bea57600080fd5b60008060408385031215610c1657600080fd5b610c1f83610bc6565b9150610c2d60208401610bc6565b90509250929050565b60006020808385031215610c4957600080fd5b823567ffffffffffffffff811115610c6057600080fd5b8301601f81018513610c7157600080fd5b8035610c84610c7f82610fd2565b610f83565b80828252848201915084840188868560051b8701011115610ca457600080fd5b600094505b83851015610cce57610cba81610bc6565b835260019490940193918501918501610ca9565b50979650505050505050565b600060208284031215610cec57600080fd5b81518015158114610cfc57600080fd5b9392505050565b600080600080600060a08688031215610d1b57600080fd5b8535945060208601359350604086013561ffff81168114610d3b57600080fd5b9250610d4960608701610bef565b9150610d5760808701610bef565b90509295509295909350565b600060208284031215610d7557600080fd5b5035919050565b600060208284031215610d8e57600080fd5b5051919050565b60008060408385031215610da857600080fd5b8235915060208084013567ffffffffffffffff811115610dc757600080fd5b8401601f81018613610dd857600080fd5b8035610de6610c7f82610fd2565b80828252848201915084840189868560051b8701011115610e0657600080fd5b600094505b83851015610e29578035835260019490940193918501918501610e0b565b5080955050505050509250929050565b600060208284031215610e4b57600080fd5b81356bffffffffffffffffffffffff81168114610cfc57600080fd5b6000815180845260005b81811015610e8d57602081850181015186830182015201610e71565b81811115610e9f576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff841681526bffffffffffffffffffffffff83166020820152606060408201526000610f156060830184610e67565b95945050505050565b60208152815160208201526020820151604082015261ffff60408301511660608201526000606083015163ffffffff80821660808501528060808601511660a0850152505060a083015160c080840152610f7b60e0840182610e67565b949350505050565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610fca57610fca611085565b604052919050565b600067ffffffffffffffff821115610fec57610fec611085565b5060051b60200190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561104f577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var VRFConsumerV25UpgradeableExampleABI = VRFConsumerV25UpgradeableExampleMetaData.ABI

var VRFConsumerV25UpgradeableExampleBin = VRFConsumerV25UpgradeableExampleMetaData.Bin

func DeployVRFConsumerV25UpgradeableExample(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFConsumerV25UpgradeableExample, error) {
	parsed, err := VRFConsumerV25UpgradeableExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFConsumerV25UpgradeableExampleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFConsumerV25UpgradeableExample{VRFConsumerV25UpgradeableExampleCaller: VRFConsumerV25UpgradeableExampleCaller{contract: contract}, VRFConsumerV25UpgradeableExampleTransactor: VRFConsumerV25UpgradeableExampleTransactor{contract: contract}, VRFConsumerV25UpgradeableExampleFilterer: VRFConsumerV25UpgradeableExampleFilterer{contract: contract}}, nil
}

type VRFConsumerV25UpgradeableExample struct {
	address common.Address
	abi     abi.ABI
	VRFConsumerV25UpgradeableExampleCaller
	VRFConsumerV25UpgradeableExampleTransactor
	VRFConsumerV25UpgradeableExampleFilterer
}

type VRFConsumerV25UpgradeableExampleCaller struct {
	contract *bind.BoundContract
}

type VRFConsumerV25UpgradeableExampleTransactor struct {
	contract *bind.BoundContract
}

type VRFConsumerV25UpgradeableExampleFilterer struct {
	contract *bind.BoundContract
}

type VRFConsumerV25UpgradeableExampleSession struct {
	Contract     *VRFConsumerV25UpgradeableExample
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFConsumerV25UpgradeableExampleCallerSession struct {
	Contract *VRFConsumerV25UpgradeableExampleCaller
	CallOpts bind.CallOpts
}

type VRFConsumerV25UpgradeableExampleTransactorSession struct {
	Contract     *VRFConsumerV25UpgradeableExampleTransactor
	TransactOpts bind.TransactOpts
}

type VRFConsumerV25UpgradeableExampleRaw struct {
	Contract *VRFConsumerV25UpgradeableExample
}

type VRFConsumerV25UpgradeableExampleCallerRaw struct {
	Contract *VRFConsumerV25UpgradeableExampleCaller
}

type VRFConsumerV25UpgradeableExampleTransactorRaw struct {
	Contract *VRFConsumerV25UpgradeableExampleTransactor
}

func NewVRFConsumerV25UpgradeableExample(address common.Address, backend bind.ContractBackend) (*VRFConsumerV25UpgradeableExample, error) {
	abi, err := abi.JSON(strings.NewReader(VRFConsumerV25UpgradeableExampleABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFConsumerV25UpgradeableExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerV25UpgradeableExample{address: address, abi: abi, VRFConsumerV25UpgradeableExampleCaller: VRFConsumerV25UpgradeableExampleCaller{contract: contract}, VRFConsumerV25UpgradeableExampleTransactor: VRFConsumerV25UpgradeableExampleTransactor{contract: contract}, VRFConsumerV25UpgradeableExampleFilterer: VRFConsumerV25UpgradeableExampleFilterer{contract: contract}}, nil
}

func NewVRFConsumerV25UpgradeableExampleCaller(address common.Address, caller bind.ContractCaller) (*VRFConsumerV25UpgradeableExampleCaller, error) {
	contract, err := bindVRFConsumerV25UpgradeableExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerV25UpgradeableExampleCaller{contract: contract}, nil
}

func NewVRFConsumerV25UpgradeableExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFConsumerV25UpgradeableExampleTransactor, error) {
	contract, err := bindVRFConsumerV25UpgradeableExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerV25UpgradeableExampleTransactor{contract: contract}, nil
}

func NewVRFConsumerV25UpgradeableExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFConsumerV25UpgradeableExampleFilterer, error) {
	contract, err := bindVRFConsumerV25UpgradeableExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerV25UpgradeableExampleFilterer{contract: contract}, nil
}

func bindVRFConsumerV25UpgradeableExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFConsumerV25UpgradeableExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFConsumerV25UpgradeableExample.Contract.VRFConsumerV25UpgradeableExampleCaller.contract.Call(opts, result, method, params...)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.VRFConsumerV25UpgradeableExampleTransactor.contract.Transfer(opts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.VRFConsumerV25UpgradeableExampleTransactor.contract.Transact(opts, method, params...)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFConsumerV25UpgradeableExample.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.contract.Transfer(opts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.contract.Transact(opts, method, params...)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCaller) COORDINATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFConsumerV25UpgradeableExample.contract.Call(opts, &out, "COORDINATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) COORDINATOR() (common.Address, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.COORDINATOR(&_VRFConsumerV25UpgradeableExample.CallOpts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCallerSession) COORDINATOR() (common.Address, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.COORDINATOR(&_VRFConsumerV25UpgradeableExample.CallOpts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCaller) LINKTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFConsumerV25UpgradeableExample.contract.Call(opts, &out, "LINKTOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) LINKTOKEN() (common.Address, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.LINKTOKEN(&_VRFConsumerV25UpgradeableExample.CallOpts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCallerSession) LINKTOKEN() (common.Address, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.LINKTOKEN(&_VRFConsumerV25UpgradeableExample.CallOpts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCaller) SGasAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFConsumerV25UpgradeableExample.contract.Call(opts, &out, "s_gasAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) SGasAvailable() (*big.Int, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.SGasAvailable(&_VRFConsumerV25UpgradeableExample.CallOpts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCallerSession) SGasAvailable() (*big.Int, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.SGasAvailable(&_VRFConsumerV25UpgradeableExample.CallOpts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCaller) SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VRFConsumerV25UpgradeableExample.contract.Call(opts, &out, "s_randomWords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.SRandomWords(&_VRFConsumerV25UpgradeableExample.CallOpts, arg0)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCallerSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.SRandomWords(&_VRFConsumerV25UpgradeableExample.CallOpts, arg0)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCaller) SRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFConsumerV25UpgradeableExample.contract.Call(opts, &out, "s_requestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) SRequestId() (*big.Int, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.SRequestId(&_VRFConsumerV25UpgradeableExample.CallOpts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCallerSession) SRequestId() (*big.Int, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.SRequestId(&_VRFConsumerV25UpgradeableExample.CallOpts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCaller) SSubId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFConsumerV25UpgradeableExample.contract.Call(opts, &out, "s_subId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) SSubId() (*big.Int, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.SSubId(&_VRFConsumerV25UpgradeableExample.CallOpts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleCallerSession) SSubId() (*big.Int, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.SSubId(&_VRFConsumerV25UpgradeableExample.CallOpts)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactor) CreateSubscriptionAndFund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.contract.Transact(opts, "createSubscriptionAndFund", amount)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) CreateSubscriptionAndFund(amount *big.Int) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.CreateSubscriptionAndFund(&_VRFConsumerV25UpgradeableExample.TransactOpts, amount)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactorSession) CreateSubscriptionAndFund(amount *big.Int) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.CreateSubscriptionAndFund(&_VRFConsumerV25UpgradeableExample.TransactOpts, amount)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactor) Initialize(opts *bind.TransactOpts, _vrfCoordinator common.Address, _link common.Address) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.contract.Transact(opts, "initialize", _vrfCoordinator, _link)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) Initialize(_vrfCoordinator common.Address, _link common.Address) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.Initialize(&_VRFConsumerV25UpgradeableExample.TransactOpts, _vrfCoordinator, _link)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactorSession) Initialize(_vrfCoordinator common.Address, _link common.Address) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.Initialize(&_VRFConsumerV25UpgradeableExample.TransactOpts, _vrfCoordinator, _link)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.RawFulfillRandomWords(&_VRFConsumerV25UpgradeableExample.TransactOpts, requestId, randomWords)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.RawFulfillRandomWords(&_VRFConsumerV25UpgradeableExample.TransactOpts, requestId, randomWords)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactor) RequestRandomness(opts *bind.TransactOpts, keyHash [32]byte, subId *big.Int, minReqConfs uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.contract.Transact(opts, "requestRandomness", keyHash, subId, minReqConfs, callbackGasLimit, numWords)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) RequestRandomness(keyHash [32]byte, subId *big.Int, minReqConfs uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.RequestRandomness(&_VRFConsumerV25UpgradeableExample.TransactOpts, keyHash, subId, minReqConfs, callbackGasLimit, numWords)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactorSession) RequestRandomness(keyHash [32]byte, subId *big.Int, minReqConfs uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.RequestRandomness(&_VRFConsumerV25UpgradeableExample.TransactOpts, keyHash, subId, minReqConfs, callbackGasLimit, numWords)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactor) TopUpSubscription(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.contract.Transact(opts, "topUpSubscription", amount)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) TopUpSubscription(amount *big.Int) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.TopUpSubscription(&_VRFConsumerV25UpgradeableExample.TransactOpts, amount)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactorSession) TopUpSubscription(amount *big.Int) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.TopUpSubscription(&_VRFConsumerV25UpgradeableExample.TransactOpts, amount)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactor) UpdateSubscription(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.contract.Transact(opts, "updateSubscription", consumers)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleSession) UpdateSubscription(consumers []common.Address) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.UpdateSubscription(&_VRFConsumerV25UpgradeableExample.TransactOpts, consumers)
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleTransactorSession) UpdateSubscription(consumers []common.Address) (*types.Transaction, error) {
	return _VRFConsumerV25UpgradeableExample.Contract.UpdateSubscription(&_VRFConsumerV25UpgradeableExample.TransactOpts, consumers)
}

type VRFConsumerV25UpgradeableExampleInitializedIterator struct {
	Event *VRFConsumerV25UpgradeableExampleInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFConsumerV25UpgradeableExampleInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFConsumerV25UpgradeableExampleInitialized)
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
		it.Event = new(VRFConsumerV25UpgradeableExampleInitialized)
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

func (it *VRFConsumerV25UpgradeableExampleInitializedIterator) Error() error {
	return it.fail
}

func (it *VRFConsumerV25UpgradeableExampleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFConsumerV25UpgradeableExampleInitialized struct {
	Version uint8
	Raw     types.Log
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleFilterer) FilterInitialized(opts *bind.FilterOpts) (*VRFConsumerV25UpgradeableExampleInitializedIterator, error) {

	logs, sub, err := _VRFConsumerV25UpgradeableExample.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VRFConsumerV25UpgradeableExampleInitializedIterator{contract: _VRFConsumerV25UpgradeableExample.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VRFConsumerV25UpgradeableExampleInitialized) (event.Subscription, error) {

	logs, sub, err := _VRFConsumerV25UpgradeableExample.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFConsumerV25UpgradeableExampleInitialized)
				if err := _VRFConsumerV25UpgradeableExample.contract.UnpackLog(event, "Initialized", log); err != nil {
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

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExampleFilterer) ParseInitialized(log types.Log) (*VRFConsumerV25UpgradeableExampleInitialized, error) {
	event := new(VRFConsumerV25UpgradeableExampleInitialized)
	if err := _VRFConsumerV25UpgradeableExample.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExample) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFConsumerV25UpgradeableExample.abi.Events["Initialized"].ID:
		return _VRFConsumerV25UpgradeableExample.ParseInitialized(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFConsumerV25UpgradeableExampleInitialized) Topic() common.Hash {
	return common.HexToHash("0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498")
}

func (_VRFConsumerV25UpgradeableExample *VRFConsumerV25UpgradeableExample) Address() common.Address {
	return _VRFConsumerV25UpgradeableExample.address
}

type VRFConsumerV25UpgradeableExampleInterface interface {
	COORDINATOR(opts *bind.CallOpts) (common.Address, error)

	LINKTOKEN(opts *bind.CallOpts) (common.Address, error)

	SGasAvailable(opts *bind.CallOpts) (*big.Int, error)

	SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)

	SRequestId(opts *bind.CallOpts) (*big.Int, error)

	SSubId(opts *bind.CallOpts) (*big.Int, error)

	CreateSubscriptionAndFund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _vrfCoordinator common.Address, _link common.Address) (*types.Transaction, error)

	RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error)

	RequestRandomness(opts *bind.TransactOpts, keyHash [32]byte, subId *big.Int, minReqConfs uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error)

	TopUpSubscription(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	UpdateSubscription(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error)

	FilterInitialized(opts *bind.FilterOpts) (*VRFConsumerV25UpgradeableExampleInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *VRFConsumerV25UpgradeableExampleInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*VRFConsumerV25UpgradeableExampleInitialized, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
