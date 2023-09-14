// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrfv2_5_wrapper_consumer_example

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

var VRFV25WrapperConsumerExampleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vrfV2Wrapper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LINKAlreadySet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"WrappedRequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"}],\"name\":\"WrapperRequestMade\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_numWords\",\"type\":\"uint32\"}],\"name\":\"makeRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_numWords\",\"type\":\"uint32\"}],\"name\":\"makeRequestNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"native\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"}],\"name\":\"setLinkToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001651380380620016518339810160408190526200003491620001e2565b3380600084846001600160a01b038216156200006657600080546001600160a01b0319166001600160a01b0384161790555b600180546001600160a01b0319166001600160a01b03928316179055831615159050620000da5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600280546001600160a01b0319166001600160a01b03848116919091179091558116156200010d576200010d8162000118565b50505050506200021a565b6001600160a01b038116331415620001735760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000d1565b600380546001600160a01b0319166001600160a01b03838116918217909255600254604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b80516001600160a01b0381168114620001dd57600080fd5b919050565b60008060408385031215620001f657600080fd5b6200020183620001c5565b91506200021160208401620001c5565b90509250929050565b611427806200022a6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806384276d8111610081578063a168fa891161005b578063a168fa8914610185578063d8a4676f146101d7578063f2fde38b146101f957600080fd5b806384276d81146101375780638da5cb5b1461014a5780639c24ea401461017257600080fd5b80631fe543e3116100b25780631fe543e31461010757806379ba50971461011c5780637a8042bd1461012457600080fd5b80630c09b832146100ce5780631e1a3499146100f4575b600080fd5b6100e16100dc366004611282565b61020c565b6040519081526020015b60405180910390f35b6100e1610102366004611282565b6103d1565b61011a610115366004611193565b61051e565b005b61011a6105d8565b61011a610132366004611161565b6106d9565b61011a610145366004611161565b6107c3565b60025460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100eb565b61011a610180366004611102565b6108b3565b6101ba610193366004611161565b600460205260009081526040902080546001820154600390920154909160ff908116911683565b6040805193845291151560208401521515908201526060016100eb565b6101ea6101e5366004611161565b61094a565b6040516100eb939291906113ca565b61011a610207366004611102565b610a6c565b6000610216610a80565b610221848484610b03565b6001546040517f4306d35400000000000000000000000000000000000000000000000000000000815263ffffffff8716600482015291925060009173ffffffffffffffffffffffffffffffffffffffff90911690634306d3549060240160206040518083038186803b15801561029657600080fd5b505afa1580156102aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ce919061117a565b6040805160808101825282815260006020808301828152845183815280830186528486019081526060850184905288845260048352949092208351815591516001830180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055925180519495509193909261035a926002850192910190611089565b5060609190910151600390910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905560405181815282907f5f56b4c20db9f5b294cbf6f681368de4a992a27e2de2ee702dcf2cbbfa791ec49060200160405180910390a2509392505050565b60006103db610a80565b6103e6848484610d08565b6001546040517f4b16093500000000000000000000000000000000000000000000000000000000815263ffffffff8716600482015291925060009173ffffffffffffffffffffffffffffffffffffffff90911690634b1609359060240160206040518083038186803b15801561045b57600080fd5b505afa15801561046f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610493919061117a565b604080516080810182528281526000602080830182815284518381528083018652848601908152600160608601819052898552600484529590932084518155905194810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169515159590951790945590518051949550919361035a9260028501920190611089565b60015473ffffffffffffffffffffffffffffffffffffffff1633146105ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f6f6e6c79205652462056325f3520777261707065722063616e2066756c66696c60448201527f6c0000000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6105d48282610e7b565b5050565b60035473ffffffffffffffffffffffffffffffffffffffff163314610659576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016105c1565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560038054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b6106e1610a80565b60005473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb61071e60025473ffffffffffffffffffffffffffffffffffffffff1690565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff909116600482015260248101849052604401602060405180830381600087803b15801561078b57600080fd5b505af115801561079f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d4919061113f565b6107cb610a80565b60006107ec60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610843576040519150601f19603f3d011682016040523d82523d6000602084013e610848565b606091505b50509050806105d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f77697468647261774e6174697665206661696c6564000000000000000000000060448201526064016105c1565b60005473ffffffffffffffffffffffffffffffffffffffff1615610903576040517f64f778ae00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60008181526004602052604081205481906060906109c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f72657175657374206e6f7420666f756e6400000000000000000000000000000060448201526064016105c1565b6000848152600460209081526040808320815160808101835281548152600182015460ff16151581850152600282018054845181870281018701865281815292959394860193830182828015610a3957602002820191906000526020600020905b815481526020019060010190808311610a25575b50505091835250506003919091015460ff1615156020918201528151908201516040909201519097919650945092505050565b610a74610a80565b610a7d81610f92565b50565b60025473ffffffffffffffffffffffffffffffffffffffff163314610b01576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016105c1565b565b600080546001546040517f4306d35400000000000000000000000000000000000000000000000000000000815263ffffffff8716600482015273ffffffffffffffffffffffffffffffffffffffff92831692634000aea09216908190634306d3549060240160206040518083038186803b158015610b8057600080fd5b505afa158015610b94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bb8919061117a565b6040805163ffffffff808b16602083015261ffff8a169282019290925290871660608201526080016040516020818303038152906040526040518463ffffffff1660e01b8152600401610c0d93929190611309565b602060405180830381600087803b158015610c2757600080fd5b505af1158015610c3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5f919061113f565b50600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc2a88c36040518163ffffffff1660e01b815260040160206040518083038186803b158015610cc857600080fd5b505afa158015610cdc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d00919061117a565b949350505050565b6001546040517f4b16093500000000000000000000000000000000000000000000000000000000815263ffffffff85166004820152600091829173ffffffffffffffffffffffffffffffffffffffff90911690634b1609359060240160206040518083038186803b158015610d7c57600080fd5b505afa158015610d90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db4919061117a565b6001546040517f62a504fc00000000000000000000000000000000000000000000000000000000815263ffffffff808916600483015261ffff881660248301528616604482015291925073ffffffffffffffffffffffffffffffffffffffff16906362a504fc9083906064016020604051808303818588803b158015610e3957600080fd5b505af1158015610e4d573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610e72919061117a565b95945050505050565b600082815260046020526040902054610ef0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f72657175657374206e6f7420666f756e6400000000000000000000000000000060448201526064016105c1565b6000828152600460209081526040909120600181810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790558251610f4392600290920191840190611089565b50600082815260046020526040908190205490517f6c84e12b4c188e61f1b4727024a5cf05c025fa58467e5eedf763c0744c89da7b91610f8691859185916113a1565b60405180910390a15050565b73ffffffffffffffffffffffffffffffffffffffff8116331415611012576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016105c1565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600254604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b8280548282559060005260206000209081019282156110c4579160200282015b828111156110c45782518255916020019190600101906110a9565b506110d09291506110d4565b5090565b5b808211156110d057600081556001016110d5565b803563ffffffff811681146110fd57600080fd5b919050565b60006020828403121561111457600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461113857600080fd5b9392505050565b60006020828403121561115157600080fd5b8151801515811461113857600080fd5b60006020828403121561117357600080fd5b5035919050565b60006020828403121561118c57600080fd5b5051919050565b600080604083850312156111a657600080fd5b8235915060208084013567ffffffffffffffff808211156111c657600080fd5b818601915086601f8301126111da57600080fd5b8135818111156111ec576111ec6113eb565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f8301168101818110858211171561122f5761122f6113eb565b604052828152858101935084860182860187018b101561124e57600080fd5b600095505b83861015611271578035855260019590950194938601938601611253565b508096505050505050509250929050565b60008060006060848603121561129757600080fd5b6112a0846110e9565b9250602084013561ffff811681146112b757600080fd5b91506112c5604085016110e9565b90509250925092565b600081518084526020808501945080840160005b838110156112fe578151875295820195908201906001016112e2565b509495945050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260006020848184015260606040840152835180606085015260005b818110156113595785810183015185820160800152820161133d565b8181111561136b576000608083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160800195945050505050565b8381526060602082015260006113ba60608301856112ce565b9050826040830152949350505050565b8381528215156020820152606060408201526000610e7260608301846112ce565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var VRFV25WrapperConsumerExampleABI = VRFV25WrapperConsumerExampleMetaData.ABI

