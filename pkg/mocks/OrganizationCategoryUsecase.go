// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	models "server/internal/core/models"

	mock "github.com/stretchr/testify/mock"
)

// OrganizationCategoryUsecase is an autogenerated mock type for the OrganizationCategoryUsecase type
type OrganizationCategoryUsecase struct {
	mock.Mock
}

// CreateOrganizationCategory provides a mock function with given fields: organizationCategory, requesterUserId
func (_m *OrganizationCategoryUsecase) CreateOrganizationCategory(organizationCategory *models.OrganizationCategory, requesterUserId int) (*models.OrganizationCategory, error) {
	ret := _m.Called(organizationCategory, requesterUserId)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrganizationCategory")
	}

	var r0 *models.OrganizationCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.OrganizationCategory, int) (*models.OrganizationCategory, error)); ok {
		return rf(organizationCategory, requesterUserId)
	}
	if rf, ok := ret.Get(0).(func(*models.OrganizationCategory, int) *models.OrganizationCategory); ok {
		r0 = rf(organizationCategory, requesterUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OrganizationCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.OrganizationCategory, int) error); ok {
		r1 = rf(organizationCategory, requesterUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOrganizationCategory provides a mock function with given fields: organizationCategory, requesterUserId
func (_m *OrganizationCategoryUsecase) DeleteOrganizationCategory(organizationCategory *models.OrganizationCategory, requesterUserId int) error {
	ret := _m.Called(organizationCategory, requesterUserId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOrganizationCategory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.OrganizationCategory, int) error); ok {
		r0 = rf(organizationCategory, requesterUserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrganizationCategory provides a mock function with given fields: organizationCategory, requesterUserId
func (_m *OrganizationCategoryUsecase) UpdateOrganizationCategory(organizationCategory *models.OrganizationCategory, requesterUserId int) (*models.OrganizationCategory, error) {
	ret := _m.Called(organizationCategory, requesterUserId)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrganizationCategory")
	}

	var r0 *models.OrganizationCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.OrganizationCategory, int) (*models.OrganizationCategory, error)); ok {
		return rf(organizationCategory, requesterUserId)
	}
	if rf, ok := ret.Get(0).(func(*models.OrganizationCategory, int) *models.OrganizationCategory); ok {
		r0 = rf(organizationCategory, requesterUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OrganizationCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.OrganizationCategory, int) error); ok {
		r1 = rf(organizationCategory, requesterUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrganizationCategoryUsecase creates a new instance of OrganizationCategoryUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrganizationCategoryUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrganizationCategoryUsecase {
	mock := &OrganizationCategoryUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
