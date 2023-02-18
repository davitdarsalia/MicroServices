// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	entities "auth/internal/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorizer is a mock of Authorizer interface.
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer.
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance.
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// CheckToken mocks base method.
func (m *MockAuthorizer) CheckToken(authToken, signKey string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckToken", authToken, signKey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckToken indicates an expected call of CheckToken.
func (mr *MockAuthorizerMockRecorder) CheckToken(authToken, signKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckToken", reflect.TypeOf((*MockAuthorizer)(nil).CheckToken), authToken, signKey)
}

// CreateUser mocks base method.
func (m *MockAuthorizer) CreateUser(u entities.User) (entities.AuthenticatedUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", u)
	ret0, _ := ret[0].(entities.AuthenticatedUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizerMockRecorder) CreateUser(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorizer)(nil).CreateUser), u)
}

// LoginUser mocks base method.
func (m *MockAuthorizer) LoginUser(u entities.UserInput) (entities.AuthenticatedUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", u)
	ret0, _ := ret[0].(entities.AuthenticatedUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockAuthorizerMockRecorder) LoginUser(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockAuthorizer)(nil).LoginUser), u)
}

// RecoverPassword mocks base method.
func (m *MockAuthorizer) RecoverPassword(u entities.RecoverPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecoverPassword", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecoverPassword indicates an expected call of RecoverPassword.
func (mr *MockAuthorizerMockRecorder) RecoverPassword(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverPassword", reflect.TypeOf((*MockAuthorizer)(nil).RecoverPassword), u)
}
