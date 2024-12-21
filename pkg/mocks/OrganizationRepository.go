// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	models "server/internal/core/models"

	mock "github.com/stretchr/testify/mock"
)

// OrganizationRepository is an autogenerated mock type for the OrganizationRepository type
type OrganizationRepository struct {
	mock.Mock
}

// CreateOrganization provides a mock function with given fields: organization
func (_m *OrganizationRepository) CreateOrganization(organization *models.Organization) (*models.Organization, error) {
	ret := _m.Called(organization)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrganization")
	}

	var r0 *models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Organization) (*models.Organization, error)); ok {
		return rf(organization)
	}
	if rf, ok := ret.Get(0).(func(*models.Organization) *models.Organization); ok {
		r0 = rf(organization)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Organization) error); ok {
		r1 = rf(organization)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizationById provides a mock function with given fields: id
func (_m *OrganizationRepository) GetOrganizationById(id int) (*models.Organization, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetOrganizationById")
	}

	var r0 *models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Organization, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Organization); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizationListByUserId provides a mock function with given fields: id
func (_m *OrganizationRepository) GetOrganizationListByUserId(id int) (*[]models.Organization, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetOrganizationListByUserId")
	}

	var r0 *[]models.Organization
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*[]models.Organization, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *[]models.Organization); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Organization)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrganizationRepository creates a new instance of OrganizationRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrganizationRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrganizationRepository {
	mock := &OrganizationRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
