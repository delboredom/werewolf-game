// Code generated by MockGen. DO NOT EDIT.
// Source: game/contract/role.go

// Package gamemock is a generated GoMock package.
package gamemock

import (
	reflect "reflect"
	types "uwwolf/game/types"

	gomock "github.com/golang/mock/gomock"
)

// MockRole is a mock of Role interface.
type MockRole struct {
	ctrl     *gomock.Controller
	recorder *MockRoleMockRecorder
}

// MockRoleMockRecorder is the mock recorder for MockRole.
type MockRoleMockRecorder struct {
	mock *MockRole
}

// NewMockRole creates a new mock instance.
func NewMockRole(ctrl *gomock.Controller) *MockRole {
	mock := &MockRole{ctrl: ctrl}
	mock.recorder = &MockRoleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRole) EXPECT() *MockRoleMockRecorder {
	return m.recorder
}

// ActivateAbility mocks base method.
func (m *MockRole) ActivateAbility(req *types.ActivateAbilityRequest) *types.ActionResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateAbility", req)
	ret0, _ := ret[0].(*types.ActionResponse)
	return ret0
}

// ActivateAbility indicates an expected call of ActivateAbility.
func (mr *MockRoleMockRecorder) ActivateAbility(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateAbility", reflect.TypeOf((*MockRole)(nil).ActivateAbility), req)
}

// ActiveLimit mocks base method.
func (m *MockRole) ActiveLimit(index int) types.Limit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveLimit", index)
	ret0, _ := ret[0].(types.Limit)
	return ret0
}

// ActiveLimit indicates an expected call of ActiveLimit.
func (mr *MockRoleMockRecorder) ActiveLimit(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveLimit", reflect.TypeOf((*MockRole)(nil).ActiveLimit), index)
}

// AfterDeath mocks base method.
func (m *MockRole) AfterDeath() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AfterDeath")
}

// AfterDeath indicates an expected call of AfterDeath.
func (mr *MockRoleMockRecorder) AfterDeath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterDeath", reflect.TypeOf((*MockRole)(nil).AfterDeath))
}

// BeforeDeath mocks base method.
func (m *MockRole) BeforeDeath() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeDeath")
	ret0, _ := ret[0].(bool)
	return ret0
}

// BeforeDeath indicates an expected call of BeforeDeath.
func (mr *MockRoleMockRecorder) BeforeDeath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDeath", reflect.TypeOf((*MockRole)(nil).BeforeDeath))
}

// FactionID mocks base method.
func (m *MockRole) FactionID() types.FactionID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FactionID")
	ret0, _ := ret[0].(types.FactionID)
	return ret0
}

// FactionID indicates an expected call of FactionID.
func (mr *MockRoleMockRecorder) FactionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FactionID", reflect.TypeOf((*MockRole)(nil).FactionID))
}

// ID mocks base method.
func (m *MockRole) ID() types.RoleID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(types.RoleID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockRoleMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockRole)(nil).ID))
}

// RegisterTurns mocks base method.
func (m *MockRole) RegisterTurns() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterTurns")
}

// RegisterTurns indicates an expected call of RegisterTurns.
func (mr *MockRoleMockRecorder) RegisterTurns() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTurns", reflect.TypeOf((*MockRole)(nil).RegisterTurns))
}
