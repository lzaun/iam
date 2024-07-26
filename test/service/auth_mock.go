// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/auth.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	dto "github.com/lantonster/iam/internal/dto"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthService) Login(c *gin.Context, req *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", c, req)
	ret0, _ := ret[0].(*dto.AuthLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceMockRecorder) Login(c, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), c, req)
}

// UserInfo mocks base method.
func (m *MockAuthService) UserInfo(c *gin.Context) (*dto.AuthUserInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserInfo", c)
	ret0, _ := ret[0].(*dto.AuthUserInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserInfo indicates an expected call of UserInfo.
func (mr *MockAuthServiceMockRecorder) UserInfo(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserInfo", reflect.TypeOf((*MockAuthService)(nil).UserInfo), c)
}

// UsernameAvailable mocks base method.
func (m *MockAuthService) UsernameAvailable(c *gin.Context, req *dto.AuthUsernameAvailableRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsernameAvailable", c, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UsernameAvailable indicates an expected call of UsernameAvailable.
func (mr *MockAuthServiceMockRecorder) UsernameAvailable(c, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsernameAvailable", reflect.TypeOf((*MockAuthService)(nil).UsernameAvailable), c, req)
}
