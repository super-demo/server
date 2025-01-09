// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	models "server/internal/core/models"

	mock "github.com/stretchr/testify/mock"
)

// AuthenticationRepository is an autogenerated mock type for the AuthenticationRepository type
type AuthenticationRepository struct {
	mock.Mock
}

// GetUserInfoByAccessToken provides a mock function with given fields: token
func (_m *AuthenticationRepository) GetUserInfoByAccessToken(token string) (*models.UserInfoResponse, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for GetUserInfoByAccessToken")
	}

	var r0 *models.UserInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.UserInfoResponse, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *models.UserInfoResponse); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserInfoResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthenticationRepository creates a new instance of AuthenticationRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthenticationRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthenticationRepository {
	mock := &AuthenticationRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
