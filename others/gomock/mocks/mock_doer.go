// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sgreben/testing-with-gomock/doer (interfaces: Doer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDoer is a mock of Doer interface
type MockDoer struct {
	ctrl     *gomock.Controller
	recorder *MockDoerMockRecorder
}

// MockDoerMockRecorder is the mock recorder for MockDoer
type MockDoerMockRecorder struct {
	mock *MockDoer
}

// NewMockDoer creates a new mock instance
func NewMockDoer(ctrl *gomock.Controller) *MockDoer {
	mock := &MockDoer{ctrl: ctrl}
	mock.recorder = &MockDoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDoer) EXPECT() *MockDoerMockRecorder {
	return m.recorder
}

// DoSomething mocks base method
func (m *MockDoer) DoSomething(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoSomething", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoSomething indicates an expected call of DoSomething
func (mr *MockDoerMockRecorder) DoSomething(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoSomething", reflect.TypeOf((*MockDoer)(nil).DoSomething), arg0, arg1)
}
