// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// OrganizationServiceUsecase is an autogenerated mock type for the OrganizationServiceUsecase type
type OrganizationServiceUsecase struct {
	mock.Mock
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
