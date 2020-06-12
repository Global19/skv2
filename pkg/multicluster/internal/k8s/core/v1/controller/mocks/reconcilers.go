// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/skv2/pkg/multicluster/internal/k8s/core/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "k8s.io/api/core/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockSecretReconciler is a mock of SecretReconciler interface.
type MockSecretReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockSecretReconcilerMockRecorder
}

// MockSecretReconcilerMockRecorder is the mock recorder for MockSecretReconciler.
type MockSecretReconcilerMockRecorder struct {
	mock *MockSecretReconciler
}

// NewMockSecretReconciler creates a new mock instance.
func NewMockSecretReconciler(ctrl *gomock.Controller) *MockSecretReconciler {
	mock := &MockSecretReconciler{ctrl: ctrl}
	mock.recorder = &MockSecretReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretReconciler) EXPECT() *MockSecretReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSecret mocks base method.
func (m *MockSecretReconciler) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSecret", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSecret indicates an expected call of ReconcileSecret.
func (mr *MockSecretReconcilerMockRecorder) ReconcileSecret(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSecret", reflect.TypeOf((*MockSecretReconciler)(nil).ReconcileSecret), obj)
}

// MockSecretDeletionReconciler is a mock of SecretDeletionReconciler interface.
type MockSecretDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockSecretDeletionReconcilerMockRecorder
}

// MockSecretDeletionReconcilerMockRecorder is the mock recorder for MockSecretDeletionReconciler.
type MockSecretDeletionReconcilerMockRecorder struct {
	mock *MockSecretDeletionReconciler
}

// NewMockSecretDeletionReconciler creates a new mock instance.
func NewMockSecretDeletionReconciler(ctrl *gomock.Controller) *MockSecretDeletionReconciler {
	mock := &MockSecretDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockSecretDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretDeletionReconciler) EXPECT() *MockSecretDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSecretDeletion mocks base method.
func (m *MockSecretDeletionReconciler) ReconcileSecretDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileSecretDeletion", req)
}

// ReconcileSecretDeletion indicates an expected call of ReconcileSecretDeletion.
func (mr *MockSecretDeletionReconcilerMockRecorder) ReconcileSecretDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSecretDeletion", reflect.TypeOf((*MockSecretDeletionReconciler)(nil).ReconcileSecretDeletion), req)
}

// MockSecretFinalizer is a mock of SecretFinalizer interface.
type MockSecretFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockSecretFinalizerMockRecorder
}

// MockSecretFinalizerMockRecorder is the mock recorder for MockSecretFinalizer.
type MockSecretFinalizerMockRecorder struct {
	mock *MockSecretFinalizer
}

// NewMockSecretFinalizer creates a new mock instance.
func NewMockSecretFinalizer(ctrl *gomock.Controller) *MockSecretFinalizer {
	mock := &MockSecretFinalizer{ctrl: ctrl}
	mock.recorder = &MockSecretFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretFinalizer) EXPECT() *MockSecretFinalizerMockRecorder {
	return m.recorder
}

// ReconcileSecret mocks base method.
func (m *MockSecretFinalizer) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSecret", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSecret indicates an expected call of ReconcileSecret.
func (mr *MockSecretFinalizerMockRecorder) ReconcileSecret(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSecret", reflect.TypeOf((*MockSecretFinalizer)(nil).ReconcileSecret), obj)
}

// SecretFinalizerName mocks base method.
func (m *MockSecretFinalizer) SecretFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecretFinalizerName indicates an expected call of SecretFinalizerName.
func (mr *MockSecretFinalizerMockRecorder) SecretFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretFinalizerName", reflect.TypeOf((*MockSecretFinalizer)(nil).SecretFinalizerName))
}

// FinalizeSecret mocks base method.
func (m *MockSecretFinalizer) FinalizeSecret(obj *v1.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeSecret", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeSecret indicates an expected call of FinalizeSecret.
func (mr *MockSecretFinalizerMockRecorder) FinalizeSecret(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeSecret", reflect.TypeOf((*MockSecretFinalizer)(nil).FinalizeSecret), obj)
}

// MockSecretReconcileLoop is a mock of SecretReconcileLoop interface.
type MockSecretReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockSecretReconcileLoopMockRecorder
}

