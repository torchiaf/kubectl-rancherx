package rancher

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/torchiaf/kubectl-rancherx/pkg/manager"
	"k8s.io/client-go/rest"
)

func ListClusters(ctx context.Context, client *rest.RESTClient) (*v3.ClusterList, error) {

	clusters := &v3.ClusterList{}

	err := manager.List(ctx, client, "clusters", "", clusters)
	if err != nil {
		return &v3.ClusterList{}, err
	}

	return clusters, nil
}
