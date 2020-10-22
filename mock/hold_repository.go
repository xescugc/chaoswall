// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xescugc/chaoswall/hold (interfaces: Repository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	hold "github.com/xescugc/chaoswall/hold"
	reflect "reflect"
)

// HoldRepository is a mock of Repository interface
type HoldRepository struct {
	ctrl     *gomock.Controller
	recorder *HoldRepositoryMockRecorder
}

// HoldRepositoryMockRecorder is the mock recorder for HoldRepository
type HoldRepositoryMockRecorder struct {
	mock *HoldRepository
}

// NewHoldRepository creates a new mock instance
func NewHoldRepository(ctrl *gomock.Controller) *HoldRepository {
	mock := &HoldRepository{ctrl: ctrl}
	mock.recorder = &HoldRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *HoldRepository) EXPECT() *HoldRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *HoldRepository) Create(arg0 context.Context, arg1, arg2 string, arg3 hold.Hold) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *HoldRepositoryMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*HoldRepository)(nil).Create), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *HoldRepository) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *HoldRepositoryMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*HoldRepository)(nil).Delete), arg0, arg1, arg2)
}

// Filter mocks base method
func (m *HoldRepository) Filter(arg0 context.Context, arg1, arg2 string) ([]*hold.Hold, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*hold.Hold)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *HoldRepositoryMockRecorder) Filter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*HoldRepository)(nil).Filter), arg0, arg1, arg2)
}