// MockSecretReconcileLoopMockRecorder is the mock recorder for MockSecretReconcileLoop.
type MockSecretReconcileLoopMockRecorder struct {
	mock *MockSecretReconcileLoop
}

// NewMockSecretReconcileLoop creates a new mock instance.
func NewMockSecretReconcileLoop(ctrl *gomock.Controller) *MockSecretReconcileLoop {
	mock := &MockSecretReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockSecretReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretReconcileLoop) EXPECT() *MockSecretReconcileLoopMockRecorder {
	return m.recorder
}

// RunSecretReconciler mocks base method.
func (m *MockSecretReconcileLoop) RunSecretReconciler(ctx context.Context, rec controller.SecretReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunSecretReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunSecretReconciler indicates an expected call of RunSecretReconciler.
func (mr *MockSecretReconcileLoopMockRecorder) RunSecretReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSecretReconciler", reflect.TypeOf((*MockSecretReconcileLoop)(nil).RunSecretReconciler), varargs...)
}

// MockServiceAccountReconciler is a mock of ServiceAccountReconciler interface.
type MockServiceAccountReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountReconcilerMockRecorder
}

// MockServiceAccountReconcilerMockRecorder is the mock recorder for MockServiceAccountReconciler.
type MockServiceAccountReconcilerMockRecorder struct {
	mock *MockServiceAccountReconciler
}

// NewMockServiceAccountReconciler creates a new mock instance.
func NewMockServiceAccountReconciler(ctrl *gomock.Controller) *MockServiceAccountReconciler {
	mock := &MockServiceAccountReconciler{ctrl: ctrl}
	mock.recorder = &MockServiceAccountReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountReconciler) EXPECT() *MockServiceAccountReconcilerMockRecorder {
	return m.recorder
}

// ReconcileServiceAccount mocks base method.
func (m *MockServiceAccountReconciler) ReconcileServiceAccount(obj *v1.ServiceAccount) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileServiceAccount", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileServiceAccount indicates an expected call of ReconcileServiceAccount.
func (mr *MockServiceAccountReconcilerMockRecorder) ReconcileServiceAccount(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileServiceAccount", reflect.TypeOf((*MockServiceAccountReconciler)(nil).ReconcileServiceAccount), obj)
}

// MockServiceAccountDeletionReconciler is a mock of ServiceAccountDeletionReconciler interface.
type MockServiceAccountDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountDeletionReconcilerMockRecorder
}

// MockServiceAccountDeletionReconcilerMockRecorder is the mock recorder for MockServiceAccountDeletionReconciler.
type MockServiceAccountDeletionReconcilerMockRecorder struct {
	mock *MockServiceAccountDeletionReconciler
}

// NewMockServiceAccountDeletionReconciler creates a new mock instance.
func NewMockServiceAccountDeletionReconciler(ctrl *gomock.Controller) *MockServiceAccountDeletionReconciler {
	mock := &MockServiceAccountDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockServiceAccountDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountDeletionReconciler) EXPECT() *MockServiceAccountDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileServiceAccountDeletion mocks base method.
func (m *MockServiceAccountDeletionReconciler) ReconcileServiceAccountDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileServiceAccountDeletion", req)
}

// ReconcileServiceAccountDeletion indicates an expected call of ReconcileServiceAccountDeletion.
func (mr *MockServiceAccountDeletionReconcilerMockRecorder) ReconcileServiceAccountDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileServiceAccountDeletion", reflect.TypeOf((*MockServiceAccountDeletionReconciler)(nil).ReconcileServiceAccountDeletion), req)
}

// MockServiceAccountFinalizer is a mock of ServiceAccountFinalizer interface.
type MockServiceAccountFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountFinalizerMockRecorder
}

// MockServiceAccountFinalizerMockRecorder is the mock recorder for MockServiceAccountFinalizer.
type MockServiceAccountFinalizerMockRecorder struct {
	mock *MockServiceAccountFinalizer
}

// NewMockServiceAccountFinalizer creates a new mock instance.
func NewMockServiceAccountFinalizer(ctrl *gomock.Controller) *MockServiceAccountFinalizer {
	mock := &MockServiceAccountFinalizer{ctrl: ctrl}
	mock.recorder = &MockServiceAccountFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountFinalizer) EXPECT() *MockServiceAccountFinalizerMockRecorder {
	return m.recorder
}

