// Code generated by MockGen. DO NOT EDIT.
// Source: module/game/contract/game.go

// Package game is a generated GoMock package.
package game

import (
	reflect "reflect"
	contract "uwwolf/module/game/contract"
	state "uwwolf/module/game/state"
	types "uwwolf/types"

	gomock "github.com/golang/mock/gomock"
)

// MockGame is a mock of Game interface.
type MockGame struct {
	ctrl     *gomock.Controller
	recorder *MockGameMockRecorder
}

// MockGameMockRecorder is the mock recorder for MockGame.
type MockGameMockRecorder struct {
	mock *MockGame
}

// NewMockGame creates a new mock instance.
func NewMockGame(ctrl *gomock.Controller) *MockGame {
	mock := &MockGame{ctrl: ctrl}
	mock.recorder = &MockGameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGame) EXPECT() *MockGameMockRecorder {
	return m.recorder
}

// GetCurrentPhaseId mocks base method.
func (m *MockGame) GetCurrentPhaseId() types.PhaseId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentPhaseId")
	ret0, _ := ret[0].(types.PhaseId)
	return ret0
}

// GetCurrentPhaseId indicates an expected call of GetCurrentPhaseId.
func (mr *MockGameMockRecorder) GetCurrentPhaseId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentPhaseId", reflect.TypeOf((*MockGame)(nil).GetCurrentPhaseId))
}

// GetCurrentRoleId mocks base method.
func (m *MockGame) GetCurrentRoleId() types.RoleId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentRoleId")
	ret0, _ := ret[0].(types.RoleId)
	return ret0
}

// GetCurrentRoleId indicates an expected call of GetCurrentRoleId.
func (mr *MockGameMockRecorder) GetCurrentRoleId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentRoleId", reflect.TypeOf((*MockGame)(nil).GetCurrentRoleId))
}

// GetCurrentRoundId mocks base method.
func (m *MockGame) GetCurrentRoundId() types.RoundId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentRoundId")
	ret0, _ := ret[0].(types.RoundId)
	return ret0
}

// GetCurrentRoundId indicates an expected call of GetCurrentRoundId.
func (mr *MockGameMockRecorder) GetCurrentRoundId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentRoundId", reflect.TypeOf((*MockGame)(nil).GetCurrentRoundId))
}

// GetPlayer mocks base method.
func (m *MockGame) GetPlayer(playerId types.PlayerId) contract.Player {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlayer", playerId)
	ret0, _ := ret[0].(contract.Player)
	return ret0
}

// GetPlayer indicates an expected call of GetPlayer.
func (mr *MockGameMockRecorder) GetPlayer(playerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlayer", reflect.TypeOf((*MockGame)(nil).GetPlayer), playerId)
}

// GetPoll mocks base method.
func (m *MockGame) GetPoll(factionId types.FactionId) *state.Poll {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoll", factionId)
	ret0, _ := ret[0].(*state.Poll)
	return ret0
}

// GetPoll indicates an expected call of GetPoll.
func (mr *MockGameMockRecorder) GetPoll(factionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoll", reflect.TypeOf((*MockGame)(nil).GetPoll), factionId)
}

// IsStarted mocks base method.
func (m *MockGame) IsStarted() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStarted")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStarted indicates an expected call of IsStarted.
func (mr *MockGameMockRecorder) IsStarted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStarted", reflect.TypeOf((*MockGame)(nil).IsStarted))
}

// KillPlayer mocks base method.
func (m *MockGame) KillPlayer(playerId types.PlayerId) contract.Player {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KillPlayer", playerId)
	ret0, _ := ret[0].(contract.Player)
	return ret0
}

// KillPlayer indicates an expected call of KillPlayer.
func (mr *MockGameMockRecorder) KillPlayer(playerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KillPlayer", reflect.TypeOf((*MockGame)(nil).KillPlayer), playerId)
}

// RequestAction mocks base method.
func (m *MockGame) RequestAction(playerId types.PlayerId, req *types.ActionRequest) *types.ActionResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestAction", playerId, req)
	ret0, _ := ret[0].(*types.ActionResponse)
	return ret0
}

// RequestAction indicates an expected call of RequestAction.
func (mr *MockGameMockRecorder) RequestAction(playerId, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAction", reflect.TypeOf((*MockGame)(nil).RequestAction), playerId, req)
}

// Start mocks base method.
func (m *MockGame) Start() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockGameMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockGame)(nil).Start))
}
