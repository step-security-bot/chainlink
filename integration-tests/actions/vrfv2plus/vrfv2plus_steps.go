package vrfv2plus

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/vrfv2plus_wrapper_load_test_consumer"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/actions/vrfv2plus/vrfv2plus_config"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/docker/test_env"
	"github.com/smartcontractkit/chainlink/integration-tests/types/config/node"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/vrf_coordinator_v2_5"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/vrf_v2plus_upgraded_version"
	chainlinkutils "github.com/smartcontractkit/chainlink/v2/core/utils"
)

var (
	ErrNodePrimaryKey                              = "error getting node's primary ETH key"
	ErrCreatingProvingKeyHash                      = "error creating a keyHash from the proving key"
	ErrRegisteringProvingKey                       = "error registering a proving key on Coordinator contract"
	ErrRegisterProvingKey                          = "error registering proving keys"
	ErrEncodingProvingKey                          = "error encoding proving key"
	ErrCreatingVRFv2PlusKey                        = "error creating VRFv2Plus key"
	ErrDeployBlockHashStore                        = "error deploying blockhash store"
	ErrDeployCoordinator                           = "error deploying VRF CoordinatorV2Plus"
	ErrAdvancedConsumer                            = "error deploying VRFv2Plus Advanced Consumer"
	ErrABIEncodingFunding                          = "error Abi encoding subscriptionID"
	ErrSendingLinkToken                            = "error sending Link token"
	ErrCreatingVRFv2PlusJob                        = "error creating VRFv2Plus job"
	ErrParseJob                                    = "error parsing job definition"
	ErrDeployVRFV2_5Contracts                      = "error deploying VRFV2_5 contracts"
	ErrSetVRFCoordinatorConfig                     = "error setting config for VRF Coordinator contract"
	ErrCreateVRFSubscription                       = "error creating VRF Subscription"
	ErrFindSubID                                   = "error finding created subscription ID"
	ErrAddConsumerToSub                            = "error adding consumer to VRF Subscription"
	ErrFundSubWithNativeToken                      = "error funding subscription with native token"
	ErrSetLinkNativeLinkFeed                       = "error setting Link and ETH/LINK feed for VRF Coordinator contract"
	ErrFundSubWithLinkToken                        = "error funding subscription with Link tokens"
	ErrCreateVRFV2PlusJobs                         = "error creating VRF V2 Plus Jobs"
	ErrGetPrimaryKey                               = "error getting primary ETH key address"
	ErrRestartCLNode                               = "error restarting CL node"
	ErrWaitTXsComplete                             = "error waiting for TXs to complete"
	ErrRequestRandomness                           = "error requesting randomness"
	ErrRequestRandomnessDirectFundingLinkPayment   = "error requesting randomness with direct funding and link payment"
	ErrRequestRandomnessDirectFundingNativePayment = "error requesting randomness with direct funding and native payment"

	ErrWaitRandomWordsRequestedEvent = "error waiting for RandomWordsRequested event"
	ErrWaitRandomWordsFulfilledEvent = "error waiting for RandomWordsFulfilled event"
	ErrLinkTotalBalance              = "error waiting for RandomWordsFulfilled event"
	ErrNativeTokenBalance            = "error waiting for RandomWordsFulfilled event"
	ErrDeployWrapper                 = "error deploying VRFV2PlusWrapper"
)

func DeployVRFV2_5Contracts(
	contractDeployer contracts.ContractDeployer,
	chainClient blockchain.EVMClient,
	consumerContractsAmount int,
) (*VRFV2_5Contracts, error) {
	bhs, err := contractDeployer.DeployBlockhashStore()
	if err != nil {
		return nil, errors.Wrap(err, ErrDeployBlockHashStore)
	}
	err = chainClient.WaitForEvents()
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	coordinator, err := contractDeployer.DeployVRFCoordinatorV2_5(bhs.Address())
	if err != nil {
		return nil, errors.Wrap(err, ErrDeployCoordinator)
	}
	err = chainClient.WaitForEvents()
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	consumers, err := DeployVRFV2PlusConsumers(contractDeployer, coordinator, consumerContractsAmount)
	if err != nil {
		return nil, err
	}
	err = chainClient.WaitForEvents()
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	return &VRFV2_5Contracts{coordinator, bhs, consumers}, nil
}

