// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_coordinator_v2_5

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

type VRFCoordinatorV25FeeConfig struct {
	FulfillmentFlatFeeLinkPPM uint32
	FulfillmentFlatFeeEthPPM  uint32
}

type VRFCoordinatorV25RequestCommitment struct {
	BlockNum         uint64
	SubId            uint64
	CallbackGasLimit uint32
	NumWords         uint32
	Sender           common.Address
	NativePayment    bool
}

type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

var VRFCoordinatorV25MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blockhashStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"internalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalBalance\",\"type\":\"uint256\"}],\"name\":\"BalanceInvariantViolated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"BlockhashNotInStore\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"GasLimitTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"InsufficientGasForConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"linkWei\",\"type\":\"int256\"}],\"name\":\"InvalidLinkWeiPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"have\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"min\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"max\",\"type\":\"uint16\"}],\"name\":\"InvalidRequestConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoCorrespondingRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"NoSuchProvingKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"NumWordsTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"ProvingKeyAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structVRFCoordinatorV2_5.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthFundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ProvingKeyDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ProvingKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputSeed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"payment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"nativePayment\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"nativePayment\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountEth\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldEthBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEthBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFundedWithEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCKHASH_STORE\",\"outputs\":[{\"internalType\":\"contractBlockhashStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK_ETH_FEED\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_WORDS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REQUEST_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"deregisterProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"nativePayment\",\"type\":\"bool\"}],\"internalType\":\"structVRFCoordinatorV2_5.RequestCommitment\",\"name\":\"rc\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"fundSubscriptionWithEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"ethBalance\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"ownerCancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"pendingRequestExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverEthFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"payInEth\",\"type\":\"bool\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_config\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_currentSubId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_fallbackWeiPerUnitLink\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_feeConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_provingKeyHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"s_provingKeys\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requestCommitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalEthBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"internalType\":\"structVRFCoordinatorV2_5.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"}],\"name\":\"setLINK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"linkEthFeed\",\"type\":\"address\"}],\"name\":\"setLinkEthFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200560338038062005603833981016040819052620000349162000188565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000dc565b505060016002555060601b6001600160601b031916608052620001ba565b6001600160a01b038116331415620001375760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156200019b57600080fd5b81516001600160a01b0381168114620001b357600080fd5b9392505050565b60805160601c615423620001e06000396000818161055d015261370901526154236000f3fe6080604052600436106102db5760003560e01c806386fe91c711610184578063b08c8795116100d6578063e72f6e301161008a578063ee9d2d3811610064578063ee9d2d3814610926578063efcf1d9414610953578063f2fde38b1461097357600080fd5b8063e72f6e30146108af578063e82ad7d4146108cf578063e95704bd146108ff57600080fd5b8063d7ae1d30116100bb578063d7ae1d3014610839578063d98e620e14610859578063da2f26101461087957600080fd5b8063b08c8795146107f9578063caf70c4a1461081957600080fd5b8063a02e061611610138578063a4c0ed3611610112578063a4c0ed3614610799578063a8cb447b146107b9578063ad178361146107d957600080fd5b8063a02e061614610734578063a21a23e414610754578063a47c76961461076957600080fd5b80638da5cb5b116101695780638da5cb5b146106d6578063981837d5146106f45780639f87fad71461071457600080fd5b806386fe91c7146106725780638741e3ee1461069e57600080fd5b806340d6bb821161023d5780636b6feccc116101f157806379ba5097116101cb57806379ba509714610605578063823597401461061a578063827f1c5d1461063a57600080fd5b80636b6feccc1461057f5780636f64f03f146105c55780637341c10c146105e557600080fd5b806364d51a2a1161022257806364d51a2a1461051657806366316d8d1461052b578063689c45171461054b57600080fd5b806340d6bb82146104cb57806346d8d486146104f657600080fd5b806308821d5811610294578063181f5a7711610279578063181f5a77146104345780631b6b6d23146104805780633697af8b146104b857600080fd5b806308821d58146103ec57806315c48b841461040c57600080fd5b8063043bd6ae116102c5578063043bd6ae1461032f57806304c357cb14610353578063088070f51461037357600080fd5b8062012291146102e057806302bcc5b61461030d575b600080fd5b3480156102ec57600080fd5b506102f5610993565b604051610304939291906150b5565b60405180910390f35b34801561031957600080fd5b5061032d610328366004614ef5565b610a0f565b005b34801561033b57600080fd5b50610345600e5481565b604051908152602001610304565b34801561035f57600080fd5b5061032d61036e366004614f10565b610a86565b34801561037f57600080fd5b50600f546103bb9061ffff81169063ffffffff620100008204811691660100000000000081048216916a01000000000000000000009091041684565b6040805161ffff909516855263ffffffff938416602086015291831691840191909152166060820152608001610304565b3480156103f857600080fd5b5061032d610407366004614c97565b610bef565b34801561041857600080fd5b5061042160c881565b60405161ffff9091168152602001610304565b34801561044057600080fd5b50604080518082018252601681527f565246436f6f7264696e61746f72563220322e302e3000000000000000000000602082015290516103049190615060565b34801561048c57600080fd5b506003546104a0906001600160a01b031681565b6040516001600160a01b039091168152602001610304565b61032d6104c6366004614ef5565b610d83565b3480156104d757600080fd5b506104e16101f481565b60405163ffffffff9091168152602001610304565b34801561050257600080fd5b5061032d610511366004614bad565b610f05565b34801561052257600080fd5b50610421606481565b34801561053757600080fd5b5061032d610546366004614bad565b6110a0565b34801561055757600080fd5b506104a07f000000000000000000000000000000000000000000000000000000000000000081565b34801561058b57600080fd5b506010546105a89063ffffffff8082169164010000000090041682565b6040805163ffffffff938416815292909116602083015201610304565b3480156105d157600080fd5b5061032d6105e0366004614be2565b61125d565b3480156105f157600080fd5b5061032d610600366004614f10565b611375565b34801561061157600080fd5b5061032d611558565b34801561062657600080fd5b5061032d610635366004614ef5565b611609565b34801561064657600080fd5b5061065a610655366004614d8f565b611790565b6040516001600160601b039091168152602001610304565b34801561067e57600080fd5b5060075461065a906801000000000000000090046001600160601b031681565b3480156106aa57600080fd5b506007546106be906001600160401b031681565b6040516001600160401b039091168152602001610304565b3480156106e257600080fd5b506000546001600160a01b03166104a0565b34801561070057600080fd5b5061032d61070f366004614b90565b611c43565b34801561072057600080fd5b5061032d61072f366004614f10565b611c6d565b34801561074057600080fd5b5061032d61074f366004614b90565b611fdd565b34801561076057600080fd5b506106be61204a565b34801561077557600080fd5b50610789610784366004614ef5565b612225565b6040516103049493929190615176565b3480156107a557600080fd5b5061032d6107b4366004614c0f565b61231c565b3480156107c557600080fd5b5061032d6107d4366004614b90565b612518565b3480156107e557600080fd5b50600a546104a0906001600160a01b031681565b34801561080557600080fd5b5061032d610814366004614e57565b612634565b34801561082557600080fd5b50610345610834366004614cb3565b6127be565b34801561084557600080fd5b5061032d610854366004614f10565b6127ee565b34801561086557600080fd5b50610345610874366004614cec565b6128d9565b34801561088557600080fd5b506104a0610894366004614cec565b600b602052600090815260409020546001600160a01b031681565b3480156108bb57600080fd5b5061032d6108ca366004614b90565b6128fa565b3480156108db57600080fd5b506108ef6108ea366004614ef5565b612ac2565b6040519015158152602001610304565b34801561090b57600080fd5b5060075461065a90600160a01b90046001600160601b031681565b34801561093257600080fd5b50610345610941366004614cec565b600d6020526000908152604090205481565b34801561095f57600080fd5b5061034561096e366004614d1e565b612ce2565b34801561097f57600080fd5b5061032d61098e366004614b90565b613061565b600f54600c805460408051602080840282018101909252828152600094859460609461ffff8316946201000090930463ffffffff169391928391908301828280156109fd57602002820191906000526020600020905b8154815260200190600101908083116109e9575b50505050509050925092509250909192565b610a17613072565b6001600160401b0381166000908152600560205260409020546001600160a01b0316610a5657604051630fb532db60e11b815260040160405180910390fd5b6001600160401b038116600090815260056020526040902054610a839082906001600160a01b03166130ce565b50565b6001600160401b03821660009081526005602052604090205482906001600160a01b031680610ac857604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b03821614610b0157604051636c51fda960e11b81526001600160a01b03821660048201526024015b60405180910390fd5b600280541415610b415760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b600280556001600160401b0384166000908152600560205260409020600101546001600160a01b03848116911614610be4576001600160401b03841660008181526005602090815260409182902060010180546001600160a01b0319166001600160a01b0388169081179091558251338152918201527f69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be91015b60405180910390a25b505060016002555050565b610bf7613072565b604080518082018252600091610c269190849060029083908390808284376000920191909152506127be915050565b6000818152600b60205260409020549091506001600160a01b031680610c6257604051631dfd6e1360e21b815260048101839052602401610af8565b6000828152600b6020526040812080546001600160a01b03191690555b600c54811015610d3a5782600c8281548110610c9d57610c9d6153a7565b90600052602060002001541415610d2857600c805460009190610cc2906001906152d0565b81548110610cd257610cd26153a7565b9060005260206000200154905080600c8381548110610cf357610cf36153a7565b600091825260209091200155600c805480610d1057610d10615391565b60019003818190600052602060002001600090559055505b80610d328161530f565b915050610c7f565b50806001600160a01b03167f72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d83604051610d7691815260200190565b60405180910390a2505050565b600280541415610dc35760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b600280556001600160401b0381166000908152600560205260409020546001600160a01b0316610e0657604051630fb532db60e11b815260040160405180910390fd5b6001600160401b03811660009081526006602052604090208054600160601b90046001600160601b0316903490600c610e3f838561527b565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555034600760148282829054906101000a90046001600160601b0316610e87919061527b565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550816001600160401b03167f4e4421a74ab9b76003de9d9527153d002f1338c36c1bbc180069671631cf0958823484610ee39190615238565b604080519283526020830191909152015b60405180910390a250506001600255565b600280541415610f455760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b60028055336000908152600960205260409020546001600160601b0380831691161015610f8557604051631e9acf1760e31b815260040160405180910390fd5b3360009081526009602052604081208054839290610fad9084906001600160601b03166152e7565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600760148282829054906101000a90046001600160601b0316610ff591906152e7565b92506101000a8154816001600160601b0302191690836001600160601b031602179055506000826001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d806000811461106f576040519150601f19603f3d011682016040523d82523d6000602084013e611074565b606091505b505090508061109657604051630dcf35db60e41b815260040160405180910390fd5b5050600160025550565b6002805414156110e05760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b60028055336000908152600860205260409020546001600160601b038083169116101561112057604051631e9acf1760e31b815260040160405180910390fd5b33600090815260086020526040812080548392906111489084906001600160601b03166152e7565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600760088282829054906101000a90046001600160601b031661119091906152e7565b82546101009290920a6001600160601b0381810219909316918316021790915560035460405163a9059cbb60e01b81526001600160a01b03868116600483015292851660248201529116915063a9059cbb90604401602060405180830381600087803b1580156111ff57600080fd5b505af1158015611213573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112379190614ccf565b61125457604051631e9acf1760e31b815260040160405180910390fd5b50506001600255565b611265613072565b6040805180820182526000916112949190849060029083908390808284376000920191909152506127be915050565b6000818152600b60205260409020549091506001600160a01b0316156112e9576040517f4a0b8fa700000000000000000000000000000000000000000000000000000000815260048101829052602401610af8565b6000818152600b6020908152604080832080546001600160a01b0319166001600160a01b038816908117909155600c805460018101825594527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c7909301849055518381527fe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b89101610d76565b6001600160401b03821660009081526005602052604090205482906001600160a01b0316806113b757604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b038216146113eb57604051636c51fda960e11b81526001600160a01b0382166004820152602401610af8565b60028054141561142b5760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b60028080556001600160401b0385166000908152600560205260409020015460641415611484576040517f05a48e0f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831660009081526004602090815260408083206001600160401b03808916855292529091205416156114bd57610be4565b6001600160a01b03831660008181526004602090815260408083206001600160401b038916808552908352818420805467ffffffffffffffff19166001908117909155600584528285206002018054918201815585529383902090930180546001600160a01b031916851790555192835290917f43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e09101610bdb565b6001546001600160a01b031633146115b25760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610af8565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6002805414156116495760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b600280556001600160401b0381166000908152600560205260409020546001600160a01b031661168c57604051630fb532db60e11b815260040160405180910390fd5b6001600160401b0381166000908152600560205260409020600101546001600160a01b03163314611712576001600160401b038116600090815260056020526040908190206001015490517fd084e9750000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610af8565b6001600160401b0381166000818152600560209081526040918290208054336001600160a01b0319808316821784556001909301805490931690925583516001600160a01b03909116808252928101919091529092917f6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f09101610ef4565b60006002805414156117d25760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b6002805560005a905060006117e78585613517565b90506000846060015163ffffffff166001600160401b0381111561180d5761180d6153bd565b604051908082528060200260200182016040528015611836578160200160208202803683370190505b50905060005b856060015163ffffffff168110156118b65782604001518160405160200161186e929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c828281518110611899576118996153a7565b6020908102919091010152806118ae8161530f565b91505061183c565b50602080830180516000908152600d90925260408083208390559051905182917f1fe543e3000000000000000000000000000000000000000000000000000000009161190791908690602401615128565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090506000611981886040015163ffffffff16896080015184613848565b905060006119b087600f600001600a9054906101000a900463ffffffff1663ffffffff163a8c60a00151613896565b90508860a0015115611acf576020808a01516001600160401b03166000908152600690915260409020546001600160601b03808316600160601b909204161015611a0d57604051631e9acf1760e31b815260040160405180910390fd5b6020898101516001600160401b031660009081526006909152604090208054829190600c90611a4d908490600160601b90046001600160601b03166152e7565b82546101009290920a6001600160601b0381810219909316918316021790915587516000908152600b60209081526040808320546001600160a01b031683526009909152812080548594509092611aa69185911661527b565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550611bcd565b6020808a01516001600160401b03166000908152600690915260409020546001600160601b0380831691161015611b1957604051631e9acf1760e31b815260040160405180910390fd5b6020808a01516001600160401b031660009081526006909152604081208054839290611b4f9084906001600160601b03166152e7565b82546101009290920a6001600160601b0381810219909316918316021790915587516000908152600b60209081526040808320546001600160a01b031683526008909152812080548594509092611ba89185911661527b565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505b60208087015160408089015160a08d015182519182526001600160601b03861694820194909452921515908301528315156060830152907f04624b1adcd18b732732ed3009cbdcc7dfc5f0340105149d0c193e8361d955e19060800160405180910390a260016002559998505050505050505050565b611c4b613072565b600a80546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160401b03821660009081526005602052604090205482906001600160a01b031680611caf57604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b03821614611ce357604051636c51fda960e11b81526001600160a01b0382166004820152602401610af8565b600280541415611d235760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b60028055611d3084612ac2565b15611d4e57604051631685ecdd60e31b815260040160405180910390fd5b6001600160a01b03831660009081526004602090815260408083206001600160401b03808916855292529091205416611db457604051637800cff360e11b81526001600160401b03851660048201526001600160a01b0384166024820152604401610af8565b6001600160401b038416600090815260056020908152604080832060020180548251818502810185019093528083529192909190830182828015611e2157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611e03575b50505050509050600060018251611e3891906152d0565b905060005b8251811015611f5f57856001600160a01b0316838281518110611e6257611e626153a7565b60200260200101516001600160a01b03161415611f4d576000838381518110611e8d57611e8d6153a7565b6020026020010151905080600560008a6001600160401b03166001600160401b031681526020019081526020016000206002018381548110611ed157611ed16153a7565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394909416939093179092556001600160401b038a168152600590915260409020600201805480611f2557611f25615391565b600082815260209020810160001990810180546001600160a01b031916905501905550611f5f565b80611f578161530f565b915050611e3d565b506001600160a01b03851660008181526004602090815260408083206001600160401b038b1680855290835292819020805467ffffffffffffffff191690555192835290917f182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b91015b60405180910390a25050600160025550505050565b611fe5613072565b6003546001600160a01b031615612028576040517f2d118a6e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b600060028054141561208c5760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b60028055600780546001600160401b03169060006120a98361532a565b82546101009290920a6001600160401b038181021990931691831602179091556007541690506000806040519080825280602002602001820160405280156120fb578160200160208202803683370190505b50604080518082018252600080825260208083018281526001600160401b038816808452600683528584209451855492516001600160601b03908116600160601b027fffffffffffffffff0000000000000000000000000000000000000000000000009094169116179190911790935583516060810185523381528082018381528186018781529484526005835294909220825181546001600160a01b039182166001600160a01b03199182161783559551600183018054919092169616959095179094559151805194955090936121d99260028501920190614916565b50506040513381526001600160401b03841691507f464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf9060200160405180910390a2509050600160025590565b6001600160401b038116600090815260056020526040812054819081906060906001600160a01b031661226b57604051630fb532db60e11b815260040160405180910390fd5b6001600160401b03851660009081526006602090815260408083205460058352928190208054600290910180548351818602810186019094528084526001600160601b0380871696600160601b900416946001600160a01b0390931693919283919083018282801561230657602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116122e8575b5050505050905093509350935093509193509193565b60028054141561235c5760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b600280556003546001600160a01b031633146123a4576040517f44b0e3c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602081146123de576040517f8129bbcd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006123ec82840184614ef5565b6001600160401b0381166000908152600560205260409020549091506001600160a01b031661242e57604051630fb532db60e11b815260040160405180910390fd5b6001600160401b038116600090815260066020526040812080546001600160601b03169186919061245f838561527b565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555084600760088282829054906101000a90046001600160601b03166124a7919061527b565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550816001600160401b03167fd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f88287846125039190615238565b60408051928352602083019190915201611fc8565b612520613072565b6007544790600160a01b90046001600160601b031681811115612560576040516354ced18160e11b81526004810182905260248101839052604401610af8565b8181101561262f57600061257482846152d0565b90506000846001600160a01b03168260405160006040518083038185875af1925050503d80600081146125c3576040519150601f19603f3d011682016040523d82523d6000602084013e6125c8565b606091505b50509050806125ea57604051630dcf35db60e41b815260040160405180910390fd5b604080516001600160a01b0387168152602081018490527f879c9ea2b9d5345b84ccd12610b032602808517cebdb795007f3dcb4df377317910160405180910390a150505b505050565b61263c613072565b60c861ffff871611156126765760405163539c34bb60e11b815261ffff871660048201819052602482015260c86044820152606401610af8565b6000821361269a576040516321ea67b360e11b815260048101839052602401610af8565b604080516080808201835261ffff891680835263ffffffff89811660208086018290528a83168688018190528a84166060978801819052600f805465ffffffffffff19168717620100008602176dffffffffffffffff0000000000001916660100000000000084026dffffffff000000000000000000001916176a010000000000000000000083021790558951601080548c86015192881667ffffffffffffffff1990911617640100000000928816929092029190911790819055600e8c9055895196875286840194909452978501529483019590955291810186905283821660a08201529290911c1660c08201527f777357bb93f63d088f18112d3dba38457aec633eb8f1341e1d418380ad328e789060e00160405180910390a1505050505050565b6000816040516020016127d19190615052565b604051602081830303815290604052805190602001209050919050565b6001600160401b03821660009081526005602052604090205482906001600160a01b03168061283057604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b0382161461286457604051636c51fda960e11b81526001600160a01b0382166004820152602401610af8565b6002805414156128a45760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b600280556128b184612ac2565b156128cf57604051631685ecdd60e31b815260040160405180910390fd5b610be484846130ce565b600c81815481106128e957600080fd5b600091825260209091200154905081565b612902613072565b6003546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000916001600160a01b0316906370a082319060240160206040518083038186803b15801561295f57600080fd5b505afa158015612973573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129979190614d05565b6007549091506801000000000000000090046001600160601b0316818111156129dd576040516354ced18160e11b81526004810182905260248101839052604401610af8565b8181101561262f5760006129f182846152d0565b60035460405163a9059cbb60e01b81526001600160a01b0387811660048301526024820184905292935091169063a9059cbb90604401602060405180830381600087803b158015612a4157600080fd5b505af1158015612a55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a799190614ccf565b50604080516001600160a01b0386168152602081018390527f59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600910160405180910390a150505050565b6001600160401b0381166000908152600560209081526040808320815160608101835281546001600160a01b0390811682526001830154168185015260028201805484518187028101870186528181528796939586019390929190830182828015612b5657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612b38575b505050505081525050905060005b816040015151811015612cd85760005b600c54811015612cc5576000612c8e600c8381548110612b9657612b966153a7565b906000526020600020015485604001518581518110612bb757612bb76153a7565b6020026020010151886004600089604001518981518110612bda57612bda6153a7565b6020908102919091018101516001600160a01b0316825281810192909252604090810160009081206001600160401b03808f16835293522054166040805160208082018790526001600160a01b0395909516818301526001600160401b039384166060820152919092166080808301919091528251808303909101815260a08201835280519084012060c082019490945260e080820185905282518083039091018152610100909101909152805191012091565b506000818152600d602052604090205490915015612cb25750600195945050505050565b5080612cbd8161530f565b915050612b74565b5080612cd08161530f565b915050612b64565b5060009392505050565b6000600280541415612d245760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b600280556001600160401b0386166000908152600560205260409020546001600160a01b0316612d6757604051630fb532db60e11b815260040160405180910390fd5b3360009081526004602090815260408083206001600160401b03808b1685529252909120541680612dbc57604051637800cff360e11b81526001600160401b0388166004820152336024820152604401610af8565b600f5461ffff9081169087161080612dd8575060c861ffff8716115b15612e0f57600f5460405163539c34bb60e11b815261ffff8089166004830152909116602482015260c86044820152606401610af8565b600f5463ffffffff6201000090910481169086161115612e7657600f546040517ff5d7e01e00000000000000000000000000000000000000000000000000000000815263ffffffff8088166004830152620100009092049091166024820152604401610af8565b6101f463ffffffff85161115612ec8576040517f47386bec00000000000000000000000000000000000000000000000000000000815263ffffffff851660048201526101f46024820152604401610af8565b6000612ed5826001615250565b6040805160208082018d905233828401526001600160401b03808d16606084015284166080808401919091528351808403909101815260a08301845280519082012060c083018e905260e080840182905284518085039091018152610100909301909352815191012091925081612f4a6138e6565b6040805160208101939093528201526001600160401b038b16606082015263ffffffff808a166080830152881660a08201523360c082015260e00160408051601f1981840301815282825280516020918201206000868152600d835283902055848352820183905261ffff8b169082015263ffffffff808a1660608301528816608082015286151560a082015233906001600160401b038c16908d907fad8d0b36e0dd1bc233432f9e3c28d376b9713fc9d26ac2dc972077be458d86fe9060c00160405180910390a4503360009081526004602090815260408083206001600160401b03808e168552925290912080549190931667ffffffffffffffff199091161790915591505060016002559695505050505050565b613069613072565b610a838161397f565b6000546001600160a01b031633146130cc5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610af8565b565b60028054141561310e5760405162461bcd60e51b815260206004820152601f60248201526000805160206153f78339815191526044820152606401610af8565b60028080556001600160401b0383166000908152600560209081526040808320815160608101835281546001600160a01b03908116825260018301541681850152948101805483518186028101860185528181529596959294938601938301828280156131a457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613186575b505050919092525050506001600160401b03841660009081526006602090815260408083208151808301909252546001600160601b03808216808452600160601b9092041692820183905293945092915b84604001515181101561327657600460008660400151838151811061321c5761321c6153a7565b6020908102919091018101516001600160a01b0316825281810192909252604090810160009081206001600160401b038b1682529092529020805467ffffffffffffffff191690558061326e8161530f565b9150506131f5565b506001600160401b038616600090815260056020526040812080546001600160a01b031990811682556001820180549091169055906132b8600283018261497b565b50506001600160401b038616600090815260066020526040902080547fffffffffffffffff000000000000000000000000000000000000000000000000169055600780548391906008906133229084906801000000000000000090046001600160601b03166152e7565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600760148282829054906101000a90046001600160601b031661336a91906152e7565b82546101009290920a6001600160601b0381810219909316918316021790915560035460405163a9059cbb60e01b81526001600160a01b03898116600483015292861660248201529116915063a9059cbb90604401602060405180830381600087803b1580156133d957600080fd5b505af11580156133ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134119190614ccf565b61342e57604051631e9acf1760e31b815260040160405180910390fd5b6000856001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d8060008114613484576040519150601f19603f3d011682016040523d82523d6000602084013e613489565b606091505b50509050806134ab57604051630dcf35db60e41b815260040160405180910390fd5b604080516001600160a01b03881681526001600160601b03858116602083015284168183015290516001600160401b038916917fb3a76e2c52bb7a484e325118e4cadc4346fdaa16a7df0621c5525e1f7ab4d495919081900360600190a2505060016002555050505050565b6040805160608101825260008082526020820181905291810191909152600061354384600001516127be565b6000818152600b60205260409020549091506001600160a01b03168061357f57604051631dfd6e1360e21b815260048101839052602401610af8565b60008286608001516040516020016135a1929190918252602082015260400190565b60408051601f1981840301815291815281516020928301206000818152600d90935291205490915080613600576040517f3688124a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b85516020808801516040808a015160608b015160808c0151925161366b96899690959491019586526001600160401b03948516602087015292909316604085015263ffffffff90811660608501529190911660808301526001600160a01b031660a082015260c00190565b6040516020818303038152906040528051906020012081146136b9576040517fd529142c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006136c88760000151613a29565b9050806137d25786516040517fe9413d380000000000000000000000000000000000000000000000000000000081526001600160401b0390911660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063e9413d389060240160206040518083038186803b15801561375357600080fd5b505afa158015613767573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061378b9190614d05565b9050806137d25786516040517f175dadad0000000000000000000000000000000000000000000000000000000081526001600160401b039091166004820152602401610af8565b60008860800151826040516020016137f4929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c9050600061381b8a83613b2c565b90506040518060600160405280888152602001868152602001828152509750505050505050505b92915050565b60005a61138881101561385a57600080fd5b61138881039050846040820482031161387257600080fd5b50823b61387e57600080fd5b60008083516020850160008789f190505b9392505050565b600081156138c4576010546138bd9086908690640100000000900463ffffffff1686613b97565b90506138de565b6010546138db908690869063ffffffff1686613c01565b90505b949350505050565b60004661a4b18114806138fb575062066eed81145b156139785760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561393a57600080fd5b505afa15801561394e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139729190614d05565b91505090565b4391505090565b6001600160a01b0381163314156139d85760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610af8565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60004661a4b1811480613a3e575062066eed81145b15613b1d57610100836001600160401b0316613a586138e6565b613a6291906152d0565b1180613a7e5750613a716138e6565b836001600160401b031610155b15613a8c5750600092915050565b6040517f2b407a820000000000000000000000000000000000000000000000000000000081526001600160401b0384166004820152606490632b407a829060240160206040518083038186803b158015613ae557600080fd5b505afa158015613af9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061388f9190614d05565b50506001600160401b03164090565b6000613b608360000151846020015185604001518660600151868860a001518960c001518a60e001518b6101000151613d08565b60038360200151604051602001613b78929190615114565b60408051601f1981840301815291905280516020909101209392505050565b600080613ba2613f43565b905060005a613bb18888615238565b613bbb91906152d0565b613bc590856152b1565b90506000613bde63ffffffff871664e8d4a510006152b1565b905082613beb8284615238565b613bf59190615238565b98975050505050505050565b600080613c0c613f9f565b905060008113613c32576040516321ea67b360e11b815260048101829052602401610af8565b6000613c3c613f43565b9050600082825a613c4d8b8b615238565b613c5791906152d0565b613c6190886152b1565b613c6b9190615238565b613c7d90670de0b6b3a76400006152b1565b613c87919061529d565b90506000613ca063ffffffff881664e8d4a510006152b1565b9050613cb8816b033b2e3c9fd0803ce80000006152d0565b821115613cf1576040517fe80fa38100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613cfb8183615238565b9998505050505050505050565b613d1189614086565b613d5d5760405162461bcd60e51b815260206004820152601a60248201527f7075626c6963206b6579206973206e6f74206f6e2063757276650000000000006044820152606401610af8565b613d6688614086565b613db25760405162461bcd60e51b815260206004820152601560248201527f67616d6d61206973206e6f74206f6e20637572766500000000000000000000006044820152606401610af8565b613dbb83614086565b613e075760405162461bcd60e51b815260206004820152601d60248201527f6347616d6d615769746e657373206973206e6f74206f6e2063757276650000006044820152606401610af8565b613e1082614086565b613e5c5760405162461bcd60e51b815260206004820152601c60248201527f73486173685769746e657373206973206e6f74206f6e206375727665000000006044820152606401610af8565b613e68878a888761415f565b613eb45760405162461bcd60e51b815260206004820152601960248201527f6164647228632a706b2b732a6729213d5f755769746e657373000000000000006044820152606401610af8565b6000613ec08a87614294565b90506000613ed3898b878b8689896142f8565b90506000613ee4838d8d8a86614418565b9050808a14613f355760405162461bcd60e51b815260206004820152600d60248201527f696e76616c69642070726f6f66000000000000000000000000000000000000006044820152606401610af8565b505050505050505050505050565b60004661a4b1811480613f58575062066eed81145b15613f9757606c6001600160a01b031663c6f7de0e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561393a57600080fd5b600091505090565b600f54600a54604080517ffeaf968c00000000000000000000000000000000000000000000000000000000815290516000936601000000000000900463ffffffff169283151592859283926001600160a01b03169163feaf968c9160048083019260a0929190829003018186803b15801561401957600080fd5b505afa15801561402d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906140519190614f47565b509450909250849150508015614075575061406c82426152d0565b8463ffffffff16105b156138de5750600e54949350505050565b80516000906401000003d019116140df5760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420782d6f7264696e61746500000000000000000000000000006044820152606401610af8565b60208201516401000003d019116141385760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420792d6f7264696e61746500000000000000000000000000006044820152606401610af8565b60208201516401000003d0199080096141588360005b6020020151614458565b1492915050565b60006001600160a01b0382166141b75760405162461bcd60e51b815260206004820152600b60248201527f626164207769746e6573730000000000000000000000000000000000000000006044820152606401610af8565b6020840151600090600116156141ce57601c6141d1565b601b5b9050600070014551231950b75fc4402da1732fc9bebe1985876000602002015109865170014551231950b75fc4402da1732fc9bebe19918203925060009190890987516040805160008082526020820180845287905260ff88169282019290925260608101929092526080820183905291925060019060a0016020604051602081039080840390855afa15801561426c573d6000803e3d6000fd5b5050604051601f1901516001600160a01b039081169088161495505050505050949350505050565b61429c614999565b6142c9600184846040516020016142b593929190615031565b60405160208183030381529060405261447c565b90505b6142d581614086565b6138425780516040805160208101929092526142f191016142b5565b90506142cc565b614300614999565b825186516401000003d019908190069106141561435f5760405162461bcd60e51b815260206004820152601e60248201527f706f696e747320696e2073756d206d7573742062652064697374696e637400006044820152606401610af8565b61436a8789886144cb565b6143b65760405162461bcd60e51b815260206004820152601660248201527f4669727374206d756c20636865636b206661696c6564000000000000000000006044820152606401610af8565b6143c18486856144cb565b61440d5760405162461bcd60e51b815260206004820152601760248201527f5365636f6e64206d756c20636865636b206661696c65640000000000000000006044820152606401610af8565b613bf5868484614605565b60006002868686858760405160200161443696959493929190614fd2565b60408051601f1981840301815291905280516020909101209695505050505050565b6000806401000003d01980848509840990506401000003d019600782089392505050565b614484614999565b61448d826146cc565b81526144a261449d82600061414e565b614707565b6020820181905260029006600114156144c6576020810180516401000003d0190390525b919050565b60008261451a5760405162461bcd60e51b815260206004820152600b60248201527f7a65726f207363616c61720000000000000000000000000000000000000000006044820152606401610af8565b8351602085015160009061453090600290615351565b1561453c57601c61453f565b601b5b9050600070014551231950b75fc4402da1732fc9bebe198387096040805160008082526020820180845281905260ff86169282019290925260608101869052608081018390529192509060019060a0016020604051602081039080840390855afa1580156145b1573d6000803e3d6000fd5b5050506020604051035190506000866040516020016145d09190614fc0565b60408051601f1981840301815291905280516020909101206001600160a01b0392831692169190911498975050505050505050565b61460d614999565b83516020808601518551918601516000938493849361462e93909190614727565b919450925090506401000003d01985820960011461468e5760405162461bcd60e51b815260206004820152601960248201527f696e765a206d75737420626520696e7665727365206f66207a000000000000006044820152606401610af8565b60405180604001604052806401000003d019806146ad576146ad61537b565b87860981526020016401000003d0198785099052979650505050505050565b805160208201205b6401000003d01981106144c6576040805160208082019390935281518082038401815290820190915280519101206146d4565b60006138428260026147206401000003d0196001615238565b901c614807565b60008080600180826401000003d019896401000003d019038808905060006401000003d0198b6401000003d019038a0890506000614767838385856148a9565b909850905061477888828e886148cd565b909850905061478988828c876148cd565b9098509050600061479c8d878b856148cd565b90985090506147ad888286866148a9565b90985090506147be88828e896148cd565b90985090508181146147f3576401000003d019818a0998506401000003d01982890997506401000003d01981830996506147f7565b8196505b5050505050509450945094915050565b6000806148126149b7565b6020808252818101819052604082015260608101859052608081018490526401000003d01960a08201526148446149d5565b60208160c0846005600019fa92508261489f5760405162461bcd60e51b815260206004820152601260248201527f6269674d6f64457870206661696c7572652100000000000000000000000000006044820152606401610af8565b5195945050505050565b6000806401000003d0198487096401000003d0198487099097909650945050505050565b600080806401000003d019878509905060006401000003d01987876401000003d019030990506401000003d0198183086401000003d01986890990999098509650505050505050565b82805482825590600052602060002090810192821561496b579160200282015b8281111561496b57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614936565b506149779291506149f3565b5090565b5080546000825590600052602060002090810190610a8391906149f3565b60405180604001604052806002906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b5b8082111561497757600081556001016149f4565b80356144c6816153d3565b806040810183101561384257600080fd5b600082601f830112614a3557600080fd5b614a3d6151ed565b808385604086011115614a4f57600080fd5b60005b6002811015614a71578135845260209384019390910190600101614a52565b509095945050505050565b600060c08284031215614a8e57600080fd5b60405160c081018181106001600160401b0382111715614ab057614ab06153bd565b604052905080614abf83614b48565b8152614acd60208401614b48565b6020820152614ade60408401614b34565b6040820152614aef60608401614b34565b60608201526080830135614b02816153d3565b608082015260a0830135614b15816153e8565b60a0919091015292915050565b803561ffff811681146144c657600080fd5b803563ffffffff811681146144c657600080fd5b80356001600160401b03811681146144c657600080fd5b805169ffffffffffffffffffff811681146144c657600080fd5b80356001600160601b03811681146144c657600080fd5b600060208284031215614ba257600080fd5b813561388f816153d3565b60008060408385031215614bc057600080fd5b8235614bcb816153d3565b9150614bd960208401614b79565b90509250929050565b60008060608385031215614bf557600080fd5b8235614c00816153d3565b9150614bd98460208501614a13565b60008060008060608587031215614c2557600080fd5b8435614c30816153d3565b93506020850135925060408501356001600160401b0380821115614c5357600080fd5b818701915087601f830112614c6757600080fd5b813581811115614c7657600080fd5b886020828501011115614c8857600080fd5b95989497505060200194505050565b600060408284031215614ca957600080fd5b61388f8383614a13565b600060408284031215614cc557600080fd5b61388f8383614a24565b600060208284031215614ce157600080fd5b815161388f816153e8565b600060208284031215614cfe57600080fd5b5035919050565b600060208284031215614d1757600080fd5b5051919050565b60008060008060008060c08789031215614d3757600080fd5b86359550614d4760208801614b48565b9450614d5560408801614b22565b9350614d6360608801614b34565b9250614d7160808801614b34565b915060a0870135614d81816153e8565b809150509295509295509295565b600080828403610260811215614da457600080fd5b6101a080821215614db457600080fd5b614dbc615215565b9150614dc88686614a24565b8252614dd78660408701614a24565b60208301526080850135604083015260a0850135606083015260c08501356080830152614e0660e08601614a08565b60a0830152610100614e1a87828801614a24565b60c0840152614e2d876101408801614a24565b60e08401526101808601358184015250819350614e4c86828701614a7c565b925050509250929050565b60008060008060008086880360e0811215614e7157600080fd5b614e7a88614b22565b9650614e8860208901614b34565b9550614e9660408901614b34565b9450614ea460608901614b34565b9350608088013592506040609f1982011215614ebf57600080fd5b50614ec86151ed565b614ed460a08901614b34565b8152614ee260c08901614b34565b6020820152809150509295509295509295565b600060208284031215614f0757600080fd5b61388f82614b48565b60008060408385031215614f2357600080fd5b614f2c83614b48565b91506020830135614f3c816153d3565b809150509250929050565b600080600080600060a08688031215614f5f57600080fd5b614f6886614b5f565b9450602086015193506040860151925060608601519150614f8b60808701614b5f565b90509295509295909350565b8060005b6002811015614fba578151845260209384019390910190600101614f9b565b50505050565b614fca8183614f97565b604001919050565b868152614fe26020820187614f97565b614fef6060820186614f97565b614ffc60a0820185614f97565b61500960e0820184614f97565b60609190911b6bffffffffffffffffffffffff19166101208201526101340195945050505050565b8381526150416020820184614f97565b606081019190915260800192915050565b604081016138428284614f97565b600060208083528351808285015260005b8181101561508d57858101830151858201604001528201615071565b8181111561509f576000604083870101525b50601f01601f1916929092016040019392505050565b60006060820161ffff86168352602063ffffffff86168185015260606040850152818551808452608086019150828701935060005b81811015615106578451835293830193918301916001016150ea565b509098975050505050505050565b8281526060810161388f6020830184614f97565b6000604082018483526020604081850152818551808452606086019150828701935060005b818110156151695784518352938301939183019160010161514d565b5090979650505050505050565b6000608082016001600160601b0380881684526020818816818601526001600160a01b03915081871660408601526080606086015282865180855260a087019150828801945060005b818110156151dd5785518516835294830194918301916001016151bf565b50909a9950505050505050505050565b604080519081016001600160401b038111828210171561520f5761520f6153bd565b60405290565b60405161012081016001600160401b038111828210171561520f5761520f6153bd565b6000821982111561524b5761524b615365565b500190565b60006001600160401b0380831681851680830382111561527257615272615365565b01949350505050565b60006001600160601b0380831681851680830382111561527257615272615365565b6000826152ac576152ac61537b565b500490565b60008160001904831182151516156152cb576152cb615365565b500290565b6000828210156152e2576152e2615365565b500390565b60006001600160601b038381169083168181101561530757615307615365565b039392505050565b600060001982141561532357615323615365565b5060010190565b60006001600160401b038083168181141561534757615347615365565b6001019392505050565b6000826153605761536061537b565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610a8357600080fd5b8015158114610a8357600080fdfe5265656e7472616e637947756172643a207265656e7472616e742063616c6c00a164736f6c6343000806000a",
}

