// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/xiaoJack/GraphQL-Golang/internal/pkg/models"
)

// ReviewsRepository is an autogenerated mock type for the ReviewsRepository type
type ReviewsRepository struct {
	mock.Mock
}

// Query provides a mock function with given fields: productID
func (_m *ReviewsRepository) Query(productID uint64) ([]*models.Review, error) {
	ret := _m.Called(productID)

	var r0 []*models.Review
	if rf, ok := ret.Get(0).(func(uint64) []*models.Review); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Review)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReviewsRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewReviewsRepository creates a new instance of ReviewsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReviewsRepository(t mockConstructorTestingTNewReviewsRepository) *ReviewsRepository {
	mock := &ReviewsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
