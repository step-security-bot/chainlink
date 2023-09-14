package contracts

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/vrf_coordinator_v2_5"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/vrf_v2_5_load_test_with_metrics"
)

type EthereumVRFCoordinatorV2_5 struct {
	address     *common.Address
	client      blockchain.EVMClient
	coordinator *vrf_coordinator_v2_5.VRFCoordinatorV25
}

// EthereumVRFv2_5LoadTestConsumer represents VRFv2_5 consumer contract for performing Load Tests
type EthereumVRFv2_5LoadTestConsumer struct {
	address  *common.Address
	client   blockchain.EVMClient
	consumer *vrf_v2_5_load_test_with_metrics.VRFV25LoadTestWithMetrics
}

// DeployVRFCoordinatorV2_5 deploys VRFV2_5 coordinator contract
func (e *EthereumContractDeployer) DeployVRFCoordinatorV2_5(bhsAddr string) (VRFCoordinatorV2_5, error) {
	address, _, instance, err := e.client.DeployContract("VRFCoordinatorV2_5", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return vrf_coordinator_v2_5.DeployVRFCoordinatorV25(auth, backend, common.HexToAddress(bhsAddr))
	})
	if err != nil {
		return nil, err
	}
	return &EthereumVRFCoordinatorV2_5{
		client:      e.client,
		coordinator: instance.(*vrf_coordinator_v2_5.VRFCoordinatorV25),
		address:     address,
	}, err
}

func (v *EthereumVRFCoordinatorV2_5) Address() string {
	return v.address.Hex()
}

func (v *EthereumVRFCoordinatorV2_5) HashOfKey(ctx context.Context, pubKey [2]*big.Int) ([32]byte, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(v.client.GetDefaultWallet().Address()),
		Context: ctx,
	}
	hash, err := v.coordinator.HashOfKey(opts, pubKey)
	if err != nil {
		return [32]byte{}, err
	}
	return hash, nil
}

func (v *EthereumVRFCoordinatorV2_5) GetSubscription(ctx context.Context, subID *big.Int) (vrf_coordinator_v2_5.GetSubscription, error) {
	opts := &bind.CallOpts{
		From:    common.HexToAddress(v.client.GetDefaultWallet().Address()),
		Context: ctx,
	}
	subscription, err := v.coordinator.GetSubscription(opts, subID)
	if err != nil {
		return vrf_coordinator_v2_5.GetSubscription{}, err
	}
	return subscription, nil
}