func DeployVRFV2PlusDirectFundingContracts(
	contractDeployer contracts.ContractDeployer,
	chainClient blockchain.EVMClient,
	linkTokenAddress string,
	linkEthFeedAddress string,
	coordinator contracts.VRFCoordinatorV2_5,
	consumerContractsAmount int,
) (*VRFV2PlusWrapperContracts, error) {

	vrfv2PlusWrapper, err := contractDeployer.DeployVRFV2PlusWrapper(linkTokenAddress, linkEthFeedAddress, coordinator.Address())
	if err != nil {
		return nil, errors.Wrap(err, ErrDeployWrapper)
	}
	err = chainClient.WaitForEvents()
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitTXsComplete)
	}

	consumers, err := DeployVRFV2PlusWrapperConsumers(contractDeployer, linkTokenAddress, vrfv2PlusWrapper, consumerContractsAmount)
	if err != nil {
		return nil, err
	}
	err = chainClient.WaitForEvents()
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	return &VRFV2PlusWrapperContracts{vrfv2PlusWrapper, consumers}, nil
}

func DeployVRFV2PlusConsumers(contractDeployer contracts.ContractDeployer, coordinator contracts.VRFCoordinatorV2_5, consumerContractsAmount int) ([]contracts.VRFv2PlusLoadTestConsumer, error) {
	var consumers []contracts.VRFv2PlusLoadTestConsumer
	for i := 1; i <= consumerContractsAmount; i++ {
		loadTestConsumer, err := contractDeployer.DeployVRFv2PlusLoadTestConsumer(coordinator.Address())
		if err != nil {
			return nil, errors.Wrap(err, ErrAdvancedConsumer)
		}
		consumers = append(consumers, loadTestConsumer)
	}
	return consumers, nil
}

func DeployVRFV2PlusWrapperConsumers(contractDeployer contracts.ContractDeployer, linkTokenAddress string, vrfV2PlusWrapper contracts.VRFV2PlusWrapper, consumerContractsAmount int) ([]contracts.VRFv2PlusWrapperLoadTestConsumer, error) {
	var consumers []contracts.VRFv2PlusWrapperLoadTestConsumer
	for i := 1; i <= consumerContractsAmount; i++ {
		loadTestConsumer, err := contractDeployer.DeployVRFV2PlusWrapperLoadTestConsumer(linkTokenAddress, vrfV2PlusWrapper.Address())
		if err != nil {
			return nil, errors.Wrap(err, ErrAdvancedConsumer)
		}
		consumers = append(consumers, loadTestConsumer)
	}
	return consumers, nil
}

func CreateVRFV2PlusJob(
	chainlinkNode *client.ChainlinkClient,
	coordinatorAddress string,
	nativeTokenPrimaryKeyAddress string,
	pubKeyCompressed string,
	chainID string,
	minIncomingConfirmations uint16,
) (*client.Job, error) {
	jobUUID := uuid.New()
	os := &client.VRFV2PlusTxPipelineSpec{
		Address: coordinatorAddress,
	}
	ost, err := os.String()
	if err != nil {
		return nil, errors.Wrap(err, ErrParseJob)
	}

	job, err := chainlinkNode.MustCreateJob(&client.VRFV2PlusJobSpec{
		Name:                     fmt.Sprintf("vrf-v2-plus-%s", jobUUID),
		CoordinatorAddress:       coordinatorAddress,
		FromAddresses:            []string{nativeTokenPrimaryKeyAddress},
		EVMChainID:               chainID,
		MinIncomingConfirmations: int(minIncomingConfirmations),
		PublicKey:                pubKeyCompressed,
		ExternalJobID:            jobUUID.String(),
		ObservationSource:        ost,
		BatchFulfillmentEnabled:  false,
	})
	if err != nil {
		return nil, errors.Wrap(err, ErrCreatingVRFv2PlusJob)
	}

	return job, nil
}

func VRFV2_5RegisterProvingKey(
	vrfKey *client.VRFKey,
	oracleAddress string,
	coordinator contracts.VRFCoordinatorV2_5,
) (VRFV2PlusEncodedProvingKey, error) {
	provingKey, err := actions.EncodeOnChainVRFProvingKey(*vrfKey)
	if err != nil {
		return VRFV2PlusEncodedProvingKey{}, errors.Wrap(err, ErrEncodingProvingKey)
	}
	err = coordinator.RegisterProvingKey(
		oracleAddress,
		provingKey,
	)
	if err != nil {
		return VRFV2PlusEncodedProvingKey{}, errors.Wrap(err, ErrRegisterProvingKey)
	}
	return provingKey, nil
}

