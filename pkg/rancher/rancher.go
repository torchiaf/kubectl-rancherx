package rancher

import (
	"context"

	apiv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func CreateProject(ctx context.Context, client *rest.RESTClient, name string, displayName string, clusterName string) error {

	obj := &apiv3.Project{
		ObjectMeta: v1.ObjectMeta{
			Name: name,
		},
		Spec: apiv3.ProjectSpec{
			ClusterName: clusterName,
			DisplayName: displayName,
		},
	}

	// Object: map[string]interface{}{
	// 	"apiVersion": "management.cattle.io/v3",
	// 	"kind":       "Project",
	// 	"metadata": map[string]interface{}{
	// 		"name": name,
	// 	},
	// 	"spec": map[string]interface{}{
	// 		"clusterName": clusterName,
	// 		"displayName": displayName,
	// 	},
	// }

	result := apiv3.ProjectList{}

	err1 := client.
		Get().
		Resource("projects").
		Do(ctx).
		Into(&result)

	if err1 != nil {
		return err1
	}

	result1 := apiv3.Project{}

	err := client.
		Post().
		Resource("projects").
		Namespace(clusterName).
		Body(obj).
		Do(ctx).
		Into(&result1)

	// _, err := client.Resource(schema.GroupVersionResource{
	// 	Group:    "management.cattle.io",
	// 	Version:  "v3",
	// 	Resource: "projects",
	// }).Namespace(clusterName).Create(ctx, obj, v1.CreateOptions{})

	if err != nil {
		return err
	}

	return nil
}
