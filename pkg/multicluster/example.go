package multicluster

import (
	"context"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	skv2_corev1 "github.com/solo-io/skv2/pkg/api/kube/core/v1"
	"github.com/solo-io/skv2/pkg/api/kube/core/v1/controller"
	"go.uber.org/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

/*
Example of how the watcher could fit into an app setup flow.
*/
func example(local manager.Manager) {
	mccs := NewClientSet()

	go func() {
		err := RunClusterWatcher(context.TODO(), local, mccs, StartConfigMapReconcileLoop)
		if err != nil {
			panic("cluster watcher errored")
		}
	}()

	multiclusterClients := NewTypedClientSet(mccs)
	fooSet, err := multiclusterClients.Cluster("foo")
	if err != nil {
		// uh oh
	}

	fooSet.Secrets().DeleteAllOfSecret(context.TODO())

}

// User-provided reconcile loop starters
func StartConfigMapReconcileLoop(ctx context.Context, cluster string, mgr manager.Manager) error {
	go func() {
		err := controller.NewConfigMapReconcileLoop(cluster, mgr).RunConfigMapReconciler(ctx, *new(controller.ConfigMapReconciler))
		if err != nil {
			contextutils.LoggerFrom(ctx).DPanicw("ConfigMap reconcile loop stopped with error", zap.Error(err))
		}
	}()
	return nil
}

/**
Rough sketch of a typed multicluster clientset

Alternative is to have a structure like "setManager.Resource().Cluster().Action()", wdyt?
*/
// TODO generate

type typedClientSet interface {
	Cluster(cluster string) (skv2_corev1.Clientset, error)
}

type typedCs struct{ getter ClientSet }

func (m typedCs) Cluster(cluster string) (skv2_corev1.Clientset, error) {
	c, err := m.getter.Cluster(cluster)
	if err != nil {
		return nil, eris.Wrapf(err, "Failed to getManager client for cluster %v")
	}
	return skv2_corev1.NewClientset(c), nil
}

func NewTypedClientSet(getter ClientSet) typedClientSet {
	return typedCs{getter: getter}
}