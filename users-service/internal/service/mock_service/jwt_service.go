// Code generated by MockGen. DO NOT EDIT.
// Source: users-service/internal/service (interfaces: JwtService)
//
// Generated by this command:
//
//	mockgen -destination mock_service/jwt_service.go . JwtService
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockJwtService is a mock of JwtService interface.
type MockJwtService struct {
	ctrl     *gomock.Controller
	recorder *MockJwtServiceMockRecorder
	isgomock struct{}
}

// MockJwtServiceMockRecorder is the mock recorder for MockJwtService.
type MockJwtServiceMockRecorder struct {
	mock *MockJwtService
}

// NewMockJwtService creates a new mock instance.
func NewMockJwtService(ctrl *gomock.Controller) *MockJwtService {
	mock := &MockJwtService{ctrl: ctrl}
	mock.recorder = &MockJwtServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtService) EXPECT() *MockJwtServiceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockJwtService) GenerateToken(id uuid.UUID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockJwtServiceMockRecorder) GenerateToken(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockJwtService)(nil).GenerateToken), id)
}

// ValidateToken mocks base method.
func (m *MockJwtService) ValidateToken(token string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", token)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockJwtServiceMockRecorder) ValidateToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockJwtService)(nil).ValidateToken), token)
}