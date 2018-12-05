// Code generated by MockGen. DO NOT EDIT.
// Source: ./actpool/actioniterator/actioniterator.go

// Package mock_actioniterator is a generated GoMock package.
package mock_actioniterator

import (
	gomock "github.com/golang/mock/gomock"
	action "github.com/iotexproject/iotex-core/action"
	reflect "reflect"
)

// MockActionIterator is a mock of ActionIterator interface
type MockActionIterator struct {
	ctrl     *gomock.Controller
	recorder *MockActionIteratorMockRecorder
}

// MockActionIteratorMockRecorder is the mock recorder for MockActionIterator
type MockActionIteratorMockRecorder struct {
	mock *MockActionIterator
}

// NewMockActionIterator creates a new mock instance
func NewMockActionIterator(ctrl *gomock.Controller) *MockActionIterator {
	mock := &MockActionIterator{ctrl: ctrl}
	mock.recorder = &MockActionIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActionIterator) EXPECT() *MockActionIteratorMockRecorder {
	return m.recorder
}

// TopAction mocks base method
func (m *MockActionIterator) TopAction() *action.Action {
	ret := m.ctrl.Call(m, "TopAction")
	ret0, _ := ret[0].(*action.Action)
	return ret0
}

// TopAction indicates an expected call of TopAction
func (mr *MockActionIteratorMockRecorder) TopAction() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopAction", reflect.TypeOf((*MockActionIterator)(nil).TopAction))
}

// PopAction mocks base method
func (m *MockActionIterator) PopAction() {
	m.ctrl.Call(m, "PopAction")
}

// PopAction indicates an expected call of PopAction
func (mr *MockActionIteratorMockRecorder) PopAction() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopAction", reflect.TypeOf((*MockActionIterator)(nil).PopAction))
}

// LoadNextAction mocks base method
func (m *MockActionIterator) LoadNextAction() {
	m.ctrl.Call(m, "LoadNextAction")
}

// LoadNextAction indicates an expected call of LoadNextAction
func (mr *MockActionIteratorMockRecorder) LoadNextAction() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadNextAction", reflect.TypeOf((*MockActionIterator)(nil).LoadNextAction))
}

// AddAction mocks base method
func (m *MockActionIterator) AddAction(action action.Action) {
	m.ctrl.Call(m, "AddAction", action)
}

// AddAction indicates an expected call of AddAction
func (mr *MockActionIteratorMockRecorder) AddAction(action interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAction", reflect.TypeOf((*MockActionIterator)(nil).AddAction), action)
}
