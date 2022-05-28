// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package mock_seller is a generated GoMock package.
package mock_seller

import (
	gomock "github.com/golang/mock/gomock"
)

// MockSellerRepo is a mock of SellerRepo interface.
type MockSellerRepo struct {
	ctrl     *gomock.Controller
	recorder *MockSellerRepoMockRecorder
}

// MockSellerRepoMockRecorder is the mock recorder for MockSellerRepo.
type MockSellerRepoMockRecorder struct {
	mock *MockSellerRepo
}

// NewMockSellerRepo creates a new mock instance.
func NewMockSellerRepo(ctrl *gomock.Controller) *MockSellerRepo {
	mock := &MockSellerRepo{ctrl: ctrl}
	mock.recorder = &MockSellerRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSellerRepo) EXPECT() *MockSellerRepoMockRecorder {
	return m.recorder
}