var VRFV25WrapperConsumerExampleBin = VRFV25WrapperConsumerExampleMetaData.Bin

func DeployVRFV25WrapperConsumerExample(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _vrfV2Wrapper common.Address) (common.Address, *types.Transaction, *VRFV25WrapperConsumerExample, error) {
	parsed, err := VRFV25WrapperConsumerExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFV25WrapperConsumerExampleBin), backend, _link, _vrfV2Wrapper)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFV25WrapperConsumerExample{VRFV25WrapperConsumerExampleCaller: VRFV25WrapperConsumerExampleCaller{contract: contract}, VRFV25WrapperConsumerExampleTransactor: VRFV25WrapperConsumerExampleTransactor{contract: contract}, VRFV25WrapperConsumerExampleFilterer: VRFV25WrapperConsumerExampleFilterer{contract: contract}}, nil
}

type VRFV25WrapperConsumerExample struct {
	address common.Address
	abi     abi.ABI
	VRFV25WrapperConsumerExampleCaller
	VRFV25WrapperConsumerExampleTransactor
	VRFV25WrapperConsumerExampleFilterer
}

type VRFV25WrapperConsumerExampleCaller struct {
	contract *bind.BoundContract
}

type VRFV25WrapperConsumerExampleTransactor struct {
	contract *bind.BoundContract
}