var VRFCoordinatorV25ABI = VRFCoordinatorV25MetaData.ABI

var VRFCoordinatorV25Bin = VRFCoordinatorV25MetaData.Bin

func DeployVRFCoordinatorV25(auth *bind.TransactOpts, backend bind.ContractBackend, blockhashStore common.Address) (common.Address, *types.Transaction, *VRFCoordinatorV25, error) {
	parsed, err := VRFCoordinatorV25MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFCoordinatorV25Bin), backend, blockhashStore)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFCoordinatorV25{VRFCoordinatorV25Caller: VRFCoordinatorV25Caller{contract: contract}, VRFCoordinatorV25Transactor: VRFCoordinatorV25Transactor{contract: contract}, VRFCoordinatorV25Filterer: VRFCoordinatorV25Filterer{contract: contract}}, nil
}

type VRFCoordinatorV25 struct {
	address common.Address
	abi     abi.ABI
	VRFCoordinatorV25Caller
	VRFCoordinatorV25Transactor
	VRFCoordinatorV25Filterer
}

type VRFCoordinatorV25Caller struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV25Transactor struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV25Filterer struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV25Session struct {
	Contract     *VRFCoordinatorV25
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV25CallerSession struct {
	Contract *VRFCoordinatorV25Caller
	CallOpts bind.CallOpts
}

type VRFCoordinatorV25TransactorSession struct {
	Contract     *VRFCoordinatorV25Transactor
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV25Raw struct {
	Contract *VRFCoordinatorV25
}

type VRFCoordinatorV25CallerRaw struct {
	Contract *VRFCoordinatorV25Caller
}

type VRFCoordinatorV25TransactorRaw struct {
	Contract *VRFCoordinatorV25Transactor
}

func NewVRFCoordinatorV25(address common.Address, backend bind.ContractBackend) (*VRFCoordinatorV25, error) {
	abi, err := abi.JSON(strings.NewReader(VRFCoordinatorV25ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFCoordinatorV25(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25{address: address, abi: abi, VRFCoordinatorV25Caller: VRFCoordinatorV25Caller{contract: contract}, VRFCoordinatorV25Transactor: VRFCoordinatorV25Transactor{contract: contract}, VRFCoordinatorV25Filterer: VRFCoordinatorV25Filterer{contract: contract}}, nil
}

func NewVRFCoordinatorV25Caller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorV25Caller, error) {
	contract, err := bindVRFCoordinatorV25(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25Caller{contract: contract}, nil
}

func NewVRFCoordinatorV25Transactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorV25Transactor, error) {
	contract, err := bindVRFCoordinatorV25(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25Transactor{contract: contract}, nil
}

func NewVRFCoordinatorV25Filterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorV25Filterer, error) {
	contract, err := bindVRFCoordinatorV25(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25Filterer{contract: contract}, nil
}

func bindVRFCoordinatorV25(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFCoordinatorV25MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV25.Contract.VRFCoordinatorV25Caller.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.VRFCoordinatorV25Transactor.contract.Transfer(opts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.VRFCoordinatorV25Transactor.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV25.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.contract.Transfer(opts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "BLOCKHASH_STORE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) LINK() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.LINK(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) LINK() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.LINK(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) LINKETHFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "LINK_ETH_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) LINKETHFEED() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.LINKETHFEED(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) LINKETHFEED() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.LINKETHFEED(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV25.Contract.MAXCONSUMERS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV25.Contract.MAXCONSUMERS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) MAXNUMWORDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "MAX_NUM_WORDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV25.Contract.MAXNUMWORDS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV25.Contract.MAXNUMWORDS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "MAX_REQUEST_CONFIRMATIONS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV25.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV25.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "getRequestConfig")

	if err != nil {
		return *new(uint16), *new(uint32), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV25.Contract.GetRequestConfig(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV25.Contract.GetRequestConfig(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) GetSubscription(opts *bind.CallOpts, subId uint64) (GetSubscription,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(GetSubscription)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EthBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) GetSubscription(subId uint64) (GetSubscription,

	error) {
	return _VRFCoordinatorV25.Contract.GetSubscription(&_VRFCoordinatorV25.CallOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) GetSubscription(subId uint64) (GetSubscription,

	error) {
	return _VRFCoordinatorV25.Contract.GetSubscription(&_VRFCoordinatorV25.CallOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "hashOfKey", publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.HashOfKey(&_VRFCoordinatorV25.CallOpts, publicKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.HashOfKey(&_VRFCoordinatorV25.CallOpts, publicKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) Owner() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.Owner(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) Owner() (common.Address, error) {
	return _VRFCoordinatorV25.Contract.Owner(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) PendingRequestExists(opts *bind.CallOpts, subId uint64) (bool, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "pendingRequestExists", subId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) PendingRequestExists(subId uint64) (bool, error) {
	return _VRFCoordinatorV25.Contract.PendingRequestExists(&_VRFCoordinatorV25.CallOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) PendingRequestExists(subId uint64) (bool, error) {
	return _VRFCoordinatorV25.Contract.PendingRequestExists(&_VRFCoordinatorV25.CallOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SConfig(opts *bind.CallOpts) (SConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_config")

	outstruct := new(SConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumRequestConfirmations = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MaxGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.StalenessSeconds = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasAfterPaymentCalculation = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SConfig() (SConfig,

	error) {
	return _VRFCoordinatorV25.Contract.SConfig(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SConfig() (SConfig,

	error) {
	return _VRFCoordinatorV25.Contract.SConfig(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SCurrentSubId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_currentSubId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SCurrentSubId() (uint64, error) {
	return _VRFCoordinatorV25.Contract.SCurrentSubId(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SCurrentSubId() (uint64, error) {
	return _VRFCoordinatorV25.Contract.SCurrentSubId(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_fallbackWeiPerUnitLink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.SFallbackWeiPerUnitLink(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.SFallbackWeiPerUnitLink(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SFeeConfig(opts *bind.CallOpts) (SFeeConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_feeConfig")

	outstruct := new(SFeeConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FulfillmentFlatFeeLinkPPM = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeEthPPM = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SFeeConfig() (SFeeConfig,

	error) {
	return _VRFCoordinatorV25.Contract.SFeeConfig(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SFeeConfig() (SFeeConfig,

	error) {
	return _VRFCoordinatorV25.Contract.SFeeConfig(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_provingKeyHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.SProvingKeyHashes(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.SProvingKeyHashes(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_provingKeys", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SProvingKeys(arg0 [32]byte) (common.Address, error) {
	return _VRFCoordinatorV25.Contract.SProvingKeys(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SProvingKeys(arg0 [32]byte) (common.Address, error) {
	return _VRFCoordinatorV25.Contract.SProvingKeys(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_requestCommitments", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.SRequestCommitments(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV25.Contract.SRequestCommitments(&_VRFCoordinatorV25.CallOpts, arg0)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) STotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_totalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) STotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.STotalBalance(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) STotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.STotalBalance(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) STotalEthBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "s_totalEthBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) STotalEthBalance() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.STotalEthBalance(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) STotalEthBalance() (*big.Int, error) {
	return _VRFCoordinatorV25.Contract.STotalEthBalance(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VRFCoordinatorV25.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) TypeAndVersion() (string, error) {
	return _VRFCoordinatorV25.Contract.TypeAndVersion(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25CallerSession) TypeAndVersion() (string, error) {
	return _VRFCoordinatorV25.Contract.TypeAndVersion(&_VRFCoordinatorV25.CallOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "acceptOwnership")
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AcceptOwnership(&_VRFCoordinatorV25.TransactOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AcceptOwnership(&_VRFCoordinatorV25.TransactOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) AddConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "addConsumer", subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AddConsumer(&_VRFCoordinatorV25.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.AddConsumer(&_VRFCoordinatorV25.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) CancelSubscription(opts *bind.TransactOpts, subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "cancelSubscription", subId, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.CancelSubscription(&_VRFCoordinatorV25.TransactOpts, subId, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.CancelSubscription(&_VRFCoordinatorV25.TransactOpts, subId, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "createSubscription")
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.CreateSubscription(&_VRFCoordinatorV25.TransactOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.CreateSubscription(&_VRFCoordinatorV25.TransactOpts)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "deregisterProvingKey", publicProvingKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.DeregisterProvingKey(&_VRFCoordinatorV25.TransactOpts, publicProvingKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.DeregisterProvingKey(&_VRFCoordinatorV25.TransactOpts, publicProvingKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV25RequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "fulfillRandomWords", proof, rc)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV25RequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.FulfillRandomWords(&_VRFCoordinatorV25.TransactOpts, proof, rc)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV25RequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.FulfillRandomWords(&_VRFCoordinatorV25.TransactOpts, proof, rc)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) FundSubscriptionWithEth(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "fundSubscriptionWithEth", subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) FundSubscriptionWithEth(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.FundSubscriptionWithEth(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) FundSubscriptionWithEth(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.FundSubscriptionWithEth(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OnTokenTransfer(&_VRFCoordinatorV25.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OnTokenTransfer(&_VRFCoordinatorV25.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "oracleWithdraw", recipient, amount)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OracleWithdraw(&_VRFCoordinatorV25.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OracleWithdraw(&_VRFCoordinatorV25.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) OracleWithdrawEth(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "oracleWithdrawEth", recipient, amount)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) OracleWithdrawEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OracleWithdrawEth(&_VRFCoordinatorV25.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) OracleWithdrawEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OracleWithdrawEth(&_VRFCoordinatorV25.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) OwnerCancelSubscription(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "ownerCancelSubscription", subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) OwnerCancelSubscription(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OwnerCancelSubscription(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) OwnerCancelSubscription(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.OwnerCancelSubscription(&_VRFCoordinatorV25.TransactOpts, subId)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RecoverEthFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "recoverEthFunds", to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RecoverEthFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RecoverEthFunds(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RecoverEthFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RecoverEthFunds(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "recoverFunds", to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RecoverFunds(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RecoverFunds(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "registerProvingKey", oracle, publicProvingKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RegisterProvingKey(&_VRFCoordinatorV25.TransactOpts, oracle, publicProvingKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RegisterProvingKey(&_VRFCoordinatorV25.TransactOpts, oracle, publicProvingKey)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RemoveConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "removeConsumer", subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RemoveConsumer(&_VRFCoordinatorV25.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RemoveConsumer(&_VRFCoordinatorV25.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RequestRandomWords(opts *bind.TransactOpts, keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32, payInEth bool) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "requestRandomWords", keyHash, subId, requestConfirmations, callbackGasLimit, numWords, payInEth)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RequestRandomWords(keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32, payInEth bool) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RequestRandomWords(&_VRFCoordinatorV25.TransactOpts, keyHash, subId, requestConfirmations, callbackGasLimit, numWords, payInEth)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RequestRandomWords(keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32, payInEth bool) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RequestRandomWords(&_VRFCoordinatorV25.TransactOpts, keyHash, subId, requestConfirmations, callbackGasLimit, numWords, payInEth)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV25.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV25.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV25FeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "setConfig", minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV25FeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.SetConfig(&_VRFCoordinatorV25.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV25FeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.SetConfig(&_VRFCoordinatorV25.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) SetLINK(opts *bind.TransactOpts, link common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "setLINK", link)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SetLINK(link common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.SetLINK(&_VRFCoordinatorV25.TransactOpts, link)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) SetLINK(link common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.SetLINK(&_VRFCoordinatorV25.TransactOpts, link)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) SetLinkEthFeed(opts *bind.TransactOpts, linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "setLinkEthFeed", linkEthFeed)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) SetLinkEthFeed(linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.SetLinkEthFeed(&_VRFCoordinatorV25.TransactOpts, linkEthFeed)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) SetLinkEthFeed(linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.SetLinkEthFeed(&_VRFCoordinatorV25.TransactOpts, linkEthFeed)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.TransferOwnership(&_VRFCoordinatorV25.TransactOpts, to)
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV25.Contract.TransferOwnership(&_VRFCoordinatorV25.TransactOpts, to)
}

type VRFCoordinatorV25ConfigSetIterator struct {
	Event *VRFCoordinatorV25ConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25ConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25ConfigSet)
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
		it.Event = new(VRFCoordinatorV25ConfigSet)
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

func (it *VRFCoordinatorV25ConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25ConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25ConfigSet struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
	FallbackWeiPerUnitLink      *big.Int
	FeeConfig                   VRFCoordinatorV25FeeConfig
	Raw                         types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV25ConfigSetIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25ConfigSetIterator{contract: _VRFCoordinatorV25.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25ConfigSet)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseConfigSet(log types.Log) (*VRFCoordinatorV25ConfigSet, error) {
	event := new(VRFCoordinatorV25ConfigSet)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25EthFundsRecoveredIterator struct {
	Event *VRFCoordinatorV25EthFundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25EthFundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25EthFundsRecovered)
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
		it.Event = new(VRFCoordinatorV25EthFundsRecovered)
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

func (it *VRFCoordinatorV25EthFundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25EthFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25EthFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterEthFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV25EthFundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "EthFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25EthFundsRecoveredIterator{contract: _VRFCoordinatorV25.contract, event: "EthFundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchEthFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25EthFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "EthFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25EthFundsRecovered)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "EthFundsRecovered", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseEthFundsRecovered(log types.Log) (*VRFCoordinatorV25EthFundsRecovered, error) {
	event := new(VRFCoordinatorV25EthFundsRecovered)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "EthFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25FundsRecoveredIterator struct {
	Event *VRFCoordinatorV25FundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25FundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25FundsRecovered)
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
		it.Event = new(VRFCoordinatorV25FundsRecovered)
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

func (it *VRFCoordinatorV25FundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25FundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25FundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV25FundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25FundsRecoveredIterator{contract: _VRFCoordinatorV25.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25FundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25FundsRecovered)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseFundsRecovered(log types.Log) (*VRFCoordinatorV25FundsRecovered, error) {
	event := new(VRFCoordinatorV25FundsRecovered)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25OwnershipTransferRequestedIterator struct {
	Event *VRFCoordinatorV25OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25OwnershipTransferRequested)
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
		it.Event = new(VRFCoordinatorV25OwnershipTransferRequested)
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

func (it *VRFCoordinatorV25OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV25OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25OwnershipTransferRequestedIterator{contract: _VRFCoordinatorV25.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25OwnershipTransferRequested)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV25OwnershipTransferRequested, error) {
	event := new(VRFCoordinatorV25OwnershipTransferRequested)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25OwnershipTransferredIterator struct {
	Event *VRFCoordinatorV25OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25OwnershipTransferred)
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
		it.Event = new(VRFCoordinatorV25OwnershipTransferred)
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

func (it *VRFCoordinatorV25OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV25OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25OwnershipTransferredIterator{contract: _VRFCoordinatorV25.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25OwnershipTransferred)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV25OwnershipTransferred, error) {
	event := new(VRFCoordinatorV25OwnershipTransferred)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25ProvingKeyDeregisteredIterator struct {
	Event *VRFCoordinatorV25ProvingKeyDeregistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25ProvingKeyDeregisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25ProvingKeyDeregistered)
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
		it.Event = new(VRFCoordinatorV25ProvingKeyDeregistered)
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

func (it *VRFCoordinatorV25ProvingKeyDeregisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25ProvingKeyDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25ProvingKeyDeregistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterProvingKeyDeregistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV25ProvingKeyDeregisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "ProvingKeyDeregistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25ProvingKeyDeregisteredIterator{contract: _VRFCoordinatorV25.contract, event: "ProvingKeyDeregistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ProvingKeyDeregistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "ProvingKeyDeregistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25ProvingKeyDeregistered)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseProvingKeyDeregistered(log types.Log) (*VRFCoordinatorV25ProvingKeyDeregistered, error) {
	event := new(VRFCoordinatorV25ProvingKeyDeregistered)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25ProvingKeyRegisteredIterator struct {
	Event *VRFCoordinatorV25ProvingKeyRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25ProvingKeyRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25ProvingKeyRegistered)
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
		it.Event = new(VRFCoordinatorV25ProvingKeyRegistered)
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

func (it *VRFCoordinatorV25ProvingKeyRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25ProvingKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25ProvingKeyRegistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV25ProvingKeyRegisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25ProvingKeyRegisteredIterator{contract: _VRFCoordinatorV25.contract, event: "ProvingKeyRegistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ProvingKeyRegistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25ProvingKeyRegistered)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV25ProvingKeyRegistered, error) {
	event := new(VRFCoordinatorV25ProvingKeyRegistered)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25RandomWordsFulfilledIterator struct {
	Event *VRFCoordinatorV25RandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25RandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25RandomWordsFulfilled)
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
		it.Event = new(VRFCoordinatorV25RandomWordsFulfilled)
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

func (it *VRFCoordinatorV25RandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25RandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25RandomWordsFulfilled struct {
	RequestId     *big.Int
	OutputSeed    *big.Int
	Payment       *big.Int
	NativePayment bool
	Success       bool
	Raw           types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*VRFCoordinatorV25RandomWordsFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25RandomWordsFulfilledIterator{contract: _VRFCoordinatorV25.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25RandomWordsFulfilled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25RandomWordsFulfilled)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV25RandomWordsFulfilled, error) {
	event := new(VRFCoordinatorV25RandomWordsFulfilled)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25RandomWordsRequestedIterator struct {
	Event *VRFCoordinatorV25RandomWordsRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25RandomWordsRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25RandomWordsRequested)
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
		it.Event = new(VRFCoordinatorV25RandomWordsRequested)
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

func (it *VRFCoordinatorV25RandomWordsRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25RandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25RandomWordsRequested struct {
	KeyHash                     [32]byte
	RequestId                   *big.Int
	PreSeed                     *big.Int
	SubId                       uint64
	MinimumRequestConfirmations uint16
	CallbackGasLimit            uint32
	NumWords                    uint32
	NativePayment               bool
	Sender                      common.Address
	Raw                         types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []uint64, sender []common.Address) (*VRFCoordinatorV25RandomWordsRequestedIterator, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25RandomWordsRequestedIterator{contract: _VRFCoordinatorV25.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25RandomWordsRequested, keyHash [][32]byte, subId []uint64, sender []common.Address) (event.Subscription, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25RandomWordsRequested)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV25RandomWordsRequested, error) {
	event := new(VRFCoordinatorV25RandomWordsRequested)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionCanceledIterator struct {
	Event *VRFCoordinatorV25SubscriptionCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionCanceled)
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
		it.Event = new(VRFCoordinatorV25SubscriptionCanceled)
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

func (it *VRFCoordinatorV25SubscriptionCanceledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionCanceled struct {
	SubId      uint64
	To         common.Address
	AmountLink *big.Int
	AmountEth  *big.Int
	Raw        types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionCanceledIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionCanceled, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionCanceled)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV25SubscriptionCanceled, error) {
	event := new(VRFCoordinatorV25SubscriptionCanceled)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionConsumerAddedIterator struct {
	Event *VRFCoordinatorV25SubscriptionConsumerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionConsumerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionConsumerAdded)
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
		it.Event = new(VRFCoordinatorV25SubscriptionConsumerAdded)
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

func (it *VRFCoordinatorV25SubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionConsumerAdded struct {
	SubId    uint64
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionConsumerAddedIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionConsumerAdded, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionConsumerAdded)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV25SubscriptionConsumerAdded, error) {
	event := new(VRFCoordinatorV25SubscriptionConsumerAdded)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionConsumerRemovedIterator struct {
	Event *VRFCoordinatorV25SubscriptionConsumerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionConsumerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionConsumerRemoved)
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
		it.Event = new(VRFCoordinatorV25SubscriptionConsumerRemoved)
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

func (it *VRFCoordinatorV25SubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionConsumerRemoved struct {
	SubId    uint64
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionConsumerRemovedIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionConsumerRemoved, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionConsumerRemoved)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV25SubscriptionConsumerRemoved, error) {
	event := new(VRFCoordinatorV25SubscriptionConsumerRemoved)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionCreatedIterator struct {
	Event *VRFCoordinatorV25SubscriptionCreated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionCreatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionCreated)
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
		it.Event = new(VRFCoordinatorV25SubscriptionCreated)
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

func (it *VRFCoordinatorV25SubscriptionCreatedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionCreated struct {
	SubId uint64
	Owner common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionCreatedIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionCreated, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionCreated)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV25SubscriptionCreated, error) {
	event := new(VRFCoordinatorV25SubscriptionCreated)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionFundedIterator struct {
	Event *VRFCoordinatorV25SubscriptionFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionFunded)
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
		it.Event = new(VRFCoordinatorV25SubscriptionFunded)
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

func (it *VRFCoordinatorV25SubscriptionFundedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionFunded struct {
	SubId      uint64
	OldBalance *big.Int
	NewBalance *big.Int
	Raw        types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionFundedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionFundedIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionFunded, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionFunded)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV25SubscriptionFunded, error) {
	event := new(VRFCoordinatorV25SubscriptionFunded)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionFundedWithEthIterator struct {
	Event *VRFCoordinatorV25SubscriptionFundedWithEth

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionFundedWithEthIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionFundedWithEth)
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
		it.Event = new(VRFCoordinatorV25SubscriptionFundedWithEth)
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

func (it *VRFCoordinatorV25SubscriptionFundedWithEthIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionFundedWithEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionFundedWithEth struct {
	SubId         uint64
	OldEthBalance *big.Int
	NewEthBalance *big.Int
	Raw           types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionFundedWithEth(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionFundedWithEthIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionFundedWithEth", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionFundedWithEthIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionFundedWithEth", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionFundedWithEth(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionFundedWithEth, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionFundedWithEth", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionFundedWithEth)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionFundedWithEth", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionFundedWithEth(log types.Log) (*VRFCoordinatorV25SubscriptionFundedWithEth, error) {
	event := new(VRFCoordinatorV25SubscriptionFundedWithEth)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionFundedWithEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator struct {
	Event *VRFCoordinatorV25SubscriptionOwnerTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionOwnerTransferRequested)
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
		it.Event = new(VRFCoordinatorV25SubscriptionOwnerTransferRequested)
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

func (it *VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionOwnerTransferRequested struct {
	SubId uint64
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionOwnerTransferRequested, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionOwnerTransferRequested)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV25SubscriptionOwnerTransferRequested, error) {
	event := new(VRFCoordinatorV25SubscriptionOwnerTransferRequested)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV25SubscriptionOwnerTransferredIterator struct {
	Event *VRFCoordinatorV25SubscriptionOwnerTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV25SubscriptionOwnerTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV25SubscriptionOwnerTransferred)
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
		it.Event = new(VRFCoordinatorV25SubscriptionOwnerTransferred)
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

func (it *VRFCoordinatorV25SubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV25SubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV25SubscriptionOwnerTransferred struct {
	SubId uint64
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionOwnerTransferredIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV25SubscriptionOwnerTransferredIterator{contract: _VRFCoordinatorV25.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionOwnerTransferred, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV25.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV25SubscriptionOwnerTransferred)
				if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
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

func (_VRFCoordinatorV25 *VRFCoordinatorV25Filterer) ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV25SubscriptionOwnerTransferred, error) {
	event := new(VRFCoordinatorV25SubscriptionOwnerTransferred)
	if err := _VRFCoordinatorV25.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetSubscription struct {
	Balance    *big.Int
	EthBalance *big.Int
	Owner      common.Address
	Consumers  []common.Address
}
type SConfig struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
}
type SFeeConfig struct {
	FulfillmentFlatFeeLinkPPM uint32
	FulfillmentFlatFeeEthPPM  uint32
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFCoordinatorV25.abi.Events["ConfigSet"].ID:
		return _VRFCoordinatorV25.ParseConfigSet(log)
	case _VRFCoordinatorV25.abi.Events["EthFundsRecovered"].ID:
		return _VRFCoordinatorV25.ParseEthFundsRecovered(log)
	case _VRFCoordinatorV25.abi.Events["FundsRecovered"].ID:
		return _VRFCoordinatorV25.ParseFundsRecovered(log)
	case _VRFCoordinatorV25.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFCoordinatorV25.ParseOwnershipTransferRequested(log)
	case _VRFCoordinatorV25.abi.Events["OwnershipTransferred"].ID:
		return _VRFCoordinatorV25.ParseOwnershipTransferred(log)
	case _VRFCoordinatorV25.abi.Events["ProvingKeyDeregistered"].ID:
		return _VRFCoordinatorV25.ParseProvingKeyDeregistered(log)
	case _VRFCoordinatorV25.abi.Events["ProvingKeyRegistered"].ID:
		return _VRFCoordinatorV25.ParseProvingKeyRegistered(log)
	case _VRFCoordinatorV25.abi.Events["RandomWordsFulfilled"].ID:
		return _VRFCoordinatorV25.ParseRandomWordsFulfilled(log)
	case _VRFCoordinatorV25.abi.Events["RandomWordsRequested"].ID:
		return _VRFCoordinatorV25.ParseRandomWordsRequested(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionCanceled"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionCanceled(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionConsumerAdded"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionConsumerAdded(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionConsumerRemoved"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionConsumerRemoved(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionCreated"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionCreated(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionFunded"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionFunded(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionFundedWithEth"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionFundedWithEth(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionOwnerTransferRequested"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionOwnerTransferRequested(log)
	case _VRFCoordinatorV25.abi.Events["SubscriptionOwnerTransferred"].ID:
		return _VRFCoordinatorV25.ParseSubscriptionOwnerTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFCoordinatorV25ConfigSet) Topic() common.Hash {
	return common.HexToHash("0x777357bb93f63d088f18112d3dba38457aec633eb8f1341e1d418380ad328e78")
}

func (VRFCoordinatorV25EthFundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x879c9ea2b9d5345b84ccd12610b032602808517cebdb795007f3dcb4df377317")
}

func (VRFCoordinatorV25FundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600")
}

func (VRFCoordinatorV25OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFCoordinatorV25OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFCoordinatorV25ProvingKeyDeregistered) Topic() common.Hash {
	return common.HexToHash("0x72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d")
}

func (VRFCoordinatorV25ProvingKeyRegistered) Topic() common.Hash {
	return common.HexToHash("0xe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b8")
}

func (VRFCoordinatorV25RandomWordsFulfilled) Topic() common.Hash {
	return common.HexToHash("0x04624b1adcd18b732732ed3009cbdcc7dfc5f0340105149d0c193e8361d955e1")
}

func (VRFCoordinatorV25RandomWordsRequested) Topic() common.Hash {
	return common.HexToHash("0xad8d0b36e0dd1bc233432f9e3c28d376b9713fc9d26ac2dc972077be458d86fe")
}

func (VRFCoordinatorV25SubscriptionCanceled) Topic() common.Hash {
	return common.HexToHash("0xb3a76e2c52bb7a484e325118e4cadc4346fdaa16a7df0621c5525e1f7ab4d495")
}

func (VRFCoordinatorV25SubscriptionConsumerAdded) Topic() common.Hash {
	return common.HexToHash("0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0")
}

func (VRFCoordinatorV25SubscriptionConsumerRemoved) Topic() common.Hash {
	return common.HexToHash("0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b")
}

func (VRFCoordinatorV25SubscriptionCreated) Topic() common.Hash {
	return common.HexToHash("0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf")
}

func (VRFCoordinatorV25SubscriptionFunded) Topic() common.Hash {
	return common.HexToHash("0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8")
}

func (VRFCoordinatorV25SubscriptionFundedWithEth) Topic() common.Hash {
	return common.HexToHash("0x4e4421a74ab9b76003de9d9527153d002f1338c36c1bbc180069671631cf0958")
}

func (VRFCoordinatorV25SubscriptionOwnerTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be")
}

func (VRFCoordinatorV25SubscriptionOwnerTransferred) Topic() common.Hash {
	return common.HexToHash("0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0")
}

func (_VRFCoordinatorV25 *VRFCoordinatorV25) Address() common.Address {
	return _VRFCoordinatorV25.address
}

type VRFCoordinatorV25Interface interface {
	BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error)

	LINK(opts *bind.CallOpts) (common.Address, error)

	LINKETHFEED(opts *bind.CallOpts) (common.Address, error)

	MAXCONSUMERS(opts *bind.CallOpts) (uint16, error)

	MAXNUMWORDS(opts *bind.CallOpts) (uint32, error)

	MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error)

	GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error)

	GetSubscription(opts *bind.CallOpts, subId uint64) (GetSubscription,

		error)

	HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingRequestExists(opts *bind.CallOpts, subId uint64) (bool, error)

	SConfig(opts *bind.CallOpts) (SConfig,

		error)

	SCurrentSubId(opts *bind.CallOpts) (uint64, error)

	SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error)

	SFeeConfig(opts *bind.CallOpts) (SFeeConfig,

		error)

	SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error)

	SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error)

	SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error)

	STotalBalance(opts *bind.CallOpts) (*big.Int, error)

	STotalEthBalance(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error)

	AddConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error)

	CancelSubscription(opts *bind.TransactOpts, subId uint64, to common.Address) (*types.Transaction, error)

	CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error)

	DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error)

	FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV25RequestCommitment) (*types.Transaction, error)

	FundSubscriptionWithEth(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error)

	OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error)

	OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	OracleWithdrawEth(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	OwnerCancelSubscription(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error)

	RecoverEthFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error)

	RemoveConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error)

	RequestRandomWords(opts *bind.TransactOpts, keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32, payInEth bool) (*types.Transaction, error)

	RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64, newOwner common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV25FeeConfig) (*types.Transaction, error)

	SetLINK(opts *bind.TransactOpts, link common.Address) (*types.Transaction, error)

	SetLinkEthFeed(opts *bind.TransactOpts, linkEthFeed common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV25ConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VRFCoordinatorV25ConfigSet, error)

	FilterEthFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV25EthFundsRecoveredIterator, error)

	WatchEthFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25EthFundsRecovered) (event.Subscription, error)

	ParseEthFundsRecovered(log types.Log) (*VRFCoordinatorV25EthFundsRecovered, error)

	FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV25FundsRecoveredIterator, error)

	WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25FundsRecovered) (event.Subscription, error)

	ParseFundsRecovered(log types.Log) (*VRFCoordinatorV25FundsRecovered, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV25OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV25OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV25OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV25OwnershipTransferred, error)

	FilterProvingKeyDeregistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV25ProvingKeyDeregisteredIterator, error)

	WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ProvingKeyDeregistered, oracle []common.Address) (event.Subscription, error)

	ParseProvingKeyDeregistered(log types.Log) (*VRFCoordinatorV25ProvingKeyDeregistered, error)

	FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV25ProvingKeyRegisteredIterator, error)

	WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25ProvingKeyRegistered, oracle []common.Address) (event.Subscription, error)

	ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV25ProvingKeyRegistered, error)

	FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*VRFCoordinatorV25RandomWordsFulfilledIterator, error)

	WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25RandomWordsFulfilled, requestId []*big.Int) (event.Subscription, error)

	ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV25RandomWordsFulfilled, error)

	FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []uint64, sender []common.Address) (*VRFCoordinatorV25RandomWordsRequestedIterator, error)

	WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25RandomWordsRequested, keyHash [][32]byte, subId []uint64, sender []common.Address) (event.Subscription, error)

	ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV25RandomWordsRequested, error)

	FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionCanceledIterator, error)

	WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionCanceled, subId []uint64) (event.Subscription, error)

	ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV25SubscriptionCanceled, error)

	FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionConsumerAddedIterator, error)

	WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionConsumerAdded, subId []uint64) (event.Subscription, error)

	ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV25SubscriptionConsumerAdded, error)

	FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionConsumerRemovedIterator, error)

	WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionConsumerRemoved, subId []uint64) (event.Subscription, error)

	ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV25SubscriptionConsumerRemoved, error)

	FilterSubscriptionCreated(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionCreatedIterator, error)

	WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionCreated, subId []uint64) (event.Subscription, error)

	ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV25SubscriptionCreated, error)

	FilterSubscriptionFunded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionFundedIterator, error)

	WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionFunded, subId []uint64) (event.Subscription, error)

	ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV25SubscriptionFunded, error)

	FilterSubscriptionFundedWithEth(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionFundedWithEthIterator, error)

	WatchSubscriptionFundedWithEth(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionFundedWithEth, subId []uint64) (event.Subscription, error)

	ParseSubscriptionFundedWithEth(log types.Log) (*VRFCoordinatorV25SubscriptionFundedWithEth, error)

	FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionOwnerTransferRequestedIterator, error)

	WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionOwnerTransferRequested, subId []uint64) (event.Subscription, error)

	ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV25SubscriptionOwnerTransferRequested, error)

	FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV25SubscriptionOwnerTransferredIterator, error)

	WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV25SubscriptionOwnerTransferred, subId []uint64) (event.Subscription, error)

	ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV25SubscriptionOwnerTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
