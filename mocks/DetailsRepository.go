// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/xiaoJack/GraphQL-Golang/internal/pkg/models"
)

// DetailsRepository is an autogenerated mock type for the DetailsRepository type
type DetailsRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: ID
func (_m *DetailsRepository) Get(ID uint64) (*models.Detail, error) {
	ret := _m.Called(ID)

	var r0 *models.Detail
	if rf, ok := ret.Get(0).(func(uint64) *models.Detail); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Detail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDetailsRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewDetailsRepository creates a new instance of DetailsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDetailsRepository(t mockConstructorTestingTNewDetailsRepository) *DetailsRepository {
	mock := &DetailsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}