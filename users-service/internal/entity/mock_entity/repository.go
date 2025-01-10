// Code generated by MockGen. DO NOT EDIT.
// Source: users-service/internal/entity (interfaces: UserRepository)
//
// Generated by this command:
//
//	mockgen -destination mock_entity/repository.go . UserRepository
//

// Package mock_entity is a generated GoMock package.
package mock_entity

import (
	reflect "reflect"
	entity "users-service/internal/entity"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockUserRepository) Save(user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockUserRepositoryMockRecorder) Save(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserRepository)(nil).Save), user)
}
