package vrfv2_5

import (
	"math/big"

	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
)

type VRFV2_5EncodedProvingKey [2]*big.Int

// VRFV2_5JobInfo defines a jobs into and proving key info
type VRFV2_5JobInfo struct {
	Job               *client.Job
	VRFKey            *client.VRFKey
	EncodedProvingKey VRFV2_5EncodedProvingKey
	KeyHash           [32]byte
}

type VRFV2_5Contracts struct {
	Coordinator      contracts.VRFCoordinatorV2_5
	BHS              contracts.BlockHashStore
	LoadTestConsumer contracts.VRFv2_5LoadTestConsumer
}
