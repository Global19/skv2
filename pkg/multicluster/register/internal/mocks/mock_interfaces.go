// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_internal is a generated GoMock package.
package mock_internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClusterRBACBinder is a mock of ClusterRBACBinder interface.
type MockClusterRBACBinder struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRBACBinderMockRecorder
}

// MockClusterRBACBinderMockRecorder is the mock recorder for MockClusterRBACBinder.
type MockClusterRBACBinderMockRecorder struct {
	mock *MockClusterRBACBinder
}

// NewMockClusterRBACBinder creates a new mock instance.
func NewMockClusterRBACBinder(ctrl *gomock.Controller) *MockClusterRBACBinder {
	mock := &MockClusterRBACBinder{ctrl: ctrl}
	mock.recorder = &MockClusterRBACBinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRBACBinder) EXPECT() *MockClusterRBACBinderMockRecorder {
	return m.recorder
}

// BindClusterRoles mocks base method.
func (m *MockClusterRBACBinder) BindClusterRoles(ctx context.Context, serviceAccount client.ObjectKey, clusterRoles []client.ObjectKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindClusterRoles", ctx, serviceAccount, clusterRoles)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindClusterRoles indicates an expected call of BindClusterRoles.
func (mr *MockClusterRBACBinderMockRecorder) BindClusterRoles(ctx, serviceAccount, clusterRoles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindClusterRoles", reflect.TypeOf((*MockClusterRBACBinder)(nil).BindClusterRoles), ctx, serviceAccount, clusterRoles)
}

// BindRoles mocks base method.
func (m *MockClusterRBACBinder) BindRoles(ctx context.Context, serviceAccount client.ObjectKey, roles []client.ObjectKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindRoles", ctx, serviceAccount, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindRoles indicates an expected call of BindRoles.
func (mr *MockClusterRBACBinderMockRecorder) BindRoles(ctx, serviceAccount, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindRoles", reflect.TypeOf((*MockClusterRBACBinder)(nil).BindRoles), ctx, serviceAccount, roles)
}