type VRFV25WrapperConsumerExampleFilterer struct {
	contract *bind.BoundContract
}

type VRFV25WrapperConsumerExampleSession struct {
	Contract     *VRFV25WrapperConsumerExample
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFV25WrapperConsumerExampleCallerSession struct {
	Contract *VRFV25WrapperConsumerExampleCaller
	CallOpts bind.CallOpts
}

type VRFV25WrapperConsumerExampleTransactorSession struct {
	Contract     *VRFV25WrapperConsumerExampleTransactor
	TransactOpts bind.TransactOpts
}

type VRFV25WrapperConsumerExampleRaw struct {
	Contract *VRFV25WrapperConsumerExample
}

type VRFV25WrapperConsumerExampleCallerRaw struct {
	Contract *VRFV25WrapperConsumerExampleCaller
}

type VRFV25WrapperConsumerExampleTransactorRaw struct {
	Contract *VRFV25WrapperConsumerExampleTransactor
}

func NewVRFV25WrapperConsumerExample(address common.Address, backend bind.ContractBackend) (*VRFV25WrapperConsumerExample, error) {
	abi, err := abi.JSON(strings.NewReader(VRFV25WrapperConsumerExampleABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFV25WrapperConsumerExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFV25WrapperConsumerExample{address: address, abi: abi, VRFV25WrapperConsumerExampleCaller: VRFV25WrapperConsumerExampleCaller{contract: contract}, VRFV25WrapperConsumerExampleTransactor: VRFV25WrapperConsumerExampleTransactor{contract: contract}, VRFV25WrapperConsumerExampleFilterer: VRFV25WrapperConsumerExampleFilterer{contract: contract}}, nil
}

func NewVRFV25WrapperConsumerExampleCaller(address common.Address, caller bind.ContractCaller) (*VRFV25WrapperConsumerExampleCaller, error) {
	contract, err := bindVRFV25WrapperConsumerExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV25WrapperConsumerExampleCaller{contract: contract}, nil
}

func NewVRFV25WrapperConsumerExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFV25WrapperConsumerExampleTransactor, error) {
	contract, err := bindVRFV25WrapperConsumerExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV25WrapperConsumerExampleTransactor{contract: contract}, nil
}

func NewVRFV25WrapperConsumerExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFV25WrapperConsumerExampleFilterer, error) {
	contract, err := bindVRFV25WrapperConsumerExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFV25WrapperConsumerExampleFilterer{contract: contract}, nil
}

func bindVRFV25WrapperConsumerExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFV25WrapperConsumerExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV25WrapperConsumerExample.Contract.VRFV25WrapperConsumerExampleCaller.contract.Call(opts, result, method, params...)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.VRFV25WrapperConsumerExampleTransactor.contract.Transfer(opts)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.VRFV25WrapperConsumerExampleTransactor.contract.Transact(opts, method, params...)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV25WrapperConsumerExample.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.contract.Transfer(opts)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.contract.Transact(opts, method, params...)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (GetRequestStatus,

	error) {
	var out []interface{}
	err := _VRFV25WrapperConsumerExample.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(GetRequestStatus)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RandomWords = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) GetRequestStatus(_requestId *big.Int) (GetRequestStatus,

	error) {
	return _VRFV25WrapperConsumerExample.Contract.GetRequestStatus(&_VRFV25WrapperConsumerExample.CallOpts, _requestId)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleCallerSession) GetRequestStatus(_requestId *big.Int) (GetRequestStatus,

	error) {
	return _VRFV25WrapperConsumerExample.Contract.GetRequestStatus(&_VRFV25WrapperConsumerExample.CallOpts, _requestId)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFV25WrapperConsumerExample.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) Owner() (common.Address, error) {
	return _VRFV25WrapperConsumerExample.Contract.Owner(&_VRFV25WrapperConsumerExample.CallOpts)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleCallerSession) Owner() (common.Address, error) {
	return _VRFV25WrapperConsumerExample.Contract.Owner(&_VRFV25WrapperConsumerExample.CallOpts)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleCaller) SRequests(opts *bind.CallOpts, arg0 *big.Int) (SRequests,

	error) {
	var out []interface{}
	err := _VRFV25WrapperConsumerExample.contract.Call(opts, &out, "s_requests", arg0)

	outstruct := new(SRequests)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Native = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) SRequests(arg0 *big.Int) (SRequests,

	error) {
	return _VRFV25WrapperConsumerExample.Contract.SRequests(&_VRFV25WrapperConsumerExample.CallOpts, arg0)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleCallerSession) SRequests(arg0 *big.Int) (SRequests,

	error) {
	return _VRFV25WrapperConsumerExample.Contract.SRequests(&_VRFV25WrapperConsumerExample.CallOpts, arg0)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.contract.Transact(opts, "acceptOwnership")
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.AcceptOwnership(&_VRFV25WrapperConsumerExample.TransactOpts)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.AcceptOwnership(&_VRFV25WrapperConsumerExample.TransactOpts)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactor) MakeRequest(opts *bind.TransactOpts, _callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.contract.Transact(opts, "makeRequest", _callbackGasLimit, _requestConfirmations, _numWords)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) MakeRequest(_callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.MakeRequest(&_VRFV25WrapperConsumerExample.TransactOpts, _callbackGasLimit, _requestConfirmations, _numWords)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactorSession) MakeRequest(_callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.MakeRequest(&_VRFV25WrapperConsumerExample.TransactOpts, _callbackGasLimit, _requestConfirmations, _numWords)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactor) MakeRequestNative(opts *bind.TransactOpts, _callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.contract.Transact(opts, "makeRequestNative", _callbackGasLimit, _requestConfirmations, _numWords)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) MakeRequestNative(_callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.MakeRequestNative(&_VRFV25WrapperConsumerExample.TransactOpts, _callbackGasLimit, _requestConfirmations, _numWords)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactorSession) MakeRequestNative(_callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.MakeRequestNative(&_VRFV25WrapperConsumerExample.TransactOpts, _callbackGasLimit, _requestConfirmations, _numWords)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.contract.Transact(opts, "rawFulfillRandomWords", _requestId, _randomWords)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) RawFulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.RawFulfillRandomWords(&_VRFV25WrapperConsumerExample.TransactOpts, _requestId, _randomWords)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactorSession) RawFulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.RawFulfillRandomWords(&_VRFV25WrapperConsumerExample.TransactOpts, _requestId, _randomWords)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactor) SetLinkToken(opts *bind.TransactOpts, _link common.Address) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.contract.Transact(opts, "setLinkToken", _link)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) SetLinkToken(_link common.Address) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.SetLinkToken(&_VRFV25WrapperConsumerExample.TransactOpts, _link)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactorSession) SetLinkToken(_link common.Address) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.SetLinkToken(&_VRFV25WrapperConsumerExample.TransactOpts, _link)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.TransferOwnership(&_VRFV25WrapperConsumerExample.TransactOpts, to)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.TransferOwnership(&_VRFV25WrapperConsumerExample.TransactOpts, to)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactor) WithdrawLink(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.contract.Transact(opts, "withdrawLink", amount)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) WithdrawLink(amount *big.Int) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.WithdrawLink(&_VRFV25WrapperConsumerExample.TransactOpts, amount)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactorSession) WithdrawLink(amount *big.Int) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.WithdrawLink(&_VRFV25WrapperConsumerExample.TransactOpts, amount)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactor) WithdrawNative(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.contract.Transact(opts, "withdrawNative", amount)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleSession) WithdrawNative(amount *big.Int) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.WithdrawNative(&_VRFV25WrapperConsumerExample.TransactOpts, amount)
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleTransactorSession) WithdrawNative(amount *big.Int) (*types.Transaction, error) {
	return _VRFV25WrapperConsumerExample.Contract.WithdrawNative(&_VRFV25WrapperConsumerExample.TransactOpts, amount)
}

