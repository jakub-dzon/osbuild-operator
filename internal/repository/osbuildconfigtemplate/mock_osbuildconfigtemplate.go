// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-flotta/osbuild-operator/internal/repository/osbuildconfigtemplate (interfaces: Repository)

// Package osbuildconfigtemplate is a generated GoMock package.
package osbuildconfigtemplate

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/project-flotta/osbuild-operator/api/v1alpha1"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockRepository) Read(arg0 context.Context, arg1, arg2 string) (*v1alpha1.OSBuildConfigTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.OSBuildConfigTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRepositoryMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRepository)(nil).Read), arg0, arg1, arg2)
}