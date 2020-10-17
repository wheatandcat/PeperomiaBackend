// Code generated by MockGen. DO NOT EDIT.
// Source: domain/item_detail.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	firestore "cloud.google.com/go/firestore"
	context "context"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/wheatandcat/PeperomiaBackend/domain"
	reflect "reflect"
)

// MockItemDetailRepository is a mock of ItemDetailRepository interface
type MockItemDetailRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemDetailRepositoryMockRecorder
}

// MockItemDetailRepositoryMockRecorder is the mock recorder for MockItemDetailRepository
type MockItemDetailRepositoryMockRecorder struct {
	mock *MockItemDetailRepository
}

// NewMockItemDetailRepository creates a new mock instance
func NewMockItemDetailRepository(ctrl *gomock.Controller) *MockItemDetailRepository {
	mock := &MockItemDetailRepository{ctrl: ctrl}
	mock.recorder = &MockItemDetailRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockItemDetailRepository) EXPECT() *MockItemDetailRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockItemDetailRepository) Create(ctx context.Context, f *firestore.Client, i domain.ItemDetailRecord, key domain.ItemDetailKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, f, i, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockItemDetailRepositoryMockRecorder) Create(ctx, f, i, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockItemDetailRepository)(nil).Create), ctx, f, i, key)
}

// Update mocks base method
func (m *MockItemDetailRepository) Update(ctx context.Context, f *firestore.Client, i domain.ItemDetailRecord, key domain.ItemDetailKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, f, i, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockItemDetailRepositoryMockRecorder) Update(ctx, f, i, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockItemDetailRepository)(nil).Update), ctx, f, i, key)
}

// Delete mocks base method
func (m *MockItemDetailRepository) Delete(ctx context.Context, f *firestore.Client, i domain.ItemDetailRecord, key domain.ItemDetailKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, f, i, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockItemDetailRepositoryMockRecorder) Delete(ctx, f, i, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockItemDetailRepository)(nil).Delete), ctx, f, i, key)
}

// Get mocks base method
func (m *MockItemDetailRepository) Get(ctx context.Context, f *firestore.Client, i domain.ItemDetailRecord, key domain.ItemDetailKey) (domain.ItemDetailRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, f, i, key)
	ret0, _ := ret[0].(domain.ItemDetailRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockItemDetailRepositoryMockRecorder) Get(ctx, f, i, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockItemDetailRepository)(nil).Get), ctx, f, i, key)
}

// FindByItemID mocks base method
func (m *MockItemDetailRepository) FindByItemID(ctx context.Context, f *firestore.Client, itemID string) ([]domain.ItemDetailRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByItemID", ctx, f, itemID)
	ret0, _ := ret[0].([]domain.ItemDetailRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByItemID indicates an expected call of FindByItemID
func (mr *MockItemDetailRepositoryMockRecorder) FindByItemID(ctx, f, itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByItemID", reflect.TypeOf((*MockItemDetailRepository)(nil).FindByItemID), ctx, f, itemID)
}
