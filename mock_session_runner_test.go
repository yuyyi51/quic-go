// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yuyyi51/quic-go (interfaces: SessionRunner)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/yuyyi51/quic-go/internal/protocol"
)

// MockSessionRunner is a mock of SessionRunner interface
type MockSessionRunner struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRunnerMockRecorder
}

// MockSessionRunnerMockRecorder is the mock recorder for MockSessionRunner
type MockSessionRunnerMockRecorder struct {
	mock *MockSessionRunner
}

// NewMockSessionRunner creates a new mock instance
func NewMockSessionRunner(ctrl *gomock.Controller) *MockSessionRunner {
	mock := &MockSessionRunner{ctrl: ctrl}
	mock.recorder = &MockSessionRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionRunner) EXPECT() *MockSessionRunnerMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockSessionRunner) Add(arg0 protocol.ConnectionID, arg1 packetHandler) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockSessionRunnerMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockSessionRunner)(nil).Add), arg0, arg1)
}

// AddResetToken mocks base method
func (m *MockSessionRunner) AddResetToken(arg0 protocol.StatelessResetToken, arg1 packetHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddResetToken", arg0, arg1)
}

// AddResetToken indicates an expected call of AddResetToken
func (mr *MockSessionRunnerMockRecorder) AddResetToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResetToken", reflect.TypeOf((*MockSessionRunner)(nil).AddResetToken), arg0, arg1)
}

// GetStatelessResetToken mocks base method
func (m *MockSessionRunner) GetStatelessResetToken(arg0 protocol.ConnectionID) protocol.StatelessResetToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatelessResetToken", arg0)
	ret0, _ := ret[0].(protocol.StatelessResetToken)
	return ret0
}

// GetStatelessResetToken indicates an expected call of GetStatelessResetToken
func (mr *MockSessionRunnerMockRecorder) GetStatelessResetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatelessResetToken", reflect.TypeOf((*MockSessionRunner)(nil).GetStatelessResetToken), arg0)
}

// Remove mocks base method
func (m *MockSessionRunner) Remove(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove
func (mr *MockSessionRunnerMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockSessionRunner)(nil).Remove), arg0)
}

// RemoveResetToken mocks base method
func (m *MockSessionRunner) RemoveResetToken(arg0 protocol.StatelessResetToken) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveResetToken", arg0)
}

// RemoveResetToken indicates an expected call of RemoveResetToken
func (mr *MockSessionRunnerMockRecorder) RemoveResetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveResetToken", reflect.TypeOf((*MockSessionRunner)(nil).RemoveResetToken), arg0)
}

// ReplaceWithClosed mocks base method
func (m *MockSessionRunner) ReplaceWithClosed(arg0 protocol.ConnectionID, arg1 packetHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReplaceWithClosed", arg0, arg1)
}

// ReplaceWithClosed indicates an expected call of ReplaceWithClosed
func (mr *MockSessionRunnerMockRecorder) ReplaceWithClosed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceWithClosed", reflect.TypeOf((*MockSessionRunner)(nil).ReplaceWithClosed), arg0, arg1)
}

// Retire mocks base method
func (m *MockSessionRunner) Retire(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Retire", arg0)
}

// Retire indicates an expected call of Retire
func (mr *MockSessionRunnerMockRecorder) Retire(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retire", reflect.TypeOf((*MockSessionRunner)(nil).Retire), arg0)
}

// RetireResetToken mocks base method
func (m *MockSessionRunner) RetireResetToken(arg0 protocol.StatelessResetToken) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RetireResetToken", arg0)
}

// RetireResetToken indicates an expected call of RetireResetToken
func (mr *MockSessionRunnerMockRecorder) RetireResetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetireResetToken", reflect.TypeOf((*MockSessionRunner)(nil).RetireResetToken), arg0)
}
