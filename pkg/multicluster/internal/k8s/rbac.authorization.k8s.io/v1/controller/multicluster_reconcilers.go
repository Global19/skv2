// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	rbac_authorization_k8s_io_v1 "k8s.io/api/rbac/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Role Resource across clusters.
// implemented by the user
type MulticlusterRoleReconciler interface {
	ReconcileRole(clusterName string, obj *rbac_authorization_k8s_io_v1.Role) (reconcile.Result, error)
}

// Reconcile deletion events for the Role Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterRoleDeletionReconciler interface {
	ReconcileRoleDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterRoleReconcilerFuncs struct {
	OnReconcileRole         func(clusterName string, obj *rbac_authorization_k8s_io_v1.Role) (reconcile.Result, error)
	OnReconcileRoleDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterRoleReconcilerFuncs) ReconcileRole(clusterName string, obj *rbac_authorization_k8s_io_v1.Role) (reconcile.Result, error) {
	if f.OnReconcileRole == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileRole(clusterName, obj)
}

func (f *MulticlusterRoleReconcilerFuncs) ReconcileRoleDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileRoleDeletion == nil {
		return
	}
	f.OnReconcileRoleDeletion(clusterName, req)
}

type MulticlusterRoleReconcileLoop interface {
	// AddMulticlusterRoleReconciler adds a MulticlusterRoleReconciler to the MulticlusterRoleReconcileLoop.
	AddMulticlusterRoleReconciler(ctx context.Context, rec MulticlusterRoleReconciler, predicates ...predicate.Predicate)
}

type multiclusterRoleReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterRoleReconcileLoop) AddMulticlusterRoleReconciler(ctx context.Context, rec MulticlusterRoleReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericRoleMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterRoleReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterRoleReconcileLoop {
	return &multiclusterRoleReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &rbac_authorization_k8s_io_v1.Role{})}
}

type genericRoleMulticlusterReconciler struct {
	reconciler MulticlusterRoleReconciler
}

func (g genericRoleMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterRoleDeletionReconciler); ok {
		deletionReconciler.ReconcileRoleDeletion(cluster, req)
	}
}

func (g genericRoleMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.Role)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Role handler received event for %T", object)
	}
	return g.reconciler.ReconcileRole(cluster, obj)
}

// Reconcile Upsert events for the RoleBinding Resource across clusters.
// implemented by the user
type MulticlusterRoleBindingReconciler interface {
	ReconcileRoleBinding(clusterName string, obj *rbac_authorization_k8s_io_v1.RoleBinding) (reconcile.Result, error)
}

// Reconcile deletion events for the RoleBinding Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterRoleBindingDeletionReconciler interface {
	ReconcileRoleBindingDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterRoleBindingReconcilerFuncs struct {
	OnReconcileRoleBinding         func(clusterName string, obj *rbac_authorization_k8s_io_v1.RoleBinding) (reconcile.Result, error)
	OnReconcileRoleBindingDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterRoleBindingReconcilerFuncs) ReconcileRoleBinding(clusterName string, obj *rbac_authorization_k8s_io_v1.RoleBinding) (reconcile.Result, error) {
	if f.OnReconcileRoleBinding == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileRoleBinding(clusterName, obj)
}

func (f *MulticlusterRoleBindingReconcilerFuncs) ReconcileRoleBindingDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileRoleBindingDeletion == nil {
		return
	}
	f.OnReconcileRoleBindingDeletion(clusterName, req)
}

type MulticlusterRoleBindingReconcileLoop interface {
	// AddMulticlusterRoleBindingReconciler adds a MulticlusterRoleBindingReconciler to the MulticlusterRoleBindingReconcileLoop.
	AddMulticlusterRoleBindingReconciler(ctx context.Context, rec MulticlusterRoleBindingReconciler, predicates ...predicate.Predicate)
}

type multiclusterRoleBindingReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterRoleBindingReconcileLoop) AddMulticlusterRoleBindingReconciler(ctx context.Context, rec MulticlusterRoleBindingReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericRoleBindingMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterRoleBindingReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterRoleBindingReconcileLoop {
	return &multiclusterRoleBindingReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &rbac_authorization_k8s_io_v1.RoleBinding{})}
}

type genericRoleBindingMulticlusterReconciler struct {
	reconciler MulticlusterRoleBindingReconciler
}

func (g genericRoleBindingMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterRoleBindingDeletionReconciler); ok {
		deletionReconciler.ReconcileRoleBindingDeletion(cluster, req)
	}
}

func (g genericRoleBindingMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.RoleBinding)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: RoleBinding handler received event for %T", object)
	}
	return g.reconciler.ReconcileRoleBinding(cluster, obj)
}

