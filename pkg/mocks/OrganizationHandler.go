// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// OrganizationHandler is an autogenerated mock type for the OrganizationHandler type
type OrganizationHandler struct {
	mock.Mock
}

// NewOrganizationHandler creates a new instance of OrganizationHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrganizationHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrganizationHandler {
	mock := &OrganizationHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
