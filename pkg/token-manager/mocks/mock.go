// Code generated by MockGen. DO NOT EDIT.
// Source: type.go

// Package mock_tokenmanager is a generated GoMock package.
package mock_tokenmanager

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tokenmanager "github.com/resyahrial/go-commerce/pkg/token-manager"
)

// MockTokenManager is a mock of TokenManager interface.
type MockTokenManager struct {
	ctrl     *gomock.Controller
	recorder *MockTokenManagerMockRecorder
}

// MockTokenManagerMockRecorder is the mock recorder for MockTokenManager.
type MockTokenManagerMockRecorder struct {
	mock *MockTokenManager
}

// NewMockTokenManager creates a new mock instance.
func NewMockTokenManager(ctrl *gomock.Controller) *MockTokenManager {
	mock := &MockTokenManager{ctrl: ctrl}
	mock.recorder = &MockTokenManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenManager) EXPECT() *MockTokenManagerMockRecorder {
	return m.recorder
}

// GenerateAccess mocks base method.
func (m *MockTokenManager) GenerateAccess(claims tokenmanager.Claims) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccess", claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GenerateAccess indicates an expected call of GenerateAccess.
func (mr *MockTokenManagerMockRecorder) GenerateAccess(claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccess", reflect.TypeOf((*MockTokenManager)(nil).GenerateAccess), claims)
}

// GenerateRefresh mocks base method.
func (m *MockTokenManager) GenerateRefresh(claims tokenmanager.Claims) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefresh", claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GenerateRefresh indicates an expected call of GenerateRefresh.
func (mr *MockTokenManagerMockRecorder) GenerateRefresh(claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefresh", reflect.TypeOf((*MockTokenManager)(nil).GenerateRefresh), claims)
}

// ParseAccess mocks base method.
func (m *MockTokenManager) ParseAccess(tokenString string) (tokenmanager.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseAccess", tokenString)
	ret0, _ := ret[0].(tokenmanager.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseAccess indicates an expected call of ParseAccess.
func (mr *MockTokenManagerMockRecorder) ParseAccess(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseAccess", reflect.TypeOf((*MockTokenManager)(nil).ParseAccess), tokenString)
}

// ParseRefresh mocks base method.
func (m *MockTokenManager) ParseRefresh(tokenString string) (tokenmanager.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseRefresh", tokenString)
	ret0, _ := ret[0].(tokenmanager.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseRefresh indicates an expected call of ParseRefresh.
func (mr *MockTokenManagerMockRecorder) ParseRefresh(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseRefresh", reflect.TypeOf((*MockTokenManager)(nil).ParseRefresh), tokenString)
}
