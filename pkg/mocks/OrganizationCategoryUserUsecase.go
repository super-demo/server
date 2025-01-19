// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	models "server/internal/core/models"

	mock "github.com/stretchr/testify/mock"
)

// OrganizationCategoryUserUsecase is an autogenerated mock type for the OrganizationCategoryUserUsecase type
type OrganizationCategoryUserUsecase struct {
	mock.Mock
}

// CreateOrganizationCategoryUser provides a mock function with given fields: organizationCategoryUser, requesterUserId
func (_m *OrganizationCategoryUserUsecase) CreateOrganizationCategoryUser(organizationCategoryUser *models.OrganizationCategoryUser, requesterUserId int) (*models.OrganizationCategoryUser, error) {
	ret := _m.Called(organizationCategoryUser, requesterUserId)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrganizationCategoryUser")
	}

	var r0 *models.OrganizationCategoryUser
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.OrganizationCategoryUser, int) (*models.OrganizationCategoryUser, error)); ok {
		return rf(organizationCategoryUser, requesterUserId)
	}
	if rf, ok := ret.Get(0).(func(*models.OrganizationCategoryUser, int) *models.OrganizationCategoryUser); ok {
		r0 = rf(organizationCategoryUser, requesterUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OrganizationCategoryUser)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.OrganizationCategoryUser, int) error); ok {
		r1 = rf(organizationCategoryUser, requesterUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOrganizationCategoryUser provides a mock function with given fields: organizationCategoryUser, requesterUserId
func (_m *OrganizationCategoryUserUsecase) DeleteOrganizationCategoryUser(organizationCategoryUser *models.OrganizationCategoryUser, requesterUserId int) error {
	ret := _m.Called(organizationCategoryUser, requesterUserId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOrganizationCategoryUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.OrganizationCategoryUser, int) error); ok {
		r0 = rf(organizationCategoryUser, requesterUserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrganizationCategoryUserUsecase creates a new instance of OrganizationCategoryUserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrganizationCategoryUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrganizationCategoryUserUsecase {
	mock := &OrganizationCategoryUserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