type VRFV25WrapperConsumerExampleOwnershipTransferRequestedIterator struct {
	Event *VRFV25WrapperConsumerExampleOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV25WrapperConsumerExampleOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV25WrapperConsumerExampleOwnershipTransferRequested)
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
		it.Event = new(VRFV25WrapperConsumerExampleOwnershipTransferRequested)
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

func (it *VRFV25WrapperConsumerExampleOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFV25WrapperConsumerExampleOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV25WrapperConsumerExampleOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25WrapperConsumerExampleOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25WrapperConsumerExample.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV25WrapperConsumerExampleOwnershipTransferRequestedIterator{contract: _VRFV25WrapperConsumerExample.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV25WrapperConsumerExampleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25WrapperConsumerExample.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV25WrapperConsumerExampleOwnershipTransferRequested)
				if err := _VRFV25WrapperConsumerExample.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFV25WrapperConsumerExampleOwnershipTransferRequested, error) {
	event := new(VRFV25WrapperConsumerExampleOwnershipTransferRequested)
	if err := _VRFV25WrapperConsumerExample.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV25WrapperConsumerExampleOwnershipTransferredIterator struct {
	Event *VRFV25WrapperConsumerExampleOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV25WrapperConsumerExampleOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV25WrapperConsumerExampleOwnershipTransferred)
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
		it.Event = new(VRFV25WrapperConsumerExampleOwnershipTransferred)
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

func (it *VRFV25WrapperConsumerExampleOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFV25WrapperConsumerExampleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV25WrapperConsumerExampleOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25WrapperConsumerExampleOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25WrapperConsumerExample.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFV25WrapperConsumerExampleOwnershipTransferredIterator{contract: _VRFV25WrapperConsumerExample.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV25WrapperConsumerExampleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFV25WrapperConsumerExample.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV25WrapperConsumerExampleOwnershipTransferred)
				if err := _VRFV25WrapperConsumerExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) ParseOwnershipTransferred(log types.Log) (*VRFV25WrapperConsumerExampleOwnershipTransferred, error) {
	event := new(VRFV25WrapperConsumerExampleOwnershipTransferred)
	if err := _VRFV25WrapperConsumerExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV25WrapperConsumerExampleWrappedRequestFulfilledIterator struct {
	Event *VRFV25WrapperConsumerExampleWrappedRequestFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV25WrapperConsumerExampleWrappedRequestFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV25WrapperConsumerExampleWrappedRequestFulfilled)
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
		it.Event = new(VRFV25WrapperConsumerExampleWrappedRequestFulfilled)
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

func (it *VRFV25WrapperConsumerExampleWrappedRequestFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFV25WrapperConsumerExampleWrappedRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV25WrapperConsumerExampleWrappedRequestFulfilled struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Payment     *big.Int
	Raw         types.Log
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) FilterWrappedRequestFulfilled(opts *bind.FilterOpts) (*VRFV25WrapperConsumerExampleWrappedRequestFulfilledIterator, error) {

	logs, sub, err := _VRFV25WrapperConsumerExample.contract.FilterLogs(opts, "WrappedRequestFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFV25WrapperConsumerExampleWrappedRequestFulfilledIterator{contract: _VRFV25WrapperConsumerExample.contract, event: "WrappedRequestFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) WatchWrappedRequestFulfilled(opts *bind.WatchOpts, sink chan<- *VRFV25WrapperConsumerExampleWrappedRequestFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFV25WrapperConsumerExample.contract.WatchLogs(opts, "WrappedRequestFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV25WrapperConsumerExampleWrappedRequestFulfilled)
				if err := _VRFV25WrapperConsumerExample.contract.UnpackLog(event, "WrappedRequestFulfilled", log); err != nil {
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

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) ParseWrappedRequestFulfilled(log types.Log) (*VRFV25WrapperConsumerExampleWrappedRequestFulfilled, error) {
	event := new(VRFV25WrapperConsumerExampleWrappedRequestFulfilled)
	if err := _VRFV25WrapperConsumerExample.contract.UnpackLog(event, "WrappedRequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFV25WrapperConsumerExampleWrapperRequestMadeIterator struct {
	Event *VRFV25WrapperConsumerExampleWrapperRequestMade

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFV25WrapperConsumerExampleWrapperRequestMadeIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFV25WrapperConsumerExampleWrapperRequestMade)
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
		it.Event = new(VRFV25WrapperConsumerExampleWrapperRequestMade)
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

func (it *VRFV25WrapperConsumerExampleWrapperRequestMadeIterator) Error() error {
	return it.fail
}

func (it *VRFV25WrapperConsumerExampleWrapperRequestMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFV25WrapperConsumerExampleWrapperRequestMade struct {
	RequestId *big.Int
	Paid      *big.Int
	Raw       types.Log
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) FilterWrapperRequestMade(opts *bind.FilterOpts, requestId []*big.Int) (*VRFV25WrapperConsumerExampleWrapperRequestMadeIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFV25WrapperConsumerExample.contract.FilterLogs(opts, "WrapperRequestMade", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFV25WrapperConsumerExampleWrapperRequestMadeIterator{contract: _VRFV25WrapperConsumerExample.contract, event: "WrapperRequestMade", logs: logs, sub: sub}, nil
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) WatchWrapperRequestMade(opts *bind.WatchOpts, sink chan<- *VRFV25WrapperConsumerExampleWrapperRequestMade, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFV25WrapperConsumerExample.contract.WatchLogs(opts, "WrapperRequestMade", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFV25WrapperConsumerExampleWrapperRequestMade)
				if err := _VRFV25WrapperConsumerExample.contract.UnpackLog(event, "WrapperRequestMade", log); err != nil {
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

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExampleFilterer) ParseWrapperRequestMade(log types.Log) (*VRFV25WrapperConsumerExampleWrapperRequestMade, error) {
	event := new(VRFV25WrapperConsumerExampleWrapperRequestMade)
	if err := _VRFV25WrapperConsumerExample.contract.UnpackLog(event, "WrapperRequestMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetRequestStatus struct {
	Paid        *big.Int
	Fulfilled   bool
	RandomWords []*big.Int
}
type SRequests struct {
	Paid      *big.Int
	Fulfilled bool
	Native    bool
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExample) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFV25WrapperConsumerExample.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFV25WrapperConsumerExample.ParseOwnershipTransferRequested(log)
	case _VRFV25WrapperConsumerExample.abi.Events["OwnershipTransferred"].ID:
		return _VRFV25WrapperConsumerExample.ParseOwnershipTransferred(log)
	case _VRFV25WrapperConsumerExample.abi.Events["WrappedRequestFulfilled"].ID:
		return _VRFV25WrapperConsumerExample.ParseWrappedRequestFulfilled(log)
	case _VRFV25WrapperConsumerExample.abi.Events["WrapperRequestMade"].ID:
		return _VRFV25WrapperConsumerExample.ParseWrapperRequestMade(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFV25WrapperConsumerExampleOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFV25WrapperConsumerExampleOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFV25WrapperConsumerExampleWrappedRequestFulfilled) Topic() common.Hash {
	return common.HexToHash("0x6c84e12b4c188e61f1b4727024a5cf05c025fa58467e5eedf763c0744c89da7b")
}

func (VRFV25WrapperConsumerExampleWrapperRequestMade) Topic() common.Hash {
	return common.HexToHash("0x5f56b4c20db9f5b294cbf6f681368de4a992a27e2de2ee702dcf2cbbfa791ec4")
}

func (_VRFV25WrapperConsumerExample *VRFV25WrapperConsumerExample) Address() common.Address {
	return _VRFV25WrapperConsumerExample.address
}

type VRFV25WrapperConsumerExampleInterface interface {
	GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (GetRequestStatus,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SRequests(opts *bind.CallOpts, arg0 *big.Int) (SRequests,

		error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	MakeRequest(opts *bind.TransactOpts, _callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32) (*types.Transaction, error)

	MakeRequestNative(opts *bind.TransactOpts, _callbackGasLimit uint32, _requestConfirmations uint16, _numWords uint32) (*types.Transaction, error)

	RawFulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error)

	SetLinkToken(opts *bind.TransactOpts, _link common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawLink(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	WithdrawNative(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25WrapperConsumerExampleOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFV25WrapperConsumerExampleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFV25WrapperConsumerExampleOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFV25WrapperConsumerExampleOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFV25WrapperConsumerExampleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFV25WrapperConsumerExampleOwnershipTransferred, error)

	FilterWrappedRequestFulfilled(opts *bind.FilterOpts) (*VRFV25WrapperConsumerExampleWrappedRequestFulfilledIterator, error)

	WatchWrappedRequestFulfilled(opts *bind.WatchOpts, sink chan<- *VRFV25WrapperConsumerExampleWrappedRequestFulfilled) (event.Subscription, error)

	ParseWrappedRequestFulfilled(log types.Log) (*VRFV25WrapperConsumerExampleWrappedRequestFulfilled, error)

	FilterWrapperRequestMade(opts *bind.FilterOpts, requestId []*big.Int) (*VRFV25WrapperConsumerExampleWrapperRequestMadeIterator, error)

	WatchWrapperRequestMade(opts *bind.WatchOpts, sink chan<- *VRFV25WrapperConsumerExampleWrapperRequestMade, requestId []*big.Int) (event.Subscription, error)

	ParseWrapperRequestMade(log types.Log) (*VRFV25WrapperConsumerExampleWrapperRequestMade, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
