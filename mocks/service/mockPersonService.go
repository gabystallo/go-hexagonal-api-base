// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gabystallo/go-hexagonal-api-base/service (interfaces: PersonService)

// Package service is a generated GoMock package.
package service

import (
	dto "github.com/gabystallo/go-hexagonal-api-base/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPersonService is a mock of PersonService interface.
type MockPersonService struct {
	ctrl     *gomock.Controller
	recorder *MockPersonServiceMockRecorder
}

// MockPersonServiceMockRecorder is the mock recorder for MockPersonService.
type MockPersonServiceMockRecorder struct {
	mock *MockPersonService
}

// NewMockPersonService creates a new mock instance.
func NewMockPersonService(ctrl *gomock.Controller) *MockPersonService {
	mock := &MockPersonService{ctrl: ctrl}
	mock.recorder = &MockPersonServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonService) EXPECT() *MockPersonServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPersonService) Create(arg0 *dto.CreatePersonRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPersonServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPersonService)(nil).Create), arg0)
}

// GetAll mocks base method.
func (m *MockPersonService) GetAll() ([]dto.PersonResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]dto.PersonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPersonServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPersonService)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockPersonService) GetById(arg0 int) (*dto.PersonResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*dto.PersonResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockPersonServiceMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPersonService)(nil).GetById), arg0)
}
