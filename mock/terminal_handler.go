// Code generated by MockGen. DO NOT EDIT.
// Source: cmd/handler/terminal_handler.go
//
// Generated by this command:
//
//	mockgen -source=cmd/handler/terminal_handler.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	model "quake-log-reader/internal/application/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProcessEvent is a mock of ProcessEvent interface.
type MockProcessEvent struct {
	ctrl     *gomock.Controller
	recorder *MockProcessEventMockRecorder
}

// MockProcessEventMockRecorder is the mock recorder for MockProcessEvent.
type MockProcessEventMockRecorder struct {
	mock *MockProcessEvent
}

// NewMockProcessEvent creates a new mock instance.
func NewMockProcessEvent(ctrl *gomock.Controller) *MockProcessEvent {
	mock := &MockProcessEvent{ctrl: ctrl}
	mock.recorder = &MockProcessEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessEvent) EXPECT() *MockProcessEventMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockProcessEvent) Execute(logMessage string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Execute", logMessage)
}

// Execute indicates an expected call of Execute.
func (mr *MockProcessEventMockRecorder) Execute(logMessage any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockProcessEvent)(nil).Execute), logMessage)
}

// FinishOpenGames mocks base method.
func (m *MockProcessEvent) FinishOpenGames() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FinishOpenGames")
}

// FinishOpenGames indicates an expected call of FinishOpenGames.
func (mr *MockProcessEventMockRecorder) FinishOpenGames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishOpenGames", reflect.TypeOf((*MockProcessEvent)(nil).FinishOpenGames))
}

// GetAllGamesResult mocks base method.
func (m *MockProcessEvent) GetAllGamesResult() []*model.GameInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGamesResult")
	ret0, _ := ret[0].([]*model.GameInfo)
	return ret0
}

// GetAllGamesResult indicates an expected call of GetAllGamesResult.
func (mr *MockProcessEventMockRecorder) GetAllGamesResult() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGamesResult", reflect.TypeOf((*MockProcessEvent)(nil).GetAllGamesResult))
}
