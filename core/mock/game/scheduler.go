// Code generated by MockGen. DO NOT EDIT.
// Source: game/contract/scheduler.go

// Package gamemock is a generated GoMock package.
package gamemock

import (
	reflect "reflect"
	types "uwwolf/game/types"

	gomock "github.com/golang/mock/gomock"
)

// MockScheduler is a mock of Scheduler interface.
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler.
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance.
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// AddPlayerTurn mocks base method.
func (m *MockScheduler) AddPlayerTurn(newTurn *types.NewPlayerTurn) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlayerTurn", newTurn)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AddPlayerTurn indicates an expected call of AddPlayerTurn.
func (mr *MockSchedulerMockRecorder) AddPlayerTurn(newTurn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlayerTurn", reflect.TypeOf((*MockScheduler)(nil).AddPlayerTurn), newTurn)
}

// IsEmptyPhase mocks base method.
func (m *MockScheduler) IsEmptyPhase(phaseID types.PhaseID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmptyPhase", phaseID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmptyPhase indicates an expected call of IsEmptyPhase.
func (mr *MockSchedulerMockRecorder) IsEmptyPhase(phaseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmptyPhase", reflect.TypeOf((*MockScheduler)(nil).IsEmptyPhase), phaseID)
}

// NextTurn mocks base method.
func (m *MockScheduler) NextTurn() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextTurn")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NextTurn indicates an expected call of NextTurn.
func (mr *MockSchedulerMockRecorder) NextTurn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextTurn", reflect.TypeOf((*MockScheduler)(nil).NextTurn))
}

// Phase mocks base method.
func (m *MockScheduler) Phase() []types.Turn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Phase")
	ret0, _ := ret[0].([]types.Turn)
	return ret0
}

// Phase indicates an expected call of Phase.
func (mr *MockSchedulerMockRecorder) Phase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Phase", reflect.TypeOf((*MockScheduler)(nil).Phase))
}

// PhaseID mocks base method.
func (m *MockScheduler) PhaseID() types.PhaseID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhaseID")
	ret0, _ := ret[0].(types.PhaseID)
	return ret0
}

// PhaseID indicates an expected call of PhaseID.
func (mr *MockSchedulerMockRecorder) PhaseID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhaseID", reflect.TypeOf((*MockScheduler)(nil).PhaseID))
}

// PlayablePlayerIDs mocks base method.
func (m *MockScheduler) PlayablePlayerIDs() []types.PlayerID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlayablePlayerIDs")
	ret0, _ := ret[0].([]types.PlayerID)
	return ret0
}

// PlayablePlayerIDs indicates an expected call of PlayablePlayerIDs.
func (mr *MockSchedulerMockRecorder) PlayablePlayerIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlayablePlayerIDs", reflect.TypeOf((*MockScheduler)(nil).PlayablePlayerIDs))
}

// RemovePlayerTurn mocks base method.
func (m *MockScheduler) RemovePlayerTurn(removedTurn *types.RemovedPlayerTurn) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePlayerTurn", removedTurn)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RemovePlayerTurn indicates an expected call of RemovePlayerTurn.
func (mr *MockSchedulerMockRecorder) RemovePlayerTurn(removedTurn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePlayerTurn", reflect.TypeOf((*MockScheduler)(nil).RemovePlayerTurn), removedTurn)
}

// RoundID mocks base method.
func (m *MockScheduler) RoundID() types.RoundID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoundID")
	ret0, _ := ret[0].(types.RoundID)
	return ret0
}

// RoundID indicates an expected call of RoundID.
func (mr *MockSchedulerMockRecorder) RoundID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoundID", reflect.TypeOf((*MockScheduler)(nil).RoundID))
}

// Turn mocks base method.
func (m *MockScheduler) Turn() types.Turn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Turn")
	ret0, _ := ret[0].(types.Turn)
	return ret0
}

// Turn indicates an expected call of Turn.
func (mr *MockSchedulerMockRecorder) Turn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Turn", reflect.TypeOf((*MockScheduler)(nil).Turn))
}

// TurnID mocks base method.
func (m *MockScheduler) TurnID() types.TurnID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TurnID")
	ret0, _ := ret[0].(types.TurnID)
	return ret0
}

// TurnID indicates an expected call of TurnID.
func (mr *MockSchedulerMockRecorder) TurnID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TurnID", reflect.TypeOf((*MockScheduler)(nil).TurnID))
}
