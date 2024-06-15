// Code generated by MockGen. DO NOT EDIT.
// Source: db/sqlc/store.go
//
// Generated by this command:
//
//	mockgen -package mock_db -source=db/sqlc/store.go -destination=db/mock/store.go
//

// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	reflect "reflect"
	db "widgets/db/sqlc"

	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateCounter mocks base method.
func (m *MockStore) CreateCounter(ctx context.Context, arg db.CreateCounterParams) (db.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCounter", ctx, arg)
	ret0, _ := ret[0].(db.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCounter indicates an expected call of CreateCounter.
func (mr *MockStoreMockRecorder) CreateCounter(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCounter", reflect.TypeOf((*MockStore)(nil).CreateCounter), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), ctx, arg)
}

// CreateWidget mocks base method.
func (m *MockStore) CreateWidget(ctx context.Context, arg db.CreateWidgetParams) (db.Widget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWidget", ctx, arg)
	ret0, _ := ret[0].(db.Widget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWidget indicates an expected call of CreateWidget.
func (mr *MockStoreMockRecorder) CreateWidget(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWidget", reflect.TypeOf((*MockStore)(nil).CreateWidget), ctx, arg)
}

// DeleteCounter mocks base method.
func (m *MockStore) DeleteCounter(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCounter", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCounter indicates an expected call of DeleteCounter.
func (mr *MockStoreMockRecorder) DeleteCounter(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCounter", reflect.TypeOf((*MockStore)(nil).DeleteCounter), ctx, id)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), ctx, id)
}

// DeleteWidget mocks base method.
func (m *MockStore) DeleteWidget(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWidget", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWidget indicates an expected call of DeleteWidget.
func (mr *MockStoreMockRecorder) DeleteWidget(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWidget", reflect.TypeOf((*MockStore)(nil).DeleteWidget), ctx, id)
}

// GetCounter mocks base method.
func (m *MockStore) GetCounter(ctx context.Context, id int64) (db.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounter", ctx, id)
	ret0, _ := ret[0].(db.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCounter indicates an expected call of GetCounter.
func (mr *MockStoreMockRecorder) GetCounter(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounter", reflect.TypeOf((*MockStore)(nil).GetCounter), ctx, id)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(ctx context.Context, id int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), ctx, id)
}

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(ctx context.Context, username string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, username)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), ctx, username)
}

// GetWidget mocks base method.
func (m *MockStore) GetWidget(ctx context.Context, id int64) (db.Widget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWidget", ctx, id)
	ret0, _ := ret[0].(db.Widget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWidget indicates an expected call of GetWidget.
func (mr *MockStoreMockRecorder) GetWidget(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWidget", reflect.TypeOf((*MockStore)(nil).GetWidget), ctx, id)
}

// ListCounters mocks base method.
func (m *MockStore) ListCounters(ctx context.Context, arg db.ListCountersParams) ([]db.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCounters", ctx, arg)
	ret0, _ := ret[0].([]db.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCounters indicates an expected call of ListCounters.
func (mr *MockStoreMockRecorder) ListCounters(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCounters", reflect.TypeOf((*MockStore)(nil).ListCounters), ctx, arg)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, arg)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), ctx, arg)
}

// ListWidgets mocks base method.
func (m *MockStore) ListWidgets(ctx context.Context, arg db.ListWidgetsParams) ([]db.Widget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWidgets", ctx, arg)
	ret0, _ := ret[0].([]db.Widget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWidgets indicates an expected call of ListWidgets.
func (mr *MockStoreMockRecorder) ListWidgets(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWidgets", reflect.TypeOf((*MockStore)(nil).ListWidgets), ctx, arg)
}

// UpdateCounter mocks base method.
func (m *MockStore) UpdateCounter(ctx context.Context, arg db.UpdateCounterParams) (db.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCounter", ctx, arg)
	ret0, _ := ret[0].(db.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCounter indicates an expected call of UpdateCounter.
func (mr *MockStoreMockRecorder) UpdateCounter(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCounter", reflect.TypeOf((*MockStore)(nil).UpdateCounter), ctx, arg)
}

// UpdateCounterValue mocks base method.
func (m *MockStore) UpdateCounterValue(ctx context.Context, arg db.UpdateCounterValueParams) (db.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCounterValue", ctx, arg)
	ret0, _ := ret[0].(db.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCounterValue indicates an expected call of UpdateCounterValue.
func (mr *MockStoreMockRecorder) UpdateCounterValue(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCounterValue", reflect.TypeOf((*MockStore)(nil).UpdateCounterValue), ctx, arg)
}

// UpdateUserPassword mocks base method.
func (m *MockStore) UpdateUserPassword(ctx context.Context, arg db.UpdateUserPasswordParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", ctx, arg)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockStoreMockRecorder) UpdateUserPassword(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockStore)(nil).UpdateUserPassword), ctx, arg)
}

// UpdateWidget mocks base method.
func (m *MockStore) UpdateWidget(ctx context.Context, arg db.UpdateWidgetParams) (db.Widget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWidget", ctx, arg)
	ret0, _ := ret[0].(db.Widget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWidget indicates an expected call of UpdateWidget.
func (mr *MockStoreMockRecorder) UpdateWidget(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWidget", reflect.TypeOf((*MockStore)(nil).UpdateWidget), ctx, arg)
}
