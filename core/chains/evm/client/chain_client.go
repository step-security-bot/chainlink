package client

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	commonclient "github.com/smartcontractkit/chainlink/v2/common/chains/client"
	"github.com/smartcontractkit/chainlink/v2/core/assets"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/config"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

var _ Client = (*chainClient)(nil)

// TODO-1663: rename this to client, once the client.go file is deprecated.
type chainClient struct {
	evmClient commonclient.Client[
		*big.Int,
		evmtypes.Nonce,
		common.Address,
		common.Hash,
		*types.Transaction,
		common.Hash,
		types.Log,
		ethereum.FilterQuery,
		*evmtypes.Receipt,
		*assets.Wei,
		*evmtypes.Head,
		*rpcClient,
	]
	logger logger.Logger
}

func NewChainClient(
	logger logger.Logger,
	selectionMode string,
	noNewHeadsThreshold time.Duration,
	nodes []commonclient.Node[*big.Int, *evmtypes.Head, *rpcClient],
	sendonlys []commonclient.SendOnlyNode[*big.Int, *rpcClient],
	chainID *big.Int,
	chainType config.ChainType,
) Client {
	evmClient := commonclient.NewClient[
		*big.Int,
		evmtypes.Nonce,
		common.Address,
		common.Hash,
		*types.Transaction,
		common.Hash,
		types.Log,
		ethereum.FilterQuery,
		*evmtypes.Receipt,
		*assets.Wei,
		*evmtypes.Head,
		*rpcClient,
	](
		logger,
		selectionMode,
		noNewHeadsThreshold,
		nodes,
		sendonlys,
		chainID,
		chainType,
		NewSendOnlyErrorReturnCode,
	)
	return &chainClient{
		evmClient: evmClient,
		logger:    logger,
	}
}

func (c *chainClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return c.evmClient.BalanceAt(ctx, account, blockNumber)
}

func (c *chainClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	batch := make([]any, len(b))
	for i, arg := range b {
		batch[i] = any(arg)
	}
	return c.evmClient.BatchCallContext(ctx, batch)
}

func (c *chainClient) BatchCallContextAll(ctx context.Context, b []rpc.BatchElem) error {
	batch := make([]any, len(b))
	for i, arg := range b {
		batch[i] = any(arg)
	}
	return c.evmClient.BatchCallContextAll(ctx, batch)
}

// TODO-1663: return custom Block type instead of geth's once client.go is deprecated.
func (c *chainClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return c.evmClient.NodeRPC().BlockByHashGeth(ctx, hash)
}

// TODO-1663: return custom Block type instead of geth's once client.go is deprecated.
func (c *chainClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return c.evmClient.NodeRPC().BlockByNumberGeth(ctx, number)
}

func (c *chainClient) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return c.evmClient.CallContext(ctx, result, method)
}

func (c *chainClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return c.evmClient.CallContract(ctx, msg, blockNumber)
}

// TODO-1663: change this to actual ChainID() call once client.go is deprecated.
func (c *chainClient) ChainID() (*big.Int, error) {
	//return c.evmClient.ChainID(ctx), nil
	return c.evmClient.ConfiguredChainID(), nil
}

func (c *chainClient) Close() {
	c.evmClient.Close()
}

func (c *chainClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return c.evmClient.CodeAt(ctx, account, blockNumber)
}

func (c *chainClient) ConfiguredChainID() *big.Int {
	return c.evmClient.ConfiguredChainID()
}

func (c *chainClient) Dial(ctx context.Context) error {
	return c.evmClient.Dial(ctx)
}

func (c *chainClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return c.evmClient.EstimateGas(ctx, call)
}
func (c *chainClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return c.evmClient.FilterEvents(ctx, q)
}

func (c *chainClient) HeaderByHash(ctx context.Context, h common.Hash) (*types.Header, error) {
	return c.evmClient.NodeRPC().HeaderByHash(ctx, h)
}

func (c *chainClient) HeaderByNumber(ctx context.Context, n *big.Int) (*types.Header, error) {
	return c.evmClient.NodeRPC().HeaderByNumber(ctx, n)
}

func (c *chainClient) HeadByHash(ctx context.Context, h common.Hash) (*evmtypes.Head, error) {
	return c.evmClient.BlockByHash(ctx, h)
}

func (c *chainClient) HeadByNumber(ctx context.Context, n *big.Int) (*evmtypes.Head, error) {
	return c.evmClient.BlockByNumber(ctx, n)
}

func (c *chainClient) IsL2() bool {
	return c.evmClient.IsL2()
}

func (c *chainClient) LINKBalance(ctx context.Context, address common.Address, linkAddress common.Address) (*assets.Link, error) {
	return c.evmClient.LINKBalance(ctx, address, linkAddress)
}

func (c *chainClient) LatestBlockHeight(ctx context.Context) (*big.Int, error) {
	return c.evmClient.LatestBlockHeight(ctx)
}

func (c *chainClient) NodeStates() map[string]string {
	return c.evmClient.NodeStates()
}

func (c *chainClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return c.evmClient.NodeRPC().PendingCodeAt(ctx, account)
}

// TODO-1663: change this to evmtypes.Nonce(int64) once client.go is deprecated.
func (c *chainClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	n, err := c.evmClient.PendingSequenceAt(ctx, account)
	return uint64(n), err
}

func (c *chainClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.evmClient.SendTransaction(ctx, tx)
}

func (c *chainClient) SendTransactionReturnCode(ctx context.Context, tx *types.Transaction, fromAddress common.Address) (commonclient.SendTxReturnCode, error) {
	err := c.SendTransaction(ctx, tx)
	return NewSendErrorReturnCode(err, c.logger, tx, fromAddress, c.IsL2())
}

func (c *chainClient) SequenceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (evmtypes.Nonce, error) {
	return c.evmClient.SequenceAt(ctx, account, blockNumber)
}

func (c *chainClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return c.evmClient.NodeRPC().SubscribeFilterLogs(ctx, q, ch)
}

func (c *chainClient) SubscribeNewHead(ctx context.Context, ch chan<- *evmtypes.Head) (ethereum.Subscription, error) {
	return c.evmClient.Subscribe(ctx, ch)
}

func (c *chainClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.evmClient.NodeRPC().SuggestGasPrice(ctx)
}

func (c *chainClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return c.evmClient.NodeRPC().SuggestGasTipCap(ctx)
}

func (c *chainClient) TokenBalance(ctx context.Context, address common.Address, contractAddress common.Address) (*big.Int, error) {
	return c.evmClient.TokenBalance(ctx, address, contractAddress)
}

func (c *chainClient) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, error) {
	return c.evmClient.TransactionByHash(ctx, txHash)
}

// TODO-1663: return custom Receipt type instead of geth's once client.go is deprecated.
func (c *chainClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	//return c.evmClient.TransactionReceipt(ctx, txHash)
	return c.evmClient.NodeRPC().TransactionReceiptGeth(ctx, txHash)
}
