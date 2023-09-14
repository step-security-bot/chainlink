package vrfv2_5

import (
	"context"
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/actions/vrfv2_5/vrfv2_5_constants"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/docker/test_env"
	"github.com/smartcontractkit/chainlink/integration-tests/types/config/node"
	chainlinkutils "github.com/smartcontractkit/chainlink/v2/core/utils"
)

var (
	ErrNodePrimaryKey         = "error getting node's primary ETH key"
	ErrCreatingProvingKeyHash = "error creating a keyHash from the proving key"
	ErrRegisteringProvingKey  = "error registering a proving key on Coordinator contract"
	ErrRegisterProvingKey     = "error registering proving keys"
	ErrEncodingProvingKey     = "error encoding proving key"
	ErrCreatingVRFv2_5Key     = "error creating VRFv2_5 key"
	ErrDeployBlockHashStore   = "error deploying blockhash store"
	ErrDeployCoordinator      = "error deploying VRF CoordinatorV2_5"
	ErrAdvancedConsumer       = "error deploying VRFv2_5 Advanced Consumer"
	ErrABIEncodingFunding     = "error Abi encoding subscriptionID"
	ErrSendingLinkToken       = "error sending Link token"
	ErrCreatingVRFv2_5Job     = "error creating VRFv2_5 job"
	ErrParseJob               = "error parsing job definition"

	ErrDeployVRFV2_5Contracts  = "error deploying VRFV2_5 contracts"
	ErrSetVRFCoordinatorConfig = "error setting config for VRF Coordinator contract"
	ErrCreateVRFSubscription   = "error creating VRF Subscription"
	ErrFindSubID               = "error finding created subscription ID"
	ErrAddConsumerToSub        = "error adding consumer to VRF Subscription"
	ErrFundSubWithNativeToken  = "error funding subscription with native token"
	ErrSetLinkETHNativeFeed    = "error setting Link and ETH/LINK feed for VRF Coordinator contract"
	ErrFundSubWithLinkToken    = "error funding subscription with Link tokens"
	ErrCreateVRFV2_5Jobs       = "error creating VRF V2 Plus Jobs"
	ErrGetPrimaryKey           = "error getting primary ETH key address"
	ErrRestartCLNode           = "error restarting CL node"
	ErrWaitTXsComplete         = "error waiting for TXs to complete"
)

