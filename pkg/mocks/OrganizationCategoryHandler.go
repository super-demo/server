// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// OrganizationCategoryHandler is an autogenerated mock type for the OrganizationCategoryHandler type
type OrganizationCategoryHandler struct {
	mock.Mock
}

// NewOrganizationCategoryHandler creates a new instance of OrganizationCategoryHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrganizationCategoryHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrganizationCategoryHandler {
	mock := &OrganizationCategoryHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
