// Code generated by MockGen. DO NOT EDIT.
// Source: domain/user.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	firestore "cloud.google.com/go/firestore"
	context "context"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/wheatandcat/PeperomiaBackend/domain"
	reflect "reflect"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserRepository) Create(ctx context.Context, f *firestore.Client, u domain.UserRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, f, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockUserRepositoryMockRecorder) Create(ctx, f, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), ctx, f, u)
}

// FindByUID mocks base method
func (m *MockUserRepository) FindByUID(ctx context.Context, f *firestore.Client, uid string) (domain.UserRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUID", ctx, f, uid)
	ret0, _ := ret[0].(domain.UserRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUID indicates an expected call of FindByUID
func (mr *MockUserRepositoryMockRecorder) FindByUID(ctx, f, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUID", reflect.TypeOf((*MockUserRepository)(nil).FindByUID), ctx, f, uid)
}

// ExistsByUID mocks base method
func (m *MockUserRepository) ExistsByUID(ctx context.Context, f *firestore.Client, uid string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByUID", ctx, f, uid)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByUID indicates an expected call of ExistsByUID
func (mr *MockUserRepositoryMockRecorder) ExistsByUID(ctx, f, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByUID", reflect.TypeOf((*MockUserRepository)(nil).ExistsByUID), ctx, f, uid)
}
