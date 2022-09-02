// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// InitControllers is an autogenerated mock type for the InitControllers type
type InitControllers struct {
	mock.Mock
}

// Execute provides a mock function with given fields: r
func (_m *InitControllers) Execute(r *gin.Engine) {
	_m.Called(r)
}

type mockConstructorTestingTNewInitControllers interface {
	mock.TestingT
	Cleanup(func())
}

// NewInitControllers creates a new instance of InitControllers. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInitControllers(t mockConstructorTestingTNewInitControllers) *InitControllers {
	mock := &InitControllers{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
