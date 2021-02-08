// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/marmotedu/iam/internal/apiserver/store (interfaces: Factory,UserStore,SecretStore,PolicyStore)

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/marmotedu/api/apiserver/v1"
	v10 "github.com/marmotedu/component-base/pkg/meta/v1"
)

// MockFactory is a mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Policies mocks base method
func (m *MockFactory) Policies() PolicyStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Policies")
	ret0, _ := ret[0].(PolicyStore)
	return ret0
}

// Policies indicates an expected call of Policies
func (mr *MockFactoryMockRecorder) Policies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Policies", reflect.TypeOf((*MockFactory)(nil).Policies))
}

// Secrets mocks base method
func (m *MockFactory) Secrets() SecretStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Secrets")
	ret0, _ := ret[0].(SecretStore)
	return ret0
}

// Secrets indicates an expected call of Secrets
func (mr *MockFactoryMockRecorder) Secrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Secrets", reflect.TypeOf((*MockFactory)(nil).Secrets))
}

// Users mocks base method
func (m *MockFactory) Users() UserStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(UserStore)
	return ret0
}

// Users indicates an expected call of Users
func (mr *MockFactoryMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockFactory)(nil).Users))
}

// MockUserStore is a mock of UserStore interface
type MockUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoreMockRecorder
}

// MockUserStoreMockRecorder is the mock recorder for MockUserStore
type MockUserStoreMockRecorder struct {
	mock *MockUserStore
}

// NewMockUserStore creates a new mock instance
func NewMockUserStore(ctrl *gomock.Controller) *MockUserStore {
	mock := &MockUserStore{ctrl: ctrl}
	mock.recorder = &MockUserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserStore) EXPECT() *MockUserStoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserStore) Create(arg0 context.Context, arg1 *v1.User, arg2 v10.CreateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockUserStoreMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserStore)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method
func (m *MockUserStore) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUserStoreMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserStore)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method
func (m *MockUserStore) DeleteCollection(arg0 context.Context, arg1 []string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockUserStoreMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockUserStore)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockUserStore) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUserStoreMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserStore)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockUserStore) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockUserStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserStore)(nil).List), arg0, arg1)
}

// Update mocks base method
func (m *MockUserStore) Update(arg0 context.Context, arg1 *v1.User, arg2 v10.UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockUserStoreMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserStore)(nil).Update), arg0, arg1, arg2)
}

// MockSecretStore is a mock of SecretStore interface
type MockSecretStore struct {
	ctrl     *gomock.Controller
	recorder *MockSecretStoreMockRecorder
}

// MockSecretStoreMockRecorder is the mock recorder for MockSecretStore
type MockSecretStoreMockRecorder struct {
	mock *MockSecretStore
}

// NewMockSecretStore creates a new mock instance
func NewMockSecretStore(ctrl *gomock.Controller) *MockSecretStore {
	mock := &MockSecretStore{ctrl: ctrl}
	mock.recorder = &MockSecretStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretStore) EXPECT() *MockSecretStoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockSecretStore) Create(arg0 context.Context, arg1 *v1.Secret, arg2 v10.CreateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockSecretStoreMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSecretStore)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method
func (m *MockSecretStore) Delete(arg0 context.Context, arg1, arg2 string, arg3 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockSecretStoreMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSecretStore)(nil).Delete), arg0, arg1, arg2, arg3)
}

// DeleteCollection mocks base method
func (m *MockSecretStore) DeleteCollection(arg0 context.Context, arg1 string, arg2 []string, arg3 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockSecretStoreMockRecorder) DeleteCollection(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockSecretStore)(nil).DeleteCollection), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockSecretStore) Get(arg0 context.Context, arg1, arg2 string, arg3 v10.GetOptions) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSecretStoreMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecretStore)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockSecretStore) List(arg0 context.Context, arg1 string, arg2 v10.ListOptions) (*v1.SecretList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.SecretList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockSecretStoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecretStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method
func (m *MockSecretStore) Update(arg0 context.Context, arg1 *v1.Secret, arg2 v10.UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockSecretStoreMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSecretStore)(nil).Update), arg0, arg1, arg2)
}

// MockPolicyStore is a mock of PolicyStore interface
type MockPolicyStore struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyStoreMockRecorder
}

// MockPolicyStoreMockRecorder is the mock recorder for MockPolicyStore
type MockPolicyStoreMockRecorder struct {
	mock *MockPolicyStore
}

// NewMockPolicyStore creates a new mock instance
func NewMockPolicyStore(ctrl *gomock.Controller) *MockPolicyStore {
	mock := &MockPolicyStore{ctrl: ctrl}
	mock.recorder = &MockPolicyStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPolicyStore) EXPECT() *MockPolicyStoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockPolicyStore) Create(arg0 context.Context, arg1 *v1.Policy, arg2 v10.CreateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockPolicyStoreMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPolicyStore)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method
func (m *MockPolicyStore) Delete(arg0 context.Context, arg1, arg2 string, arg3 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPolicyStoreMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPolicyStore)(nil).Delete), arg0, arg1, arg2, arg3)
}

// DeleteByUser mocks base method
func (m *MockPolicyStore) DeleteByUser(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByUser indicates an expected call of DeleteByUser
func (mr *MockPolicyStoreMockRecorder) DeleteByUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUser", reflect.TypeOf((*MockPolicyStore)(nil).DeleteByUser), arg0, arg1, arg2)
}

// DeleteCollection mocks base method
func (m *MockPolicyStore) DeleteCollection(arg0 context.Context, arg1 string, arg2 []string, arg3 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockPolicyStoreMockRecorder) DeleteCollection(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockPolicyStore)(nil).DeleteCollection), arg0, arg1, arg2, arg3)
}

// DeleteCollectionByUser mocks base method
func (m *MockPolicyStore) DeleteCollectionByUser(arg0 context.Context, arg1 []string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollectionByUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollectionByUser indicates an expected call of DeleteCollectionByUser
func (mr *MockPolicyStoreMockRecorder) DeleteCollectionByUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollectionByUser", reflect.TypeOf((*MockPolicyStore)(nil).DeleteCollectionByUser), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockPolicyStore) Get(arg0 context.Context, arg1, arg2 string, arg3 v10.GetOptions) (*v1.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPolicyStoreMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPolicyStore)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockPolicyStore) List(arg0 context.Context, arg1 string, arg2 v10.ListOptions) (*v1.PolicyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockPolicyStoreMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPolicyStore)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method
func (m *MockPolicyStore) Update(arg0 context.Context, arg1 *v1.Policy, arg2 v10.UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockPolicyStoreMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPolicyStore)(nil).Update), arg0, arg1, arg2)
}