func (v *EthereumVRFCoordinatorV2_5) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig vrf_coordinator_v2_5.VRFCoordinatorV25FeeConfig) error {
	opts, err := v.client.TransactionOpts(v.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	tx, err := v.coordinator.SetConfig(
		opts,
		minimumRequestConfirmations,
		maxGasLimit,
		stalenessSeconds,
		gasAfterPaymentCalculation,
		fallbackWeiPerUnitLink,
		feeConfig,
	)
	if err != nil {
		return err
	}
	return v.client.ProcessTransaction(tx)
}

func (v *EthereumVRFCoordinatorV2_5) SetLINKAndLINKNativeFeed(linkAddress string, linkNativeFeedAddress string) error {
	opts, err := v.client.TransactionOpts(v.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	tx, err := v.coordinator.SetLINKAndLINKNativeFeed(
		opts,
		common.HexToAddress(linkAddress),
		common.HexToAddress(linkNativeFeedAddress),
	)
	if err != nil {
		return err
	}
	return v.client.ProcessTransaction(tx)
}

func (v *EthereumVRFCoordinatorV2_5) RegisterProvingKey(
	oracleAddr string,
	publicProvingKey [2]*big.Int,
) error {
	opts, err := v.client.TransactionOpts(v.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	tx, err := v.coordinator.RegisterProvingKey(opts, common.HexToAddress(oracleAddr), publicProvingKey)
	if err != nil {
		return err
	}
	return v.client.ProcessTransaction(tx)
}

func (v *EthereumVRFCoordinatorV2_5) CreateSubscription() error {
	opts, err := v.client.TransactionOpts(v.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	tx, err := v.coordinator.CreateSubscription(opts)
	if err != nil {
		return err
	}
	return v.client.ProcessTransaction(tx)
}

func (v *EthereumVRFCoordinatorV2_5) AddConsumer(subId *big.Int, consumerAddress string) error {
	opts, err := v.client.TransactionOpts(v.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	tx, err := v.coordinator.AddConsumer(
		opts,
		subId,
		common.HexToAddress(consumerAddress),
	)
	if err != nil {
		return err
	}
	return v.client.ProcessTransaction(tx)
}

func (v *EthereumVRFCoordinatorV2_5) FundSubscriptionWithNative(subId *big.Int, nativeTokenAmount *big.Int) error {
	opts, err := v.client.TransactionOpts(v.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	opts.Value = nativeTokenAmount
	tx, err := v.coordinator.FundSubscriptionWithNative(
		opts,
		subId,
	)
	if err != nil {
		return err
	}
	return v.client.ProcessTransaction(tx)
}

func (v *EthereumVRFCoordinatorV2_5) FindSubscriptionID() (*big.Int, error) {
	owner := v.client.GetDefaultWallet().Address()
	subscriptionIterator, err := v.coordinator.FilterSubscriptionCreated(
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if !subscriptionIterator.Next() {
		return nil, fmt.Errorf("expected at least 1 subID for the given owner %s", owner)
	}

	return subscriptionIterator.Event.SubId, nil
}

func (v *EthereumVRFCoordinatorV2_5) WaitForRandomWordsFulfilledEvent(subID []*big.Int, requestID []*big.Int, timeout time.Duration) (*vrf_coordinator_v2_5.VRFCoordinatorV25RandomWordsFulfilled, error) {
	randomWordsFulfilledEventsChannel := make(chan *vrf_coordinator_v2_5.VRFCoordinatorV25RandomWordsFulfilled)
	subscription, err := v.coordinator.WatchRandomWordsFulfilled(nil, randomWordsFulfilledEventsChannel, requestID, subID)
	if err != nil {
		return nil, err
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case err := <-subscription.Err():
			return nil, err
		case <-time.After(timeout):
			return nil, fmt.Errorf("timeout waiting for RandomWordsFulfilled event")
		case randomWordsFulfilledEvent := <-randomWordsFulfilledEventsChannel:
			return randomWordsFulfilledEvent, nil
		}
	}
}

func (v *EthereumVRFv2_5LoadTestConsumer) Address() string {
	return v.address.Hex()
}
func (v *EthereumVRFv2_5LoadTestConsumer) RequestRandomness(keyHash [32]byte, subID *big.Int, requestConfirmations uint16, callbackGasLimit uint32, nativePayment bool, numWords uint32, requestCount uint16) error {
	opts, err := v.client.TransactionOpts(v.client.GetDefaultWallet())
	if err != nil {
		return err
	}
	tx, err := v.consumer.RequestRandomWords(opts, subID, requestConfirmations, keyHash, callbackGasLimit, nativePayment, numWords, requestCount)
	if err != nil {
		return err
	}
	return v.client.ProcessTransaction(tx)
}
func (v *EthereumVRFv2_5LoadTestConsumer) GetRequestStatus(ctx context.Context, requestID *big.Int) (vrf_v2_5_load_test_with_metrics.GetRequestStatus, error) {
	return v.consumer.GetRequestStatus(&bind.CallOpts{
		From:    common.HexToAddress(v.client.GetDefaultWallet().Address()),
		Context: ctx,
	}, requestID)
}

func (v *EthereumVRFv2_5LoadTestConsumer) GetLastRequestId(ctx context.Context) (*big.Int, error) {
	return v.consumer.SLastRequestId(&bind.CallOpts{
		From:    common.HexToAddress(v.client.GetDefaultWallet().Address()),
		Context: ctx,
	})
}

func (v *EthereumVRFv2_5LoadTestConsumer) GetLoadTestMetrics(ctx context.Context) (*VRFLoadTestMetrics, error) {
	requestCount, err := v.consumer.SRequestCount(&bind.CallOpts{
		From:    common.HexToAddress(v.client.GetDefaultWallet().Address()),
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	fulfilmentCount, err := v.consumer.SResponseCount(&bind.CallOpts{
		From:    common.HexToAddress(v.client.GetDefaultWallet().Address()),
		Context: ctx,
	})

	if err != nil {
		return nil, err
	}
	averageFulfillmentInMillions, err := v.consumer.SAverageFulfillmentInMillions(&bind.CallOpts{
		From:    common.HexToAddress(v.client.GetDefaultWallet().Address()),
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	slowestFulfillment, err := v.consumer.SSlowestFulfillment(&bind.CallOpts{
		From:    common.HexToAddress(v.client.GetDefaultWallet().Address()),
		Context: ctx,
	})

	if err != nil {
		return nil, err
	}
	fastestFulfillment, err := v.consumer.SFastestFulfillment(&bind.CallOpts{
		From:    common.HexToAddress(v.client.GetDefaultWallet().Address()),
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return &VRFLoadTestMetrics{
		requestCount,
		fulfilmentCount,
		averageFulfillmentInMillions,
		slowestFulfillment,
		fastestFulfillment,
	}, nil
}