func VRFV2PlusUpgradedVersionRegisterProvingKey(
	vrfKey *client.VRFKey,
	oracleAddress string,
	coordinator contracts.VRFCoordinatorV2PlusUpgradedVersion,
) (VRFV2PlusEncodedProvingKey, error) {
	provingKey, err := actions.EncodeOnChainVRFProvingKey(*vrfKey)
	if err != nil {
		return VRFV2PlusEncodedProvingKey{}, errors.Wrap(err, ErrEncodingProvingKey)
	}
	err = coordinator.RegisterProvingKey(
		oracleAddress,
		provingKey,
	)
	if err != nil {
		return VRFV2PlusEncodedProvingKey{}, errors.Wrap(err, ErrRegisterProvingKey)
	}
	return provingKey, nil
}

func FundVRFCoordinatorV2_5Subscription(linkToken contracts.LinkToken, coordinator contracts.VRFCoordinatorV2_5, chainClient blockchain.EVMClient, subscriptionID *big.Int, linkFundingAmount *big.Int) error {
	encodedSubId, err := chainlinkutils.ABIEncode(`[{"type":"uint256"}]`, subscriptionID)
	if err != nil {
		return errors.Wrap(err, ErrABIEncodingFunding)
	}
	_, err = linkToken.TransferAndCall(coordinator.Address(), big.NewInt(0).Mul(linkFundingAmount, big.NewInt(1e18)), encodedSubId)
	if err != nil {
		return errors.Wrap(err, ErrSendingLinkToken)
	}
	return chainClient.WaitForEvents()
}

func SetupVRFV2_5Environment(
	env *test_env.CLClusterTestEnv,
	vrfv2PlusConfig vrfv2plus_config.VRFV2PlusConfig,
	linkToken contracts.LinkToken,
	mockNativeLINKFeed contracts.MockETHLINKFeed,
	consumerContractsAmount int,
) (*VRFV2_5Contracts, *big.Int, *VRFV2PlusData, error) {

	vrfv2_5Contracts, err := DeployVRFV2_5Contracts(env.ContractDeployer, env.EVMClient, consumerContractsAmount)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrDeployVRFV2_5Contracts)
	}

	err = vrfv2_5Contracts.Coordinator.SetConfig(
		vrfv2PlusConfig.MinimumConfirmations,
		vrfv2PlusConfig.MaxGasLimitCoordinatorConfig,
		vrfv2PlusConfig.StalenessSeconds,
		vrfv2PlusConfig.GasAfterPaymentCalculation,
		big.NewInt(vrfv2PlusConfig.LinkNativeFeedResponse),
		vrf_coordinator_v2_5.VRFCoordinatorV25FeeConfig{
			FulfillmentFlatFeeLinkPPM:   vrfv2PlusConfig.FulfillmentFlatFeeLinkPPM,
			FulfillmentFlatFeeNativePPM: vrfv2PlusConfig.FulfillmentFlatFeeNativePPM,
		},
	)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrSetVRFCoordinatorConfig)
	}

	subID, err := CreateSubAndFindSubID(env, vrfv2_5Contracts.Coordinator)
	if err != nil {
		return nil, nil, nil, err
	}

	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	for _, consumer := range vrfv2_5Contracts.LoadTestConsumers {
		err = vrfv2_5Contracts.Coordinator.AddConsumer(subID, consumer.Address())
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, ErrAddConsumerToSub)
		}
	}

	err = vrfv2_5Contracts.Coordinator.SetLINKAndLINKNativeFeed(linkToken.Address(), mockNativeLINKFeed.Address())
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrSetLinkNativeLinkFeed)
	}
	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	err = FundSubscription(env, vrfv2PlusConfig, linkToken, vrfv2_5Contracts.Coordinator, subID)
	if err != nil {
		return nil, nil, nil, err
	}

	vrfKey, err := env.ClCluster.NodeAPIs()[0].MustCreateVRFKey()
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrCreatingVRFv2PlusKey)
	}
	pubKeyCompressed := vrfKey.Data.ID

	nativeTokenPrimaryKeyAddress, err := env.ClCluster.NodeAPIs()[0].PrimaryEthAddress()
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrNodePrimaryKey)
	}
	provingKey, err := VRFV2_5RegisterProvingKey(vrfKey, nativeTokenPrimaryKeyAddress, vrfv2_5Contracts.Coordinator)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrRegisteringProvingKey)
	}
	keyHash, err := vrfv2_5Contracts.Coordinator.HashOfKey(context.Background(), provingKey)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrCreatingProvingKeyHash)
	}

	chainID := env.EVMClient.GetChainID()

	job, err := CreateVRFV2PlusJob(
		env.ClCluster.NodeAPIs()[0],
		vrfv2_5Contracts.Coordinator.Address(),
		nativeTokenPrimaryKeyAddress,
		pubKeyCompressed,
		chainID.String(),
		vrfv2PlusConfig.MinimumConfirmations,
	)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrCreateVRFV2PlusJobs)
	}

	// this part is here because VRFv2 can work with only a specific key
	// [[EVM.KeySpecific]]
	//	Key = '...'
	addr, err := env.ClCluster.Nodes[0].API.PrimaryEthAddress()
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrGetPrimaryKey)
	}
	nodeConfig := node.NewConfig(env.ClCluster.Nodes[0].NodeConfig,
		node.WithVRFv2EVMEstimator(addr),
	)
	err = env.ClCluster.Nodes[0].Restart(nodeConfig)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, ErrRestartCLNode)
	}

	vrfv2PlusKeyData := VRFV2PlusKeyData{
		VRFKey:            vrfKey,
		EncodedProvingKey: provingKey,
		KeyHash:           keyHash,
	}

	data := VRFV2PlusData{
		vrfv2PlusKeyData,
		job,
		nativeTokenPrimaryKeyAddress,
		chainID,
	}

	return vrfv2_5Contracts, subID, &data, nil
}

