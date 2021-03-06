// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_hasher is a generated GoMock package.
package mock_hasher

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHasher is a mock of Hasher interface.
type MockHasher struct {
	ctrl     *gomock.Controller
	recorder *MockHasherMockRecorder
}

// MockHasherMockRecorder is the mock recorder for MockHasher.
type MockHasherMockRecorder struct {
	mock *MockHasher
}

// NewMockHasher creates a new mock instance.
func NewMockHasher(ctrl *gomock.Controller) *MockHasher {
	mock := &MockHasher{ctrl: ctrl}
	mock.recorder = &MockHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasher) EXPECT() *MockHasherMockRecorder {
	return m.recorder
}

// Compare mocks base method.
func (m *MockHasher) Compare(password, hash string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", password, hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Compare indicates an expected call of Compare.
func (mr *MockHasherMockRecorder) Compare(password, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockHasher)(nil).Compare), password, hash)
}

// Hash mocks base method.
func (m *MockHasher) Hash(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash.
func (mr *MockHasherMockRecorder) Hash(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockHasher)(nil).Hash), password)
}
