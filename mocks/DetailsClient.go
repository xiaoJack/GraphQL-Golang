// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	proto "github.com/xiaoJack/GraphQL-Golang/api/proto"
)

// DetailsClient is an autogenerated mock type for the DetailsClient type
type DetailsClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, in, opts
func (_m *DetailsClient) Get(ctx context.Context, in *proto.GetDetailRequest, opts ...grpc.CallOption) (*proto.Detail, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.Detail
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetDetailRequest, ...grpc.CallOption) *proto.Detail); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.Detail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetDetailRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDetailsClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewDetailsClient creates a new instance of DetailsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDetailsClient(t mockConstructorTestingTNewDetailsClient) *DetailsClient {
	mock := &DetailsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}