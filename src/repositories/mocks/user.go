// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/user.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/trilobit/go-chat/src/models"
	reflect "reflect"
)

// MockUser is a mock of User interface
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUser) Create(email, pswdHash string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", email, pswdHash)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserMockRecorder) Create(email, pswdHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUser)(nil).Create), email, pswdHash)
}

// UpdateToken mocks base method
func (m *MockUser) UpdateToken(user *models.User, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", user, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToken indicates an expected call of UpdateToken
func (mr *MockUserMockRecorder) UpdateToken(user, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockUser)(nil).UpdateToken), user, token)
}

// FindByEmail mocks base method
func (m *MockUser) FindByEmail(email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail
func (mr *MockUserMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUser)(nil).FindByEmail), email)
}

// FindByToken mocks base method
func (m *MockUser) FindByToken(token string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByToken", token)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByToken indicates an expected call of FindByToken
func (mr *MockUserMockRecorder) FindByToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByToken", reflect.TypeOf((*MockUser)(nil).FindByToken), token)
}
