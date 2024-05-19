package manager

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
)

func List[T runtime.Object](ctx context.Context, client *rest.RESTClient, resource string, namespace string, obj T) error {
	req := client.
		Get().
		Resource(resource)

	if namespace != "" {
		req = req.Namespace(namespace)
	}

	err := req.
		Do(ctx).
		Into(obj)

	if err != nil {
		return err
	}

	return nil
}

func Get[T runtime.Object](ctx context.Context, client *rest.RESTClient, resource string, namespace string, name string, obj T) error {
	req := client.
		Get().
		Resource(resource)

	if namespace != "" {
		req = req.Namespace(namespace)
	}

	err := req.
		Name(name).
		Do(ctx).
		Into(obj)

	if err != nil {
		return err
	}

	return nil
}

func Create(ctx context.Context, client *rest.RESTClient, resource string, namespace string, obj interface{}) error {
	res := client.
		Post().
		Resource(resource).
		Namespace(namespace).
		Body(obj).
		Do(ctx)

	return res.Error()
}

func Delete(ctx context.Context, client *rest.RESTClient, resource string, namespace string, name string) error {
	res := client.
		Delete().
		Resource(resource).
		Namespace(namespace).
		Name(name).
		Do(ctx)

	if res.Error() != nil {
		return res.Error()
	}

	return nil
}
