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

// AddSlot mocks base method.
func (m *MockScheduler) AddSlot(newSlot *types.NewTurnSlot) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSlot", newSlot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AddSlot indicates an expected call of AddSlot.
func (mr *MockSchedulerMockRecorder) AddSlot(newSlot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSlot", reflect.TypeOf((*MockScheduler)(nil).AddSlot), newSlot)
}

// CanPlay mocks base method.
func (m *MockScheduler) CanPlay(playerID types.PlayerID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanPlay", playerID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanPlay indicates an expected call of CanPlay.
func (mr *MockSchedulerMockRecorder) CanPlay(playerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanPlay", reflect.TypeOf((*MockScheduler)(nil).CanPlay), playerID)
}

// FreezeSlot mocks base method.
func (m *MockScheduler) FreezeSlot(frozenSlot *types.FreezeTurnSlot) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FreezeSlot", frozenSlot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// FreezeSlot indicates an expected call of FreezeSlot.
func (mr *MockSchedulerMockRecorder) FreezeSlot(frozenSlot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreezeSlot", reflect.TypeOf((*MockScheduler)(nil).FreezeSlot), frozenSlot)
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
func (m *MockScheduler) Phase() map[types.TurnID]types.Turn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Phase")
	ret0, _ := ret[0].(map[types.TurnID]types.Turn)
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

// RemoveSlot mocks base method.
func (m *MockScheduler) RemoveSlot(removedSlot *types.RemovedTurnSlot) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSlot", removedSlot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RemoveSlot indicates an expected call of RemoveSlot.
func (mr *MockSchedulerMockRecorder) RemoveSlot(removedSlot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSlot", reflect.TypeOf((*MockScheduler)(nil).RemoveSlot), removedSlot)
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
