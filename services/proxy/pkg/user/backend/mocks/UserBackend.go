// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	mock "github.com/stretchr/testify/mock"
)

// UserBackend is an autogenerated mock type for the UserBackend type
type UserBackend struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: ctx, username, password
func (_m *UserBackend) Authenticate(ctx context.Context, username string, password string) (*userv1beta1.User, string, error) {
	ret := _m.Called(ctx, username, password)

	var r0 *userv1beta1.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *userv1beta1.User); ok {
		r0 = rf(ctx, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.User)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, string) string); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, username, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateUserFromClaims provides a mock function with given fields: ctx, claims
func (_m *UserBackend) CreateUserFromClaims(ctx context.Context, claims map[string]interface{}) (*userv1beta1.User, error) {
	ret := _m.Called(ctx, claims)

	var r0 *userv1beta1.User
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) *userv1beta1.User); ok {
		r0 = rf(ctx, claims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = rf(ctx, claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByClaims provides a mock function with given fields: ctx, claim, value
func (_m *UserBackend) GetUserByClaims(ctx context.Context, claim string, value string) (*userv1beta1.User, string, error) {
	ret := _m.Called(ctx, claim, value)

	var r0 *userv1beta1.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *userv1beta1.User); ok {
		r0 = rf(ctx, claim, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.User)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, string) string); ok {
		r1 = rf(ctx, claim, value)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, claim, value)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUserRoles provides a mock function with given fields: ctx, user
func (_m *UserBackend) GetUserRoles(ctx context.Context, user *userv1beta1.User) (*userv1beta1.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *userv1beta1.User
	if rf, ok := ret.Get(0).(func(context.Context, *userv1beta1.User) *userv1beta1.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *userv1beta1.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserBackend interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserBackend creates a new instance of UserBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserBackend(t mockConstructorTestingTNewUserBackend) *UserBackend {
	mock := &UserBackend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
