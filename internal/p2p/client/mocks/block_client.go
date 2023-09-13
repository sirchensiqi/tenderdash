// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	blocksync "github.com/dashpay/tenderdash/proto/tendermint/blocksync"

	context "context"

	mock "github.com/stretchr/testify/mock"

	promise "github.com/dashpay/tenderdash/libs/promise"

	types "github.com/dashpay/tenderdash/types"
)

// BlockClient is an autogenerated mock type for the BlockClient type
type BlockClient struct {
	mock.Mock
}

// GetBlock provides a mock function with given fields: ctx, height, peerID
func (_m *BlockClient) GetBlock(ctx context.Context, height int64, peerID types.NodeID) (*promise.Promise[*blocksync.BlockResponse], error) {
	ret := _m.Called(ctx, height, peerID)

	var r0 *promise.Promise[*blocksync.BlockResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, types.NodeID) (*promise.Promise[*blocksync.BlockResponse], error)); ok {
		return rf(ctx, height, peerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, types.NodeID) *promise.Promise[*blocksync.BlockResponse]); ok {
		r0 = rf(ctx, height, peerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise[*blocksync.BlockResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, types.NodeID) error); ok {
		r1 = rf(ctx, height, peerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSyncStatus provides a mock function with given fields: ctx
func (_m *BlockClient) GetSyncStatus(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: ctx, msg
func (_m *BlockClient) Send(ctx context.Context, msg interface{}) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBlockClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlockClient creates a new instance of BlockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlockClient(t mockConstructorTestingTNewBlockClient) *BlockClient {
	mock := &BlockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
