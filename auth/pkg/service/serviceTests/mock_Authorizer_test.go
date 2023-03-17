// Code generated by mockery v2.22.1. DO NOT EDIT.

package serviceTests

import (
	entities "auth/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockAuthorizer is an autogenerated mock type for the Authorizer type
type MockAuthorizer struct {
	mock.Mock
}

type MockAuthorizer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthorizer) EXPECT() *MockAuthorizer_Expecter {
	return &MockAuthorizer_Expecter{mock: &_m.Mock}
}

// CheckToken provides a mock function with given fields: authToken, signKey
func (_m *MockAuthorizer) CheckToken(authToken string, signKey string) (string, error) {
	ret := _m.Called(authToken, signKey)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(authToken, signKey)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(authToken, signKey)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(authToken, signKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthorizer_CheckToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckToken'
type MockAuthorizer_CheckToken_Call struct {
	*mock.Call
}

// CheckToken is a helper method to define mock.On call
//   - authToken string
//   - signKey string
func (_e *MockAuthorizer_Expecter) CheckToken(authToken interface{}, signKey interface{}) *MockAuthorizer_CheckToken_Call {
	return &MockAuthorizer_CheckToken_Call{Call: _e.mock.On("CheckToken", authToken, signKey)}
}

func (_c *MockAuthorizer_CheckToken_Call) Run(run func(authToken string, signKey string)) *MockAuthorizer_CheckToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockAuthorizer_CheckToken_Call) Return(_a0 string, _a1 error) *MockAuthorizer_CheckToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthorizer_CheckToken_Call) RunAndReturn(run func(string, string) (string, error)) *MockAuthorizer_CheckToken_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: u
func (_m *MockAuthorizer) CreateUser(u entities.User) (entities.AuthenticatedUserResponse, error) {
	ret := _m.Called(u)

	var r0 entities.AuthenticatedUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.User) (entities.AuthenticatedUserResponse, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(entities.User) entities.AuthenticatedUserResponse); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(entities.AuthenticatedUserResponse)
	}

	if rf, ok := ret.Get(1).(func(entities.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthorizer_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockAuthorizer_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - u entities.User
func (_e *MockAuthorizer_Expecter) CreateUser(u interface{}) *MockAuthorizer_CreateUser_Call {
	return &MockAuthorizer_CreateUser_Call{Call: _e.mock.On("CreateUser", u)}
}

func (_c *MockAuthorizer_CreateUser_Call) Run(run func(u entities.User)) *MockAuthorizer_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.User))
	})
	return _c
}

func (_c *MockAuthorizer_CreateUser_Call) Return(_a0 entities.AuthenticatedUserResponse, _a1 error) *MockAuthorizer_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthorizer_CreateUser_Call) RunAndReturn(run func(entities.User) (entities.AuthenticatedUserResponse, error)) *MockAuthorizer_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// LoginUser provides a mock function with given fields: u
func (_m *MockAuthorizer) LoginUser(u entities.UserInput) (entities.AuthenticatedUserResponse, error) {
	ret := _m.Called(u)

	var r0 entities.AuthenticatedUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.UserInput) (entities.AuthenticatedUserResponse, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(entities.UserInput) entities.AuthenticatedUserResponse); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(entities.AuthenticatedUserResponse)
	}

	if rf, ok := ret.Get(1).(func(entities.UserInput) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthorizer_LoginUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoginUser'
type MockAuthorizer_LoginUser_Call struct {
	*mock.Call
}

// LoginUser is a helper method to define mock.On call
//   - u entities.UserInput
func (_e *MockAuthorizer_Expecter) LoginUser(u interface{}) *MockAuthorizer_LoginUser_Call {
	return &MockAuthorizer_LoginUser_Call{Call: _e.mock.On("LoginUser", u)}
}

func (_c *MockAuthorizer_LoginUser_Call) Run(run func(u entities.UserInput)) *MockAuthorizer_LoginUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.UserInput))
	})
	return _c
}

func (_c *MockAuthorizer_LoginUser_Call) Return(_a0 entities.AuthenticatedUserResponse, _a1 error) *MockAuthorizer_LoginUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthorizer_LoginUser_Call) RunAndReturn(run func(entities.UserInput) (entities.AuthenticatedUserResponse, error)) *MockAuthorizer_LoginUser_Call {
	_c.Call.Return(run)
	return _c
}

// RecoverPassword provides a mock function with given fields: u
func (_m *MockAuthorizer) RecoverPassword(u entities.RecoverPasswordInput) error {
	ret := _m.Called(u)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.RecoverPasswordInput) error); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAuthorizer_RecoverPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecoverPassword'
type MockAuthorizer_RecoverPassword_Call struct {
	*mock.Call
}

// RecoverPassword is a helper method to define mock.On call
//   - u entities.RecoverPasswordInput
func (_e *MockAuthorizer_Expecter) RecoverPassword(u interface{}) *MockAuthorizer_RecoverPassword_Call {
	return &MockAuthorizer_RecoverPassword_Call{Call: _e.mock.On("RecoverPassword", u)}
}

func (_c *MockAuthorizer_RecoverPassword_Call) Run(run func(u entities.RecoverPasswordInput)) *MockAuthorizer_RecoverPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.RecoverPasswordInput))
	})
	return _c
}

func (_c *MockAuthorizer_RecoverPassword_Call) Return(_a0 error) *MockAuthorizer_RecoverPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthorizer_RecoverPassword_Call) RunAndReturn(run func(entities.RecoverPasswordInput) error) *MockAuthorizer_RecoverPassword_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAuthorizer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAuthorizer creates a new instance of MockAuthorizer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAuthorizer(t mockConstructorTestingTNewMockAuthorizer) *MockAuthorizer {
	mock := &MockAuthorizer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
