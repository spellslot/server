// Code generated by MockGen. DO NOT EDIT.
// Source: ./daos/spell.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/spellslot/server/models"
	reflect "reflect"
)

// MockSpellDAO is a mock of SpellDAO interface
type MockSpellDAO struct {
	ctrl     *gomock.Controller
	recorder *MockSpellDAOMockRecorder
}

// MockSpellDAOMockRecorder is the mock recorder for MockSpellDAO
type MockSpellDAOMockRecorder struct {
	mock *MockSpellDAO
}

// NewMockSpellDAO creates a new mock instance
func NewMockSpellDAO(ctrl *gomock.Controller) *MockSpellDAO {
	mock := &MockSpellDAO{ctrl: ctrl}
	mock.recorder = &MockSpellDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpellDAO) EXPECT() *MockSpellDAOMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockSpellDAO) Get() (*models.Spells, error) {
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*models.Spells)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSpellDAOMockRecorder) Get() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSpellDAO)(nil).Get))
}

// Create mocks base method
func (m *MockSpellDAO) Create(spell *models.Spell) (*models.Spell, error) {
	ret := m.ctrl.Call(m, "Create", spell)
	ret0, _ := ret[0].(*models.Spell)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockSpellDAOMockRecorder) Create(spell interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSpellDAO)(nil).Create), spell)
}