// ReconcileServiceAccount mocks base method.
func (m *MockServiceAccountFinalizer) ReconcileServiceAccount(obj *v1.ServiceAccount) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileServiceAccount", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileServiceAccount indicates an expected call of ReconcileServiceAccount.
func (mr *MockServiceAccountFinalizerMockRecorder) ReconcileServiceAccount(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileServiceAccount", reflect.TypeOf((*MockServiceAccountFinalizer)(nil).ReconcileServiceAccount), obj)
}

// ServiceAccountFinalizerName mocks base method.
func (m *MockServiceAccountFinalizer) ServiceAccountFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceAccountFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ServiceAccountFinalizerName indicates an expected call of ServiceAccountFinalizerName.
func (mr *MockServiceAccountFinalizerMockRecorder) ServiceAccountFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceAccountFinalizerName", reflect.TypeOf((*MockServiceAccountFinalizer)(nil).ServiceAccountFinalizerName))
}

// FinalizeServiceAccount mocks base method.
func (m *MockServiceAccountFinalizer) FinalizeServiceAccount(obj *v1.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeServiceAccount", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeServiceAccount indicates an expected call of FinalizeServiceAccount.
func (mr *MockServiceAccountFinalizerMockRecorder) FinalizeServiceAccount(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeServiceAccount", reflect.TypeOf((*MockServiceAccountFinalizer)(nil).FinalizeServiceAccount), obj)
}

// MockServiceAccountReconcileLoop is a mock of ServiceAccountReconcileLoop interface.
type MockServiceAccountReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountReconcileLoopMockRecorder
}

// MockServiceAccountReconcileLoopMockRecorder is the mock recorder for MockServiceAccountReconcileLoop.
type MockServiceAccountReconcileLoopMockRecorder struct {
	mock *MockServiceAccountReconcileLoop
}

// NewMockServiceAccountReconcileLoop creates a new mock instance.
func NewMockServiceAccountReconcileLoop(ctrl *gomock.Controller) *MockServiceAccountReconcileLoop {
	mock := &MockServiceAccountReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockServiceAccountReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountReconcileLoop) EXPECT() *MockServiceAccountReconcileLoopMockRecorder {
	return m.recorder
}

// RunServiceAccountReconciler mocks base method.
func (m *MockServiceAccountReconcileLoop) RunServiceAccountReconciler(ctx context.Context, rec controller.ServiceAccountReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunServiceAccountReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunServiceAccountReconciler indicates an expected call of RunServiceAccountReconciler.
func (mr *MockServiceAccountReconcileLoopMockRecorder) RunServiceAccountReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunServiceAccountReconciler", reflect.TypeOf((*MockServiceAccountReconcileLoop)(nil).RunServiceAccountReconciler), varargs...)
}

// MockNamespaceReconciler is a mock of NamespaceReconciler interface.
type MockNamespaceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceReconcilerMockRecorder
}

// MockNamespaceReconcilerMockRecorder is the mock recorder for MockNamespaceReconciler.
type MockNamespaceReconcilerMockRecorder struct {
	mock *MockNamespaceReconciler
}

// NewMockNamespaceReconciler creates a new mock instance.
func NewMockNamespaceReconciler(ctrl *gomock.Controller) *MockNamespaceReconciler {
	mock := &MockNamespaceReconciler{ctrl: ctrl}
	mock.recorder = &MockNamespaceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceReconciler) EXPECT() *MockNamespaceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileNamespace mocks base method.
func (m *MockNamespaceReconciler) ReconcileNamespace(obj *v1.Namespace) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileNamespace", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileNamespace indicates an expected call of ReconcileNamespace.
func (mr *MockNamespaceReconcilerMockRecorder) ReconcileNamespace(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNamespace", reflect.TypeOf((*MockNamespaceReconciler)(nil).ReconcileNamespace), obj)
}

// MockNamespaceDeletionReconciler is a mock of NamespaceDeletionReconciler interface.
type MockNamespaceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceDeletionReconcilerMockRecorder
}

