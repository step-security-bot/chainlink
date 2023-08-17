package evm

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ocr2keepers "github.com/smartcontractkit/ocr2keepers/pkg/v3/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ocr2keeper/evm21/core"
)

var (
	upkeepId1 = big.NewInt(100)
	upkeepId2 = big.NewInt(200)
	trigger1  = ocr2keepers.Trigger{
		BlockNumber: block1,
		BlockHash:   common.HexToHash("0x1"),
		LogTriggerExtension: &ocr2keepers.LogTriggerExtension{
			Index:  1,
			TxHash: common.HexToHash("0x1"),
		},
	}
	trigger2 = ocr2keepers.Trigger{
		BlockNumber: block3,
		BlockHash:   common.HexToHash("0x1"),
		LogTriggerExtension: &ocr2keepers.LogTriggerExtension{
			Index:  2, // a different BlockNumber isn't enough to generate a unique work ID, so we change the index here to generate a separate work ID
			TxHash: common.HexToHash("0x1"),
		},
	}
	payload2, _ = core.NewUpkeepPayload(upkeepId1, trigger1, []byte{})
	payload3, _ = core.NewUpkeepPayload(upkeepId2, trigger1, []byte{})
	payload4, _ = core.NewUpkeepPayload(upkeepId1, trigger2, []byte{})
	payload5, _ = core.NewUpkeepPayload(upkeepId1, trigger1, []byte{})
)

const (
	block1 = 111
	block3 = 113
)

func TestUpkeepStateStore_OverrideUpkeepStates(t *testing.T) {
	p := ocr2keepers.Performed
	e := ocr2keepers.Ineligible

	tests := []struct {
		name          string
		payloads      []ocr2keepers.UpkeepPayload
		states        []ocr2keepers.UpkeepState
		expectedError error
		oldIds        []string
		oldIdResult   []upkeepState
		newIds        []string
		newIdResult   []upkeepState
		upkeepIds     []*big.Int
		endBlock      int64
		startBlock    int64
		result        []upkeepState
	}{
		{
			name: "overrides existing upkeep states",
			payloads: []ocr2keepers.UpkeepPayload{
				payload2,
				payload3,
				payload4,
				payload5, // this overrides payload 2 bc they have the same payload ID
			},
			states: []ocr2keepers.UpkeepState{ocr2keepers.Performed, ocr2keepers.Performed, ocr2keepers.Performed, ocr2keepers.Ineligible},
			oldIds: []string{payload2.WorkID, payload3.WorkID, payload4.WorkID},
			oldIdResult: []upkeepState{
				{
					payload: &payload3,
					state:   &p,
				},
				{
					payload: &payload4,
					state:   &p,
				},
			},
			newIds: []string{payload3.WorkID, payload4.WorkID, payload5.WorkID},
			newIdResult: []upkeepState{
				{
					payload: &payload3,
					state:   &p,
				},
				{
					payload: &payload4,
					state:   &p,
				},
				{
					payload: &payload5,
					state:   &e,
				},
			},

			upkeepIds:  []*big.Int{upkeepId1},
			endBlock:   block3 + 1,
			startBlock: block1,
			result: []upkeepState{
				{
					payload: &payload5,
					state:   &e,
				},
				{
					payload: &payload4,
					state:   &p,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewUpkeepStateStore(logger.TestLogger(t))
			for i, p := range tc.payloads {
				err := store.SetUpkeepState(p, tc.states[i])
				require.Equal(t, err, tc.expectedError)
			}

			pl, us, err := store.SelectByUpkeepIDsAndBlockRange(tc.upkeepIds, tc.startBlock, tc.endBlock)
			require.Nil(t, err)
			require.Equal(t, len(tc.result), len(pl))
			require.Equal(t, len(tc.result), len(us))
			for j, r := range tc.result {
				require.Equal(t, r.payload, pl[j])
				require.Equal(t, r.state, us[j])
			}

		})
	}
}