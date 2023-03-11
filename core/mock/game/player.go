// Code generated by MockGen. DO NOT EDIT.
// Source: game/contract/player.go

// Package mock_game is a generated GoMock package.
package mock_game

import (
	reflect "reflect"
	contract "uwwolf/game/contract"
	types "uwwolf/game/types"

	gomock "github.com/golang/mock/gomock"
)

// MockPlayer is a mock of Player interface.
type MockPlayer struct {
	ctrl     *gomock.Controller
	recorder *MockPlayerMockRecorder
}

// MockPlayerMockRecorder is the mock recorder for MockPlayer.
type MockPlayerMockRecorder struct {
	mock *MockPlayer
}

// NewMockPlayer creates a new mock instance.
func NewMockPlayer(ctrl *gomock.Controller) *MockPlayer {
	mock := &MockPlayer{ctrl: ctrl}
	mock.recorder = &MockPlayerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlayer) EXPECT() *MockPlayerMockRecorder {
	return m.recorder
}

// ActivateAbility mocks base method.
func (m *MockPlayer) ActivateAbility(req *types.ActivateAbilityRequest) *types.ActionResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateAbility", req)
	ret0, _ := ret[0].(*types.ActionResponse)
	return ret0
}

// ActivateAbility indicates an expected call of ActivateAbility.
func (mr *MockPlayerMockRecorder) ActivateAbility(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateAbility", reflect.TypeOf((*MockPlayer)(nil).ActivateAbility), req)
}

// AssignRole mocks base method.
func (m *MockPlayer) AssignRole(roleID types.RoleID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignRole", roleID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignRole indicates an expected call of AssignRole.
func (mr *MockPlayerMockRecorder) AssignRole(roleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignRole", reflect.TypeOf((*MockPlayer)(nil).AssignRole), roleID)
}

// Die mocks base method.
func (m *MockPlayer) Die(isExited bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Die", isExited)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Die indicates an expected call of Die.
func (mr *MockPlayerMockRecorder) Die(isExited interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Die", reflect.TypeOf((*MockPlayer)(nil).Die), isExited)
}

// FactionID mocks base method.
func (m *MockPlayer) FactionID() types.FactionID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FactionID")
	ret0, _ := ret[0].(types.FactionID)
	return ret0
}

// FactionID indicates an expected call of FactionID.
func (mr *MockPlayerMockRecorder) FactionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FactionID", reflect.TypeOf((*MockPlayer)(nil).FactionID))
}

// ID mocks base method.
func (m *MockPlayer) ID() types.PlayerID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(types.PlayerID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockPlayerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockPlayer)(nil).ID))
}

// IsDead mocks base method.
func (m *MockPlayer) IsDead() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDead")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDead indicates an expected call of IsDead.
func (mr *MockPlayerMockRecorder) IsDead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDead", reflect.TypeOf((*MockPlayer)(nil).IsDead))
}

// MainRoleID mocks base method.
func (m *MockPlayer) MainRoleID() types.RoleID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MainRoleID")
	ret0, _ := ret[0].(types.RoleID)
	return ret0
}

// MainRoleID indicates an expected call of MainRoleID.
func (mr *MockPlayerMockRecorder) MainRoleID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MainRoleID", reflect.TypeOf((*MockPlayer)(nil).MainRoleID))
}

// RevokeRole mocks base method.
func (m *MockPlayer) RevokeRole(roleID types.RoleID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeRole", roleID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeRole indicates an expected call of RevokeRole.
func (mr *MockPlayerMockRecorder) RevokeRole(roleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeRole", reflect.TypeOf((*MockPlayer)(nil).RevokeRole), roleID)
}

// RoleIDs mocks base method.
func (m *MockPlayer) RoleIDs() []types.RoleID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleIDs")
	ret0, _ := ret[0].([]types.RoleID)
	return ret0
}

// RoleIDs indicates an expected call of RoleIDs.
func (mr *MockPlayerMockRecorder) RoleIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleIDs", reflect.TypeOf((*MockPlayer)(nil).RoleIDs))
}

// Roles mocks base method.
func (m *MockPlayer) Roles() map[types.RoleID]contract.Role {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Roles")
	ret0, _ := ret[0].(map[types.RoleID]contract.Role)
	return ret0
}

// Roles indicates an expected call of Roles.
func (mr *MockPlayerMockRecorder) Roles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Roles", reflect.TypeOf((*MockPlayer)(nil).Roles))
}

// SetFactionID mocks base method.
func (m *MockPlayer) SetFactionID(factionID types.FactionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFactionID", factionID)
}

// SetFactionID indicates an expected call of SetFactionID.
func (mr *MockPlayerMockRecorder) SetFactionID(factionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFactionID", reflect.TypeOf((*MockPlayer)(nil).SetFactionID), factionID)
}