func SetupVRFV2PlusWrapperEnvironment(
	env *test_env.CLClusterTestEnv,
	vrfv2PlusConfig vrfv2plus_config.VRFV2PlusConfig,
	linkToken contracts.LinkToken,
	mockNativeLINKFeed contracts.MockETHLINKFeed,
	coordinator contracts.VRFCoordinatorV2_5,
	keyHash [32]byte,
	wrapperConsumerContractsAmount int,
) (*VRFV2PlusWrapperContracts, *big.Int, error) {

	wrapperContracts, err := DeployVRFV2PlusDirectFundingContracts(
		env.ContractDeployer,
		env.EVMClient,
		linkToken.Address(),
		mockNativeLINKFeed.Address(),
		coordinator,
		wrapperConsumerContractsAmount,
	)
	if err != nil {
		return nil, nil, err
	}

	err = env.EVMClient.WaitForEvents()

	if err != nil {
		return nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	err = wrapperContracts.VRFV2PlusWrapper.SetConfig(
		vrfv2PlusConfig.WrapperGasOverhead,
		vrfv2PlusConfig.CoordinatorGasOverhead,
		vrfv2PlusConfig.WrapperPremiumPercentage,
		keyHash,
		vrfv2PlusConfig.WrapperMaxNumberOfWords,
		vrfv2PlusConfig.StalenessSeconds,
		big.NewInt(vrfv2PlusConfig.FallbackWeiPerUnitLink),
		vrfv2PlusConfig.FulfillmentFlatFeeLinkPPM,
		vrfv2PlusConfig.FulfillmentFlatFeeNativePPM,
	)
	if err != nil {
		return nil, nil, err
	}

	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}

	//fund sub
	wrapperSubID, err := wrapperContracts.VRFV2PlusWrapper.GetSubID(context.Background())
	if err != nil {
		return nil, nil, err
	}

	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}

	err = FundSubscription(env, vrfv2PlusConfig, linkToken, coordinator, wrapperSubID)
	if err != nil {
		return nil, nil, err
	}

	//fund consumer with Link
	err = linkToken.Transfer(
		wrapperContracts.LoadTestConsumers[0].Address(),
		big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(vrfv2PlusConfig.WrapperConsumerFundingAmountLink)),
	)
	if err != nil {
		return nil, nil, err
	}
	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}

	//fund consumer with Eth
	err = wrapperContracts.LoadTestConsumers[0].Fund(big.NewFloat(vrfv2PlusConfig.WrapperConsumerFundingAmountNativeToken))
	if err != nil {
		return nil, nil, err
	}
	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	return wrapperContracts, wrapperSubID, nil
}
func CreateSubAndFindSubID(env *test_env.CLClusterTestEnv, coordinator contracts.VRFCoordinatorV2_5) (*big.Int, error) {
	err := coordinator.CreateSubscription()
	if err != nil {
		return nil, errors.Wrap(err, ErrCreateVRFSubscription)
	}
	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	subID, err := coordinator.FindSubscriptionID()
	if err != nil {
		return nil, errors.Wrap(err, ErrFindSubID)
	}
	return subID, nil
}

