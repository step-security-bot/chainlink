package offchainreporting

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	"gorm.io/gorm"
)

type txManager interface {
	CreateEthTransaction(tx postgres.Queryer, fromAddress, toAddress common.Address, payload []byte, gasLimit uint64, meta interface{}, strategy bulletprooftxmanager.TxStrategy) (etx bulletprooftxmanager.EthTx, err error)
}

type transmitter struct {
	txm         txManager
	db          *gorm.DB
	fromAddress common.Address
	gasLimit    uint64
	strategy    bulletprooftxmanager.TxStrategy
}

// NewTransmitter creates a new eth transmitter
func NewTransmitter(txm txManager, db *gorm.DB, fromAddress common.Address, gasLimit uint64, strategy bulletprooftxmanager.TxStrategy) Transmitter {
	return &transmitter{
		txm:         txm,
		db:          db,
		fromAddress: fromAddress,
		gasLimit:    gasLimit,
		strategy:    strategy,
	}
}

func (t *transmitter) CreateEthTransaction(ctx context.Context, toAddress common.Address, payload []byte) error {
	db, err := t.db.DB()
	if err != nil {
		return err
	}
	return postgres.SqlTransaction(ctx, db, func(tx *sqlx.Tx) error {
		_, err := t.txm.CreateEthTransaction(tx, t.fromAddress, toAddress, payload, t.gasLimit, nil, t.strategy)
		return errors.Wrap(err, "Skipped OCR transmission")
	})
}

func (t *transmitter) FromAddress() common.Address {
	return t.fromAddress
}
