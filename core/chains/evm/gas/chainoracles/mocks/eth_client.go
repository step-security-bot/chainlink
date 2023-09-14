// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	context "context"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"
)

// ETHClient is an autogenerated mock type for the ethClient type
type ETHClient struct {
	mock.Mock
}

// CallContract provides a mock function with given fields: ctx, msg, blockNumber
func (_m *ETHClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, msg, blockNumber)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error)); ok {
		return rf(ctx, msg, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) []byte); ok {
		r0 = rf(ctx, msg, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg, *big.Int) error); ok {
		r1 = rf(ctx, msg, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewETHClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewETHClient creates a new instance of ETHClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewETHClient(t mockConstructorTestingTNewETHClient) *ETHClient {
	mock := &ETHClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
