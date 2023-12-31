// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	menu "try-go-grpc/menu/proto"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	mock "github.com/stretchr/testify/mock"
)

// MenuServer is an autogenerated mock type for the MenuServer type
type MenuServer struct {
	mock.Mock
}

// ListBeers provides a mock function with given fields: ctx, request
func (_m *MenuServer) ListBeers(ctx context.Context, request *emptypb.Empty) (*menu.ListBeersResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *menu.ListBeersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) (*menu.ListBeersResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *menu.ListBeersResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*menu.ListBeersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMenuServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMenuServer creates a new instance of MenuServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMenuServer(t mockConstructorTestingTNewMenuServer) *MenuServer {
	mock := &MenuServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
