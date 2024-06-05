package manager

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"

	"github.com/torchiaf/kubectl-rancherx/pkg/log"
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
		log.Debug(resource, namespace, err)
		return err
	}

	log.Debug(resource, namespace, obj)

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
		log.Debug(resource, namespace, name, err)
		return err
	}

	log.Debug(resource, namespace, name, obj)
	return nil
}

func Create(ctx context.Context, client *rest.RESTClient, resource string, namespace string, obj interface{}) error {
	res := client.
		Post().
		Resource(resource).
		Namespace(namespace).
		Body(obj).
		Do(ctx)

	err := res.Error()
	if err != nil {
		log.Debug(resource, namespace, err)
	}

	log.Debug(resource, namespace, obj)

	return err
}

func Delete(ctx context.Context, client *rest.RESTClient, resource string, namespace string, name string) error {
	res := client.
		Delete().
		Resource(resource).
		Namespace(namespace).
		Name(name).
		Do(ctx)

	err := res.Error()
	if err != nil {
		log.Debug(resource, namespace, name, err)
	}

	log.Debug(resource, namespace, name)

	return err
}
