// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonmp/trv-airport-api/internal/repo (interfaces: AirportRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/ozonmp/trv-airport-api/internal/model"
)

// MockAirportRepo is a mock of AirportRepo interface.
type MockAirportRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAirportRepoMockRecorder
}

// MockAirportRepoMockRecorder is the mock recorder for MockAirportRepo.
type MockAirportRepoMockRecorder struct {
	mock *MockAirportRepo
}

// NewMockAirportRepo creates a new mock instance.
func NewMockAirportRepo(ctrl *gomock.Controller) *MockAirportRepo {
	mock := &MockAirportRepo{ctrl: ctrl}
	mock.recorder = &MockAirportRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAirportRepo) EXPECT() *MockAirportRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockAirportRepo) Add(arg0 context.Context, arg1, arg2 string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockAirportRepoMockRecorder) Add(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAirportRepo)(nil).Add), arg0, arg1, arg2)
}

// Close mocks base method.
func (m *MockAirportRepo) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAirportRepoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAirportRepo)(nil).Close))
}

// Delete mocks base method.
func (m *MockAirportRepo) Delete(arg0 context.Context, arg1 uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAirportRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAirportRepo)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockAirportRepo) Get(arg0 context.Context, arg1 uint64) (*model.Airport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.Airport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAirportRepoMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAirportRepo)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockAirportRepo) List(arg0 context.Context, arg1, arg2 uint64) ([]model.Airport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Airport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAirportRepoMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAirportRepo)(nil).List), arg0, arg1, arg2)
}
