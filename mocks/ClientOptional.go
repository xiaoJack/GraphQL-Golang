// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	grpc "github.com/xiaoJack/GraphQL-Golang/internal/pkg/transports/grpc"
)

// ClientOptional is an autogenerated mock type for the ClientOptional type
type ClientOptional struct {
	mock.Mock
}

// Execute provides a mock function with given fields: o
func (_m *ClientOptional) Execute(o *grpc.ClientOptions) {
	_m.Called(o)
}

type mockConstructorTestingTNewClientOptional interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientOptional creates a new instance of ClientOptional. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientOptional(t mockConstructorTestingTNewClientOptional) *ClientOptional {
	mock := &ClientOptional{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
