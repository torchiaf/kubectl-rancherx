package rancher

import (
	"context"

	apiv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/torchiaf/kubectl-rancherx/pkg/manager"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func CreateProject(ctx context.Context, client *rest.RESTClient, name string, displayName string, clusterName string) error {

	obj := &apiv3.Project{
		ObjectMeta: v1.ObjectMeta{
			Name: name,
		},
		// TODO generate name -> hide the namespace on the UI
		// ObjectMeta: v1.ObjectMeta{
		// 	GenerateName: "p-",
		// },
		Spec: apiv3.ProjectSpec{
			ClusterName: clusterName,
			DisplayName: displayName,
		},
	}

	return manager.Create(ctx, client, "projects", clusterName, obj)
}
