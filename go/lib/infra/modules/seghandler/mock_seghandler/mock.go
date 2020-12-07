// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/lib/infra/modules/seghandler (interfaces: Storage,Verifier)

// Package mock_seghandler is a generated GoMock package.
package mock_seghandler

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	path_mgmt "github.com/scionproto/scion/go/lib/ctrl/path_mgmt"
	seg "github.com/scionproto/scion/go/lib/ctrl/seg"
	seghandler "github.com/scionproto/scion/go/lib/infra/modules/seghandler"
	segverifier "github.com/scionproto/scion/go/lib/infra/modules/segverifier"
	net "net"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// StoreRevs mocks base method
func (m *MockStorage) StoreRevs(arg0 context.Context, arg1 []*path_mgmt.SignedRevInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreRevs", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreRevs indicates an expected call of StoreRevs
func (mr *MockStorageMockRecorder) StoreRevs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreRevs", reflect.TypeOf((*MockStorage)(nil).StoreRevs), arg0, arg1)
}

// StoreSegs mocks base method
func (m *MockStorage) StoreSegs(arg0 context.Context, arg1 []*seg.Meta) (seghandler.SegStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSegs", arg0, arg1)
	ret0, _ := ret[0].(seghandler.SegStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreSegs indicates an expected call of StoreSegs
func (mr *MockStorageMockRecorder) StoreSegs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSegs", reflect.TypeOf((*MockStorage)(nil).StoreSegs), arg0, arg1)
}

// MockVerifier is a mock of Verifier interface
type MockVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockVerifierMockRecorder
}

// MockVerifierMockRecorder is the mock recorder for MockVerifier
type MockVerifierMockRecorder struct {
	mock *MockVerifier
}

// NewMockVerifier creates a new mock instance
func NewMockVerifier(ctrl *gomock.Controller) *MockVerifier {
	mock := &MockVerifier{ctrl: ctrl}
	mock.recorder = &MockVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVerifier) EXPECT() *MockVerifierMockRecorder {
	return m.recorder
}

// Verify mocks base method
func (m *MockVerifier) Verify(arg0 context.Context, arg1 seghandler.Segments, arg2 net.Addr) (chan segverifier.UnitResult, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1, arg2)
	ret0, _ := ret[0].(chan segverifier.UnitResult)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// Verify indicates an expected call of Verify
func (mr *MockVerifierMockRecorder) Verify(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockVerifier)(nil).Verify), arg0, arg1, arg2)
}