func GetUpgradedCoordinatorTotalBalance(coordinator contracts.VRFCoordinatorV2PlusUpgradedVersion) (linkTotalBalance *big.Int, nativeTokenTotalBalance *big.Int, err error) {
	linkTotalBalance, err = coordinator.GetLinkTotalBalance(context.Background())
	if err != nil {
		return nil, nil, errors.Wrap(err, ErrLinkTotalBalance)
	}
	nativeTokenTotalBalance, err = coordinator.GetNativeTokenTotalBalance(context.Background())
	if err != nil {
		return nil, nil, errors.Wrap(err, ErrNativeTokenBalance)
	}
	return
}

func GetCoordinatorTotalBalance(coordinator contracts.VRFCoordinatorV2_5) (linkTotalBalance *big.Int, nativeTokenTotalBalance *big.Int, err error) {
	linkTotalBalance, err = coordinator.GetLinkTotalBalance(context.Background())
	if err != nil {
		return nil, nil, errors.Wrap(err, ErrLinkTotalBalance)
	}
	nativeTokenTotalBalance, err = coordinator.GetNativeTokenTotalBalance(context.Background())
	if err != nil {
		return nil, nil, errors.Wrap(err, ErrNativeTokenBalance)
	}
	return
}

func FundSubscription(
	env *test_env.CLClusterTestEnv,
	vrfv2PlusConfig vrfv2plus_config.VRFV2PlusConfig,
	linkAddress contracts.LinkToken,
	coordinator contracts.VRFCoordinatorV2_5,
	subID *big.Int,
) error {
	//Native Billing
	err := coordinator.FundSubscriptionWithNative(subID, big.NewInt(0).Mul(big.NewInt(vrfv2PlusConfig.SubscriptionFundingAmountNative), big.NewInt(1e18)))
	if err != nil {
		return errors.Wrap(err, ErrFundSubWithNativeToken)
	}

	err = FundVRFCoordinatorV2_5Subscription(linkAddress, coordinator, env.EVMClient, subID, big.NewInt(vrfv2PlusConfig.SubscriptionFundingAmountLink))
	if err != nil {
		return errors.Wrap(err, ErrFundSubWithLinkToken)
	}
	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return errors.Wrap(err, ErrWaitTXsComplete)
	}

	return nil
}

func RequestRandomnessAndWaitForFulfillment(
	consumer contracts.VRFv2PlusLoadTestConsumer,
	coordinator contracts.VRFCoordinatorV2_5,
	vrfv2PlusData *VRFV2PlusData,
	subID *big.Int,
	isNativeBilling bool,
	vrfv2PlusConfig vrfv2plus_config.VRFV2PlusConfig,
	l zerolog.Logger,
) (*vrf_coordinator_v2_5.VRFCoordinatorV25RandomWordsFulfilled, error) {
	logRandRequest(consumer.Address(), coordinator.Address(), subID, isNativeBilling, vrfv2PlusConfig, l)
	_, err := consumer.RequestRandomness(
		vrfv2PlusData.KeyHash,
		subID,
		vrfv2PlusConfig.MinimumConfirmations,
		vrfv2PlusConfig.CallbackGasLimit,
		isNativeBilling,
		vrfv2PlusConfig.NumberOfWords,
		vrfv2PlusConfig.RandomnessRequestCountPerRequest,
	)
	if err != nil {
		return nil, errors.Wrap(err, ErrRequestRandomness)
	}

	return WaitForRequestAndFulfillmentEvents(consumer.Address(), coordinator, vrfv2PlusData, subID, isNativeBilling, l)
}

func RequestRandomnessAndWaitForFulfillmentUpgraded(
	consumer contracts.VRFv2PlusLoadTestConsumer,
	coordinator contracts.VRFCoordinatorV2PlusUpgradedVersion,
	vrfv2PlusData *VRFV2PlusData,
	subID *big.Int,
	isNativeBilling bool,
	vrfv2PlusConfig vrfv2plus_config.VRFV2PlusConfig,
	l zerolog.Logger,
) (*vrf_v2plus_upgraded_version.VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled, error) {
	logRandRequest(consumer.Address(), coordinator.Address(), subID, isNativeBilling, vrfv2PlusConfig, l)
	_, err := consumer.RequestRandomness(
		vrfv2PlusData.KeyHash,
		subID,
		vrfv2PlusConfig.MinimumConfirmations,
		vrfv2PlusConfig.CallbackGasLimit,
		isNativeBilling,
		vrfv2PlusConfig.NumberOfWords,
		vrfv2PlusConfig.RandomnessRequestCountPerRequest,
	)
	if err != nil {
		return nil, errors.Wrap(err, ErrRequestRandomness)
	}

	randomWordsRequestedEvent, err := coordinator.WaitForRandomWordsRequestedEvent(
		[][32]byte{vrfv2PlusData.KeyHash},
		[]*big.Int{subID},
		[]common.Address{common.HexToAddress(consumer.Address())},
		time.Minute*1,
	)
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitRandomWordsRequestedEvent)
	}

	LogRandomnessRequestedEventUpgraded(l, coordinator, randomWordsRequestedEvent)

	randomWordsFulfilledEvent, err := coordinator.WaitForRandomWordsFulfilledEvent(
		[]*big.Int{subID},
		[]*big.Int{randomWordsRequestedEvent.RequestId},
		time.Minute*2,
	)
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitRandomWordsFulfilledEvent)
	}
	LogRandomWordsFulfilledEventUpgraded(l, coordinator, randomWordsFulfilledEvent)

	return randomWordsFulfilledEvent, err
}