// MockNamespaceDeletionReconcilerMockRecorder is the mock recorder for MockNamespaceDeletionReconciler.
type MockNamespaceDeletionReconcilerMockRecorder struct {
	mock *MockNamespaceDeletionReconciler
}

// NewMockNamespaceDeletionReconciler creates a new mock instance.
func NewMockNamespaceDeletionReconciler(ctrl *gomock.Controller) *MockNamespaceDeletionReconciler {
	mock := &MockNamespaceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockNamespaceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceDeletionReconciler) EXPECT() *MockNamespaceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileNamespaceDeletion mocks base method.
func (m *MockNamespaceDeletionReconciler) ReconcileNamespaceDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileNamespaceDeletion", req)
}

// ReconcileNamespaceDeletion indicates an expected call of ReconcileNamespaceDeletion.
func (mr *MockNamespaceDeletionReconcilerMockRecorder) ReconcileNamespaceDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNamespaceDeletion", reflect.TypeOf((*MockNamespaceDeletionReconciler)(nil).ReconcileNamespaceDeletion), req)
}

// MockNamespaceFinalizer is a mock of NamespaceFinalizer interface.
type MockNamespaceFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceFinalizerMockRecorder
}

// MockNamespaceFinalizerMockRecorder is the mock recorder for MockNamespaceFinalizer.
type MockNamespaceFinalizerMockRecorder struct {
	mock *MockNamespaceFinalizer
}

// NewMockNamespaceFinalizer creates a new mock instance.
func NewMockNamespaceFinalizer(ctrl *gomock.Controller) *MockNamespaceFinalizer {
	mock := &MockNamespaceFinalizer{ctrl: ctrl}
	mock.recorder = &MockNamespaceFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceFinalizer) EXPECT() *MockNamespaceFinalizerMockRecorder {
	return m.recorder
}

// ReconcileNamespace mocks base method.
func (m *MockNamespaceFinalizer) ReconcileNamespace(obj *v1.Namespace) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileNamespace", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileNamespace indicates an expected call of ReconcileNamespace.
func (mr *MockNamespaceFinalizerMockRecorder) ReconcileNamespace(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNamespace", reflect.TypeOf((*MockNamespaceFinalizer)(nil).ReconcileNamespace), obj)
}

// NamespaceFinalizerName mocks base method.
func (m *MockNamespaceFinalizer) NamespaceFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamespaceFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// NamespaceFinalizerName indicates an expected call of NamespaceFinalizerName.
func (mr *MockNamespaceFinalizerMockRecorder) NamespaceFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamespaceFinalizerName", reflect.TypeOf((*MockNamespaceFinalizer)(nil).NamespaceFinalizerName))
}

// FinalizeNamespace mocks base method.
func (m *MockNamespaceFinalizer) FinalizeNamespace(obj *v1.Namespace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeNamespace", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeNamespace indicates an expected call of FinalizeNamespace.
func (mr *MockNamespaceFinalizerMockRecorder) FinalizeNamespace(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeNamespace", reflect.TypeOf((*MockNamespaceFinalizer)(nil).FinalizeNamespace), obj)
}

// MockNamespaceReconcileLoop is a mock of NamespaceReconcileLoop interface.
type MockNamespaceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceReconcileLoopMockRecorder
}

// MockNamespaceReconcileLoopMockRecorder is the mock recorder for MockNamespaceReconcileLoop.
type MockNamespaceReconcileLoopMockRecorder struct {
	mock *MockNamespaceReconcileLoop
}

// NewMockNamespaceReconcileLoop creates a new mock instance.
func NewMockNamespaceReconcileLoop(ctrl *gomock.Controller) *MockNamespaceReconcileLoop {
	mock := &MockNamespaceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockNamespaceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceReconcileLoop) EXPECT() *MockNamespaceReconcileLoopMockRecorder {
	return m.recorder
}

// RunNamespaceReconciler mocks base method.
func (m *MockNamespaceReconcileLoop) RunNamespaceReconciler(ctx context.Context, rec controller.NamespaceReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunNamespaceReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunNamespaceReconciler indicates an expected call of RunNamespaceReconciler.
func (mr *MockNamespaceReconcileLoopMockRecorder) RunNamespaceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunNamespaceReconciler", reflect.TypeOf((*MockNamespaceReconcileLoop)(nil).RunNamespaceReconciler), varargs...)
}