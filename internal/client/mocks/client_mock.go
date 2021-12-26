// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/omegion/s3-secret-manager/internal/client (interfaces: Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/omegion/s3-secret-manager/internal/api"
	controller "github.com/omegion/s3-secret-manager/internal/controller"
	secret "github.com/omegion/s3-secret-manager/pkg/secret"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// DeleteSecret mocks base method.
func (m *MockInterface) DeleteSecret(arg0 api.Interface, arg1 *secret.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockInterfaceMockRecorder) DeleteSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockInterface)(nil).DeleteSecret), arg0, arg1)
}

// GetS3API mocks base method.
func (m *MockInterface) GetS3API() (api.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetS3API")
	ret0, _ := ret[0].(api.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetS3API indicates an expected call of GetS3API.
func (mr *MockInterfaceMockRecorder) GetS3API() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetS3API", reflect.TypeOf((*MockInterface)(nil).GetS3API))
}

// GetSecret mocks base method.
func (m *MockInterface) GetSecret(arg0 api.Interface, arg1 *secret.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockInterfaceMockRecorder) GetSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockInterface)(nil).GetSecret), arg0, arg1)
}

// ListSecret mocks base method.
func (m *MockInterface) ListSecret(arg0 api.Interface, arg1 *controller.ListOptions) (*secret.Secrets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecret", arg0, arg1)
	ret0, _ := ret[0].(*secret.Secrets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecret indicates an expected call of ListSecret.
func (mr *MockInterfaceMockRecorder) ListSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecret", reflect.TypeOf((*MockInterface)(nil).ListSecret), arg0, arg1)
}

// ListVersions mocks base method.
func (m *MockInterface) ListVersions(arg0 api.Interface, arg1 *secret.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVersions", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListVersions indicates an expected call of ListVersions.
func (mr *MockInterfaceMockRecorder) ListVersions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVersions", reflect.TypeOf((*MockInterface)(nil).ListVersions), arg0, arg1)
}

// SetSecret mocks base method.
func (m *MockInterface) SetSecret(arg0 api.Interface, arg1 *secret.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSecret indicates an expected call of SetSecret.
func (mr *MockInterfaceMockRecorder) SetSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecret", reflect.TypeOf((*MockInterface)(nil).SetSecret), arg0, arg1)
}
