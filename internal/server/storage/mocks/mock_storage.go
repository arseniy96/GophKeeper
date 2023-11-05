// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/server/handlers/server.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	storage "github.com/arseniy96/GophKeeper/internal/server/storage"
	gomock "github.com/golang/mock/gomock"
)

// Mockrepository is a mock of repository interface.
type Mockrepository struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryMockRecorder
}

// MockrepositoryMockRecorder is the mock recorder for Mockrepository.
type MockrepositoryMockRecorder struct {
	mock *Mockrepository
}

// NewMockrepository creates a new mock instance.
func NewMockrepository(ctrl *gomock.Controller) *Mockrepository {
	mock := &Mockrepository{ctrl: ctrl}
	mock.recorder = &MockrepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepository) EXPECT() *MockrepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *Mockrepository) CreateUser(ctx context.Context, login, password string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, login, password)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockrepositoryMockRecorder) CreateUser(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*Mockrepository)(nil).CreateUser), ctx, login, password)
}

// FindUserByLogin mocks base method.
func (m *Mockrepository) FindUserByLogin(ctx context.Context, login string) (*storage.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByLogin", ctx, login)
	ret0, _ := ret[0].(*storage.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByLogin indicates an expected call of FindUserByLogin.
func (mr *MockrepositoryMockRecorder) FindUserByLogin(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByLogin", reflect.TypeOf((*Mockrepository)(nil).FindUserByLogin), ctx, login)
}

// FindUserRecord mocks base method.
func (m *Mockrepository) FindUserRecord(ctx context.Context, id, userID int64) (*storage.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserRecord", ctx, id, userID)
	ret0, _ := ret[0].(*storage.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserRecord indicates an expected call of FindUserRecord.
func (mr *MockrepositoryMockRecorder) FindUserRecord(ctx, id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserRecord", reflect.TypeOf((*Mockrepository)(nil).FindUserRecord), ctx, id, userID)
}

// GetUserData mocks base method.
func (m *Mockrepository) GetUserData(ctx context.Context, userID int64) ([]storage.ShortRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserData", ctx, userID)
	ret0, _ := ret[0].([]storage.ShortRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserData indicates an expected call of GetUserData.
func (mr *MockrepositoryMockRecorder) GetUserData(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserData", reflect.TypeOf((*Mockrepository)(nil).GetUserData), ctx, userID)
}

// HealthCheck mocks base method.
func (m *Mockrepository) HealthCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockrepositoryMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*Mockrepository)(nil).HealthCheck))
}

// SaveUserData mocks base method.
func (m *Mockrepository) SaveUserData(ctx context.Context, userID int64, name, dataType string, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUserData", ctx, userID, name, dataType, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUserData indicates an expected call of SaveUserData.
func (mr *MockrepositoryMockRecorder) SaveUserData(ctx, userID, name, dataType, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUserData", reflect.TypeOf((*Mockrepository)(nil).SaveUserData), ctx, userID, name, dataType, data)
}

// UpdateUserRecord mocks base method.
func (m *Mockrepository) UpdateUserRecord(ctx context.Context, record *storage.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserRecord", ctx, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserRecord indicates an expected call of UpdateUserRecord.
func (mr *MockrepositoryMockRecorder) UpdateUserRecord(ctx, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserRecord", reflect.TypeOf((*Mockrepository)(nil).UpdateUserRecord), ctx, record)
}

// Mockcrypt is a mock of crypt interface.
type Mockcrypt struct {
	ctrl     *gomock.Controller
	recorder *MockcryptMockRecorder
}

// MockcryptMockRecorder is the mock recorder for Mockcrypt.
type MockcryptMockRecorder struct {
	mock *Mockcrypt
}

// NewMockcrypt creates a new mock instance.
func NewMockcrypt(ctrl *gomock.Controller) *Mockcrypt {
	mock := &Mockcrypt{ctrl: ctrl}
	mock.recorder = &MockcryptMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockcrypt) EXPECT() *MockcryptMockRecorder {
	return m.recorder
}

// BuildJWT mocks base method.
func (m *Mockcrypt) BuildJWT(userID int64, secret string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildJWT", userID, secret)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildJWT indicates an expected call of BuildJWT.
func (mr *MockcryptMockRecorder) BuildJWT(userID, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildJWT", reflect.TypeOf((*Mockcrypt)(nil).BuildJWT), userID, secret)
}

// CompareHash mocks base method.
func (m *Mockcrypt) CompareHash(src, hash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareHash", src, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompareHash indicates an expected call of CompareHash.
func (mr *MockcryptMockRecorder) CompareHash(src, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareHash", reflect.TypeOf((*Mockcrypt)(nil).CompareHash), src, hash)
}

// GetUserID mocks base method.
func (m *Mockcrypt) GetUserID(tokenString, secret string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserID", tokenString, secret)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserID indicates an expected call of GetUserID.
func (mr *MockcryptMockRecorder) GetUserID(tokenString, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*Mockcrypt)(nil).GetUserID), tokenString, secret)
}

// HashFunc mocks base method.
func (m *Mockcrypt) HashFunc(src string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashFunc", src)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashFunc indicates an expected call of HashFunc.
func (mr *MockcryptMockRecorder) HashFunc(src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashFunc", reflect.TypeOf((*Mockcrypt)(nil).HashFunc), src)
}