func DirectFundingRequestRandomnessAndWaitForFulfillment(
	consumer contracts.VRFv2PlusWrapperLoadTestConsumer,
	coordinator contracts.VRFCoordinatorV2_5,
	vrfv2PlusData *VRFV2PlusData,
	subID *big.Int,
	isNativeBilling bool,
	vrfv2PlusConfig vrfv2plus_config.VRFV2PlusConfig,
	l zerolog.Logger,
) (*vrf_coordinator_v2_5.VRFCoordinatorV25RandomWordsFulfilled, error) {
	logRandRequest(consumer.Address(), coordinator.Address(), subID, isNativeBilling, vrfv2PlusConfig, l)
	if isNativeBilling {
		_, err := consumer.RequestRandomnessNative(
			vrfv2PlusConfig.MinimumConfirmations,
			vrfv2PlusConfig.CallbackGasLimit,
			vrfv2PlusConfig.NumberOfWords,
			vrfv2PlusConfig.RandomnessRequestCountPerRequest,
		)
		if err != nil {
			return nil, errors.Wrap(err, ErrRequestRandomnessDirectFundingNativePayment)
		}
	} else {
		_, err := consumer.RequestRandomness(
			vrfv2PlusConfig.MinimumConfirmations,
			vrfv2PlusConfig.CallbackGasLimit,
			vrfv2PlusConfig.NumberOfWords,
			vrfv2PlusConfig.RandomnessRequestCountPerRequest,
		)
		if err != nil {
			return nil, errors.Wrap(err, ErrRequestRandomnessDirectFundingLinkPayment)
		}
	}
	wrapperAddress, err := consumer.GetWrapper(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "error getting wrapper address")
	}
	return WaitForRequestAndFulfillmentEvents(wrapperAddress.String(), coordinator, vrfv2PlusData, subID, isNativeBilling, l)
}

func WaitForRequestAndFulfillmentEvents(
	consumerAddress string,
	coordinator contracts.VRFCoordinatorV2_5,
	vrfv2PlusData *VRFV2PlusData,
	subID *big.Int,
	isNativeBilling bool,
	l zerolog.Logger,
) (*vrf_coordinator_v2_5.VRFCoordinatorV25RandomWordsFulfilled, error) {
	randomWordsRequestedEvent, err := coordinator.WaitForRandomWordsRequestedEvent(
		[][32]byte{vrfv2PlusData.KeyHash},
		[]*big.Int{subID},
		[]common.Address{common.HexToAddress(consumerAddress)},
		time.Minute*1,
	)
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitRandomWordsRequestedEvent)
	}

	LogRandomnessRequestedEvent(l, coordinator, randomWordsRequestedEvent, isNativeBilling)

	randomWordsFulfilledEvent, err := coordinator.WaitForRandomWordsFulfilledEvent(
		[]*big.Int{subID},
		[]*big.Int{randomWordsRequestedEvent.RequestId},
		time.Minute*2,
	)
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitRandomWordsFulfilledEvent)
	}

	LogRandomWordsFulfilledEvent(l, coordinator, randomWordsFulfilledEvent, isNativeBilling)
	return randomWordsFulfilledEvent, err
}

func LogSubDetails(l zerolog.Logger, subscription vrf_coordinator_v2_5.GetSubscription, subID *big.Int, coordinator contracts.VRFCoordinatorV2_5) {
	l.Debug().
		Str("Coordinator", coordinator.Address()).
		Str("Link Balance", (*assets.Link)(subscription.Balance).Link()).
		Str("Native Token Balance", assets.FormatWei(subscription.NativeBalance)).
		Str("Subscription ID", subID.String()).
		Str("Subscription Owner", subscription.Owner.String()).
		Interface("Subscription Consumers", subscription.Consumers).
		Msg("Subscription Data")
}

