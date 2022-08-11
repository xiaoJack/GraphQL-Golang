// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	proto "github.com/xiaoJack/GraphQL-Golang/api/proto"
)

// ReviewsServer is an autogenerated mock type for the ReviewsServer type
type ReviewsServer struct {
	mock.Mock
}

// Query provides a mock function with given fields: _a0, _a1
func (_m *ReviewsServer) Query(_a0 context.Context, _a1 *proto.QueryReviewsRequest) (*proto.QueryReviewsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.QueryReviewsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.QueryReviewsRequest) *proto.QueryReviewsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.QueryReviewsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.QueryReviewsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReviewsServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewReviewsServer creates a new instance of ReviewsServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReviewsServer(t mockConstructorTestingTNewReviewsServer) *ReviewsServer {
	mock := &ReviewsServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
