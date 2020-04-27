package register

import (
	"context"

	"github.com/aws/aws-sdk-go/service/eks"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
	"k8s.io/client-go/tools/clientcmd"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

type GkeConfigBuilder interface {
	ConfigForCluster(ctx context.Context, cluster *containerpb.Cluster) (clientcmd.ClientConfig, error)
}

type EksConfigBuilder interface {
	ConfigForCluster(ctx context.Context, cluster *eks.Cluster) (clientcmd.ClientConfig, error)
}

type ClusterRegistrant interface {
	RegisterCluster(
		ctx context.Context,
		remoteCfg, remoteCtx, clusterName, namespace string,
	) error
	RegisterClusterFromConfig(
		ctx context.Context,
		config clientcmd.ClientConfig,
		clusterName, namespace string,
	) error
}