func LogRandomnessRequestedEventUpgraded(
	l zerolog.Logger,
	coordinator contracts.VRFCoordinatorV2PlusUpgradedVersion,
	randomWordsRequestedEvent *vrf_v2plus_upgraded_version.VRFCoordinatorV2PlusUpgradedVersionRandomWordsRequested,
) {
	l.Debug().
		Str("Coordinator", coordinator.Address()).
		Str("Request ID", randomWordsRequestedEvent.RequestId.String()).
		Str("Subscription ID", randomWordsRequestedEvent.SubId.String()).
		Str("Sender Address", randomWordsRequestedEvent.Sender.String()).
		Interface("Keyhash", randomWordsRequestedEvent.KeyHash).
		Uint32("Callback Gas Limit", randomWordsRequestedEvent.CallbackGasLimit).
		Uint32("Number of Words", randomWordsRequestedEvent.NumWords).
		Uint16("Minimum Request Confirmations", randomWordsRequestedEvent.MinimumRequestConfirmations).
		Msg("RandomnessRequested Event")
}

func LogRandomWordsFulfilledEventUpgraded(
	l zerolog.Logger,
	coordinator contracts.VRFCoordinatorV2PlusUpgradedVersion,
	randomWordsFulfilledEvent *vrf_v2plus_upgraded_version.VRFCoordinatorV2PlusUpgradedVersionRandomWordsFulfilled,
) {
	l.Debug().
		Str("Coordinator", coordinator.Address()).
		Str("Total Payment in Juels", randomWordsFulfilledEvent.Payment.String()).
		Str("TX Hash", randomWordsFulfilledEvent.Raw.TxHash.String()).
		Str("Subscription ID", randomWordsFulfilledEvent.SubID.String()).
		Str("Request ID", randomWordsFulfilledEvent.RequestId.String()).
		Bool("Success", randomWordsFulfilledEvent.Success).
		Msg("RandomWordsFulfilled Event (TX metadata)")
}

func LogRandomnessRequestedEvent(
	l zerolog.Logger,
	coordinator contracts.VRFCoordinatorV2_5,
	randomWordsRequestedEvent *vrf_coordinator_v2_5.VRFCoordinatorV25RandomWordsRequested,
	isNativeBilling bool,
) {
	l.Debug().
		Str("Coordinator", coordinator.Address()).
		Bool("Native Billing", isNativeBilling).
		Str("Request ID", randomWordsRequestedEvent.RequestId.String()).
		Str("Subscription ID", randomWordsRequestedEvent.SubId.String()).
		Str("Sender Address", randomWordsRequestedEvent.Sender.String()).
		Interface("Keyhash", randomWordsRequestedEvent.KeyHash).
		Uint32("Callback Gas Limit", randomWordsRequestedEvent.CallbackGasLimit).
		Uint32("Number of Words", randomWordsRequestedEvent.NumWords).
		Uint16("Minimum Request Confirmations", randomWordsRequestedEvent.MinimumRequestConfirmations).
		Msg("RandomnessRequested Event")
}

func LogRandomWordsFulfilledEvent(
	l zerolog.Logger,
	coordinator contracts.VRFCoordinatorV2_5,
	randomWordsFulfilledEvent *vrf_coordinator_v2_5.VRFCoordinatorV25RandomWordsFulfilled,
	isNativeBilling bool,
) {
	l.Debug().
		Bool("Native Billing", isNativeBilling).
		Str("Coordinator", coordinator.Address()).
		Str("Total Payment", randomWordsFulfilledEvent.Payment.String()).
		Str("TX Hash", randomWordsFulfilledEvent.Raw.TxHash.String()).
		Str("Subscription ID", randomWordsFulfilledEvent.SubId.String()).
		Str("Request ID", randomWordsFulfilledEvent.RequestId.String()).
		Bool("Success", randomWordsFulfilledEvent.Success).
		Msg("RandomWordsFulfilled Event (TX metadata)")
}

func LogMigrationCompletedEvent(l zerolog.Logger, migrationCompletedEvent *vrf_coordinator_v2_5.VRFCoordinatorV25MigrationCompleted, vrfv2PlusContracts *VRFV2_5Contracts) {
	l.Debug().
		Str("Subscription ID", migrationCompletedEvent.SubId.String()).
		Str("Migrated From Coordinator", vrfv2PlusContracts.Coordinator.Address()).
		Str("Migrated To Coordinator", migrationCompletedEvent.NewCoordinator.String()).
		Msg("MigrationCompleted Event")
}

