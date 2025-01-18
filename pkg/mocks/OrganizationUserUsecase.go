// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	models "server/internal/core/models"

	mock "github.com/stretchr/testify/mock"
)

// OrganizationUserUsecase is an autogenerated mock type for the OrganizationUserUsecase type
type OrganizationUserUsecase struct {
	mock.Mock
}

// CreateOrganizationUser provides a mock function with given fields: organizationUser, requesterUserId
func (_m *OrganizationUserUsecase) CreateOrganizationUser(organizationUser *models.OrganizationUser, requesterUserId int) (*models.OrganizationUser, error) {
	ret := _m.Called(organizationUser, requesterUserId)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrganizationUser")
	}

	var r0 *models.OrganizationUser
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.OrganizationUser, int) (*models.OrganizationUser, error)); ok {
		return rf(organizationUser, requesterUserId)
	}
	if rf, ok := ret.Get(0).(func(*models.OrganizationUser, int) *models.OrganizationUser); ok {
		r0 = rf(organizationUser, requesterUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OrganizationUser)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.OrganizationUser, int) error); ok {
		r1 = rf(organizationUser, requesterUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOrganizationUser provides a mock function with given fields: organizationUser, requesterUserId
func (_m *OrganizationUserUsecase) DeleteOrganizationUser(organizationUser *models.OrganizationUser, requesterUserId int) error {
	ret := _m.Called(organizationUser, requesterUserId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOrganizationUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.OrganizationUser, int) error); ok {
		r0 = rf(organizationUser, requesterUserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrganizationUserUsecase creates a new instance of OrganizationUserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrganizationUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrganizationUserUsecase {
	mock := &OrganizationUserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
