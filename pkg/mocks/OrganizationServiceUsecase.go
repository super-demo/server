// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	models "server/internal/core/models"

	mock "github.com/stretchr/testify/mock"
)

// OrganizationServiceUsecase is an autogenerated mock type for the OrganizationServiceUsecase type
type OrganizationServiceUsecase struct {
	mock.Mock
}

// CreateOrganizationService provides a mock function with given fields: organizationService, requesterUserId
func (_m *OrganizationServiceUsecase) CreateOrganizationService(organizationService *models.OrganizationService, requesterUserId int) (*models.OrganizationService, error) {
	ret := _m.Called(organizationService, requesterUserId)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrganizationService")
	}

	var r0 *models.OrganizationService
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.OrganizationService, int) (*models.OrganizationService, error)); ok {
		return rf(organizationService, requesterUserId)
	}
	if rf, ok := ret.Get(0).(func(*models.OrganizationService, int) *models.OrganizationService); ok {
		r0 = rf(organizationService, requesterUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OrganizationService)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.OrganizationService, int) error); ok {
		r1 = rf(organizationService, requesterUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrganizationServiceUsecase creates a new instance of OrganizationServiceUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrganizationServiceUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrganizationServiceUsecase {
	mock := &OrganizationServiceUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
