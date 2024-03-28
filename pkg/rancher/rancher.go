package rancher

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func CreateProject(ctx context.Context, client *dynamic.DynamicClient, name string, displayName string, clusterName string) error {
	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "management.cattle.io/v3",
			"kind":       "Project",
			"metadata": map[string]interface{}{
				"name": name,
			},
			"spec": map[string]interface{}{
				"clusterName": clusterName,
				"displayName": displayName,
			},
		},
	}

	_, err := client.Resource(schema.GroupVersionResource{
		Group:    "management.cattle.io",
		Version:  "v3",
		Resource: "projects",
	}).Namespace(clusterName).Create(ctx, obj, v1.CreateOptions{})

	if err != nil {
		return err
	}

	return nil
}
