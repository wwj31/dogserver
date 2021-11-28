// Code generated by MockGen. DO NOT EDIT.
// Source: foo.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLife is a mock of Life interface.
type MockLife struct {
	ctrl     *gomock.Controller
	recorder *MockLifeMockRecorder
}

// MockLifeMockRecorder is the mock recorder for MockLife.
type MockLifeMockRecorder struct {
	mock *MockLife
}

// NewMockLife creates a new mock instance.
func NewMockLife(ctrl *gomock.Controller) *MockLife {
	mock := &MockLife{ctrl: ctrl}
	mock.recorder = &MockLifeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLife) EXPECT() *MockLifeMockRecorder {
	return m.recorder
}

// BuyHouse mocks base method.
func (m *MockLife) BuyHouse(money int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyHouse", money)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyHouse indicates an expected call of BuyHouse.
func (mr *MockLifeMockRecorder) BuyHouse(money interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyHouse", reflect.TypeOf((*MockLife)(nil).BuyHouse), money)
}

// GoodGoodStudy mocks base method.
func (m *MockLife) GoodGoodStudy(money int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GoodGoodStudy", money)
	ret0, _ := ret[0].(error)
	return ret0
}

// GoodGoodStudy indicates an expected call of GoodGoodStudy.
func (mr *MockLifeMockRecorder) GoodGoodStudy(money interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoodGoodStudy", reflect.TypeOf((*MockLife)(nil).GoodGoodStudy), money)
}

// Marry mocks base method.
func (m *MockLife) Marry(money int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marry", money)
	ret0, _ := ret[0].(error)
	return ret0
}

// Marry indicates an expected call of Marry.
func (mr *MockLifeMockRecorder) Marry(money interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marry", reflect.TypeOf((*MockLife)(nil).Marry), money)
}