// Reconcile Upsert events for the ClusterRole Resource across clusters.
// implemented by the user
type MulticlusterClusterRoleReconciler interface {
	ReconcileClusterRole(clusterName string, obj *rbac_authorization_k8s_io_v1.ClusterRole) (reconcile.Result, error)
}

// Reconcile deletion events for the ClusterRole Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterClusterRoleDeletionReconciler interface {
	ReconcileClusterRoleDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterClusterRoleReconcilerFuncs struct {
	OnReconcileClusterRole         func(clusterName string, obj *rbac_authorization_k8s_io_v1.ClusterRole) (reconcile.Result, error)
	OnReconcileClusterRoleDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterClusterRoleReconcilerFuncs) ReconcileClusterRole(clusterName string, obj *rbac_authorization_k8s_io_v1.ClusterRole) (reconcile.Result, error) {
	if f.OnReconcileClusterRole == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileClusterRole(clusterName, obj)
}

func (f *MulticlusterClusterRoleReconcilerFuncs) ReconcileClusterRoleDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileClusterRoleDeletion == nil {
		return
	}
	f.OnReconcileClusterRoleDeletion(clusterName, req)
}

type MulticlusterClusterRoleReconcileLoop interface {
	// AddMulticlusterClusterRoleReconciler adds a MulticlusterClusterRoleReconciler to the MulticlusterClusterRoleReconcileLoop.
	AddMulticlusterClusterRoleReconciler(ctx context.Context, rec MulticlusterClusterRoleReconciler, predicates ...predicate.Predicate)
}

type multiclusterClusterRoleReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterClusterRoleReconcileLoop) AddMulticlusterClusterRoleReconciler(ctx context.Context, rec MulticlusterClusterRoleReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericClusterRoleMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterClusterRoleReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterClusterRoleReconcileLoop {
	return &multiclusterClusterRoleReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &rbac_authorization_k8s_io_v1.ClusterRole{})}
}

type genericClusterRoleMulticlusterReconciler struct {
	reconciler MulticlusterClusterRoleReconciler
}

func (g genericClusterRoleMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterClusterRoleDeletionReconciler); ok {
		deletionReconciler.ReconcileClusterRoleDeletion(cluster, req)
	}
}

func (g genericClusterRoleMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.ClusterRole)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ClusterRole handler received event for %T", object)
	}
	return g.reconciler.ReconcileClusterRole(cluster, obj)
}

// Reconcile Upsert events for the ClusterRoleBinding Resource across clusters.
// implemented by the user
type MulticlusterClusterRoleBindingReconciler interface {
	ReconcileClusterRoleBinding(clusterName string, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) (reconcile.Result, error)
}

// Reconcile deletion events for the ClusterRoleBinding Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterClusterRoleBindingDeletionReconciler interface {
	ReconcileClusterRoleBindingDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterClusterRoleBindingReconcilerFuncs struct {
	OnReconcileClusterRoleBinding         func(clusterName string, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) (reconcile.Result, error)
	OnReconcileClusterRoleBindingDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterClusterRoleBindingReconcilerFuncs) ReconcileClusterRoleBinding(clusterName string, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding) (reconcile.Result, error) {
	if f.OnReconcileClusterRoleBinding == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileClusterRoleBinding(clusterName, obj)
}

func (f *MulticlusterClusterRoleBindingReconcilerFuncs) ReconcileClusterRoleBindingDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileClusterRoleBindingDeletion == nil {
		return
	}
	f.OnReconcileClusterRoleBindingDeletion(clusterName, req)
}

type MulticlusterClusterRoleBindingReconcileLoop interface {
	// AddMulticlusterClusterRoleBindingReconciler adds a MulticlusterClusterRoleBindingReconciler to the MulticlusterClusterRoleBindingReconcileLoop.
	AddMulticlusterClusterRoleBindingReconciler(ctx context.Context, rec MulticlusterClusterRoleBindingReconciler, predicates ...predicate.Predicate)
}

type multiclusterClusterRoleBindingReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterClusterRoleBindingReconcileLoop) AddMulticlusterClusterRoleBindingReconciler(ctx context.Context, rec MulticlusterClusterRoleBindingReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericClusterRoleBindingMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterClusterRoleBindingReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterClusterRoleBindingReconcileLoop {
	return &multiclusterClusterRoleBindingReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &rbac_authorization_k8s_io_v1.ClusterRoleBinding{})}
}

type genericClusterRoleBindingMulticlusterReconciler struct {
	reconciler MulticlusterClusterRoleBindingReconciler
}

func (g genericClusterRoleBindingMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterClusterRoleBindingDeletionReconciler); ok {
		deletionReconciler.ReconcileClusterRoleBindingDeletion(cluster, req)
	}
}

func (g genericClusterRoleBindingMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ClusterRoleBinding handler received event for %T", object)
	}
	return g.reconciler.ReconcileClusterRoleBinding(cluster, obj)
}