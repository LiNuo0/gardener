// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/client/kubernetes (interfaces: ChartApplier)
//
// Generated by this command:
//
//	mockgen -package mock -destination=mocks_chartapplier.go github.com/gardener/gardener/pkg/client/kubernetes ChartApplier
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	embed "embed"
	reflect "reflect"

	chartrenderer "github.com/gardener/gardener/pkg/chartrenderer"
	kubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	gomock "go.uber.org/mock/gomock"
)

// MockChartApplier is a mock of ChartApplier interface.
type MockChartApplier struct {
	ctrl     *gomock.Controller
	recorder *MockChartApplierMockRecorder
}

// MockChartApplierMockRecorder is the mock recorder for MockChartApplier.
type MockChartApplierMockRecorder struct {
	mock *MockChartApplier
}

// NewMockChartApplier creates a new mock instance.
func NewMockChartApplier(ctrl *gomock.Controller) *MockChartApplier {
	mock := &MockChartApplier{ctrl: ctrl}
	mock.recorder = &MockChartApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChartApplier) EXPECT() *MockChartApplierMockRecorder {
	return m.recorder
}

// ApplyFromArchive mocks base method.
func (m *MockChartApplier) ApplyFromArchive(arg0 context.Context, arg1 []byte, arg2, arg3 string, arg4 ...kubernetes.ApplyOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApplyFromArchive", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyFromArchive indicates an expected call of ApplyFromArchive.
func (mr *MockChartApplierMockRecorder) ApplyFromArchive(arg0, arg1, arg2, arg3 any, arg4 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyFromArchive", reflect.TypeOf((*MockChartApplier)(nil).ApplyFromArchive), varargs...)
}

// ApplyFromEmbeddedFS mocks base method.
func (m *MockChartApplier) ApplyFromEmbeddedFS(arg0 context.Context, arg1 embed.FS, arg2, arg3, arg4 string, arg5 ...kubernetes.ApplyOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApplyFromEmbeddedFS", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyFromEmbeddedFS indicates an expected call of ApplyFromEmbeddedFS.
func (mr *MockChartApplierMockRecorder) ApplyFromEmbeddedFS(arg0, arg1, arg2, arg3, arg4 any, arg5 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyFromEmbeddedFS", reflect.TypeOf((*MockChartApplier)(nil).ApplyFromEmbeddedFS), varargs...)
}

// DeleteFromArchive mocks base method.
func (m *MockChartApplier) DeleteFromArchive(arg0 context.Context, arg1 []byte, arg2, arg3 string, arg4 ...kubernetes.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFromArchive", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFromArchive indicates an expected call of DeleteFromArchive.
func (mr *MockChartApplierMockRecorder) DeleteFromArchive(arg0, arg1, arg2, arg3 any, arg4 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromArchive", reflect.TypeOf((*MockChartApplier)(nil).DeleteFromArchive), varargs...)
}

// DeleteFromEmbeddedFS mocks base method.
func (m *MockChartApplier) DeleteFromEmbeddedFS(arg0 context.Context, arg1 embed.FS, arg2, arg3, arg4 string, arg5 ...kubernetes.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFromEmbeddedFS", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFromEmbeddedFS indicates an expected call of DeleteFromEmbeddedFS.
func (mr *MockChartApplierMockRecorder) DeleteFromEmbeddedFS(arg0, arg1, arg2, arg3, arg4 any, arg5 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromEmbeddedFS", reflect.TypeOf((*MockChartApplier)(nil).DeleteFromEmbeddedFS), varargs...)
}

// RenderArchive mocks base method.
func (m *MockChartApplier) RenderArchive(arg0 []byte, arg1, arg2 string, arg3 any) (*chartrenderer.RenderedChart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderArchive", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*chartrenderer.RenderedChart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenderArchive indicates an expected call of RenderArchive.
func (mr *MockChartApplierMockRecorder) RenderArchive(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderArchive", reflect.TypeOf((*MockChartApplier)(nil).RenderArchive), arg0, arg1, arg2, arg3)
}

// RenderEmbeddedFS mocks base method.
func (m *MockChartApplier) RenderEmbeddedFS(arg0 embed.FS, arg1, arg2, arg3 string, arg4 any) (*chartrenderer.RenderedChart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderEmbeddedFS", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*chartrenderer.RenderedChart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenderEmbeddedFS indicates an expected call of RenderEmbeddedFS.
func (mr *MockChartApplierMockRecorder) RenderEmbeddedFS(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderEmbeddedFS", reflect.TypeOf((*MockChartApplier)(nil).RenderEmbeddedFS), arg0, arg1, arg2, arg3, arg4)
}