func DeployVRFV2_5Contracts(
	contractDeployer contracts.ContractDeployer,
	chainClient blockchain.EVMClient,
) (*VRFV2_5Contracts, error) {
	bhs, err := contractDeployer.DeployBlockhashStore()
	if err != nil {
		return nil, errors.Wrap(err, ErrDeployBlockHashStore)
	}
	coordinator, err := contractDeployer.DeployVRFCoordinatorV2_5(bhs.Address())
	if err != nil {
		return nil, errors.Wrap(err, ErrDeployCoordinator)
	}
	loadTestConsumer, err := contractDeployer.DeployVRFv2_5LoadTestConsumer(coordinator.Address())
	if err != nil {
		return nil, errors.Wrap(err, ErrAdvancedConsumer)
	}
	err = chainClient.WaitForEvents()
	if err != nil {
		return nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	return &VRFV2_5Contracts{coordinator, bhs, loadTestConsumer}, nil
}

func CreateVRFV2_5Jobs(
	chainlinkNodes []*client.ChainlinkClient,
	coordinator contracts.VRFCoordinatorV2_5,
	c blockchain.EVMClient,
	minIncomingConfirmations uint16,
) ([]*VRFV2_5JobInfo, error) {
	jobInfo := make([]*VRFV2_5JobInfo, 0)
	for _, chainlinkNode := range chainlinkNodes {
		vrfKey, err := chainlinkNode.MustCreateVRFKey()
		if err != nil {
			return nil, errors.Wrap(err, ErrCreatingVRFv2_5Key)
		}
		pubKeyCompressed := vrfKey.Data.ID
		jobUUID := uuid.New()
		os := &client.VRFV2_5TxPipelineSpec{
			Address: coordinator.Address(),
		}
		ost, err := os.String()
		if err != nil {
			return nil, errors.Wrap(err, ErrParseJob)
		}
		nativeTokenPrimaryKeyAddress, err := chainlinkNode.PrimaryEthAddress()
		if err != nil {
			return nil, errors.Wrap(err, ErrNodePrimaryKey)
		}
		job, err := chainlinkNode.MustCreateJob(&client.VRFV2_5JobSpec{
			Name:                     fmt.Sprintf("vrf-v2-plus-%s", jobUUID),
			CoordinatorAddress:       coordinator.Address(),
			FromAddresses:            []string{nativeTokenPrimaryKeyAddress},
			EVMChainID:               c.GetChainID().String(),
			MinIncomingConfirmations: int(minIncomingConfirmations),
			PublicKey:                pubKeyCompressed,
			ExternalJobID:            jobUUID.String(),
			ObservationSource:        ost,
			BatchFulfillmentEnabled:  false,
		})
		if err != nil {
			return nil, errors.Wrap(err, ErrCreatingVRFv2_5Job)
		}
		provingKey, err := VRFV2RegisterProvingKey(vrfKey, nativeTokenPrimaryKeyAddress, coordinator)
		if err != nil {
			return nil, errors.Wrap(err, ErrRegisteringProvingKey)
		}
		keyHash, err := coordinator.HashOfKey(context.Background(), provingKey)
		if err != nil {
			return nil, errors.Wrap(err, ErrCreatingProvingKeyHash)
		}
		ji := &VRFV2_5JobInfo{
			Job:               job,
			VRFKey:            vrfKey,
			EncodedProvingKey: provingKey,
			KeyHash:           keyHash,
		}
		jobInfo = append(jobInfo, ji)
	}
	return jobInfo, nil
}

func VRFV2RegisterProvingKey(
	vrfKey *client.VRFKey,
	oracleAddress string,
	coordinator contracts.VRFCoordinatorV2_5,
) (VRFV2_5EncodedProvingKey, error) {
	provingKey, err := actions.EncodeOnChainVRFProvingKey(*vrfKey)
	if err != nil {
		return VRFV2_5EncodedProvingKey{}, errors.Wrap(err, ErrEncodingProvingKey)
	}
	err = coordinator.RegisterProvingKey(
		oracleAddress,
		provingKey,
	)
	if err != nil {
		return VRFV2_5EncodedProvingKey{}, errors.Wrap(err, ErrRegisterProvingKey)
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
	linkAddress contracts.LinkToken,
	mockETHNativeFeedAddress contracts.MockETHLINKFeed,
) (*test_env.CLClusterTestEnv, *VRFV2_5Contracts, *big.Int, *VRFV2_5JobInfo, error) {

	vrfv2_5Contracts, err := DeployVRFV2_5Contracts(env.ContractDeployer, env.EVMClient)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrDeployVRFV2_5Contracts)
	}

	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}

	err = vrfv2_5Contracts.Coordinator.SetConfig(
		vrfv2_5_constants.MinimumConfirmations,
		vrfv2_5_constants.MaxGasLimitVRFCoordinatorConfig,
		vrfv2_5_constants.StalenessSeconds,
		vrfv2_5_constants.GasAfterPaymentCalculation,
		vrfv2_5_constants.LinkEthFeedResponse,
		vrfv2_5_constants.VRFCoordinatorV2_5FeeConfig,
	)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrSetVRFCoordinatorConfig)
	}
	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}

	err = vrfv2_5Contracts.Coordinator.CreateSubscription()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrCreateVRFSubscription)
	}
	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}

	subID, err := vrfv2_5Contracts.Coordinator.FindSubscriptionID()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrFindSubID)
	}

	err = vrfv2_5Contracts.Coordinator.AddConsumer(subID, vrfv2_5Contracts.LoadTestConsumer.Address())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrAddConsumerToSub)
	}

	//Native Billing
	err = vrfv2_5Contracts.Coordinator.FundSubscriptionWithNative(subID, big.NewInt(0).Mul(vrfv2_5_constants.VRFSubscriptionFundingAmountNativeToken, big.NewInt(1e18)))
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrFundSubWithNativeToken)
	}

	//Link Billing
	err = vrfv2_5Contracts.Coordinator.SetLINKAndLINKNativeFeed(linkAddress.Address(), mockETHNativeFeedAddress.Address())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrSetLinkETHNativeFeed)
	}
	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}
	err = FundVRFCoordinatorV2_5Subscription(linkAddress, vrfv2_5Contracts.Coordinator, env.EVMClient, subID, vrfv2_5_constants.VRFSubscriptionFundingAmountLink)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrFundSubWithLinkToken)
	}
	err = env.EVMClient.WaitForEvents()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrWaitTXsComplete)
	}

	vrfV2_5Jobs, err := CreateVRFV2_5Jobs(env.GetAPIs(), vrfv2_5Contracts.Coordinator, env.EVMClient, vrfv2_5_constants.MinimumConfirmations)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrCreateVRFV2_5Jobs)
	}

	// this part is here because VRFv2 can work with only a specific key
	// [[EVM.KeySpecific]]
	//	Key = '...'
	addr, err := env.CLNodes[0].API.PrimaryEthAddress()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrGetPrimaryKey)
	}
	nodeConfig := node.NewConfig(env.CLNodes[0].NodeConfig,
		node.WithVRFv2EVMEstimator(addr),
	)
	err = env.CLNodes[0].Restart(nodeConfig)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, ErrRestartCLNode)
	}
	return env, vrfv2_5Contracts, subID, vrfV2_5Jobs[0], nil
}