func LogSubDetailsAfterMigration(l zerolog.Logger, newCoordinator contracts.VRFCoordinatorV2PlusUpgradedVersion, subID *big.Int, migratedSubscription vrf_v2plus_upgraded_version.GetSubscription) {
	l.Debug().
		Str("New Coordinator", newCoordinator.Address()).
		Str("Subscription ID", subID.String()).
		Str("Juels Balance", migratedSubscription.Balance.String()).
		Str("Native Token Balance", migratedSubscription.NativeBalance.String()).
		Str("Subscription Owner", migratedSubscription.Owner.String()).
		Interface("Subscription Consumers", migratedSubscription.Consumers).
		Msg("Subscription Data After Migration to New Coordinator")
}

func LogFulfillmentDetailsLinkBilling(
	l zerolog.Logger,
	wrapperConsumerJuelsBalanceBeforeRequest *big.Int,
	wrapperConsumerJuelsBalanceAfterRequest *big.Int,
	consumerStatus vrfv2plus_wrapper_load_test_consumer.GetRequestStatus,
	randomWordsFulfilledEvent *vrf_coordinator_v2_5.VRFCoordinatorV25RandomWordsFulfilled,
) {
	l.Debug().
		Str("Consumer Balance Before Request (Link)", (*assets.Link)(wrapperConsumerJuelsBalanceBeforeRequest).Link()).
		Str("Consumer Balance After Request (Link)", (*assets.Link)(wrapperConsumerJuelsBalanceAfterRequest).Link()).
		Bool("Fulfilment Status", consumerStatus.Fulfilled).
		Str("Paid by Consumer Contract (Link)", (*assets.Link)(consumerStatus.Paid).Link()).
		Str("Paid by Coordinator Sub (Link)", (*assets.Link)(randomWordsFulfilledEvent.Payment).Link()).
		Str("RequestTimestamp", consumerStatus.RequestTimestamp.String()).
		Str("FulfilmentTimestamp", consumerStatus.FulfilmentTimestamp.String()).
		Str("RequestBlockNumber", consumerStatus.RequestBlockNumber.String()).
		Str("FulfilmentBlockNumber", consumerStatus.FulfilmentBlockNumber.String()).
		Str("TX Hash", randomWordsFulfilledEvent.Raw.TxHash.String()).
		Msg("Random Words Fulfilment Details For Link Billing")
}

func LogFulfillmentDetailsNativeBilling(
	l zerolog.Logger,
	wrapperConsumerBalanceBeforeRequestWei *big.Int,
	wrapperConsumerBalanceAfterRequestWei *big.Int,
	consumerStatus vrfv2plus_wrapper_load_test_consumer.GetRequestStatus,
	randomWordsFulfilledEvent *vrf_coordinator_v2_5.VRFCoordinatorV25RandomWordsFulfilled,
) {
	l.Debug().
		Str("Consumer Balance Before Request", assets.FormatWei(wrapperConsumerBalanceBeforeRequestWei)).
		Str("Consumer Balance After Request", assets.FormatWei(wrapperConsumerBalanceAfterRequestWei)).
		Bool("Fulfilment Status", consumerStatus.Fulfilled).
		Str("Paid by Consumer Contract", assets.FormatWei(consumerStatus.Paid)).
		Str("Paid by Coordinator Sub", assets.FormatWei(randomWordsFulfilledEvent.Payment)).
		Str("RequestTimestamp", consumerStatus.RequestTimestamp.String()).
		Str("FulfilmentTimestamp", consumerStatus.FulfilmentTimestamp.String()).
		Str("RequestBlockNumber", consumerStatus.RequestBlockNumber.String()).
		Str("FulfilmentBlockNumber", consumerStatus.FulfilmentBlockNumber.String()).
		Str("TX Hash", randomWordsFulfilledEvent.Raw.TxHash.String()).
		Msg("Random Words Request Fulfilment Details For Native Billing")
}

func logRandRequest(
	consumer string,
	coordinator string,
	subID *big.Int,
	isNativeBilling bool,
	vrfv2PlusConfig vrfv2plus_config.VRFV2PlusConfig,
	l zerolog.Logger) {
	l.Debug().
		Str("Consumer", consumer).
		Str("Coordinator", coordinator).
		Str("subID", subID.String()).
		Bool("IsNativePayment", isNativeBilling).
		Uint16("MinimumConfirmations", vrfv2PlusConfig.MinimumConfirmations).
		Uint16("RandomnessRequestCountPerRequest", vrfv2PlusConfig.RandomnessRequestCountPerRequest).
		Msg("Requesting randomness")
}
