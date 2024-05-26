package manager

import (
	"context"

	log "github.com/sirupsen/logrus"
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
		log.Tracef("manager.List -  resource [%s] - namespace [%s] - error [%s]", resource, namespace, err)
		return err
	}

	log.Tracef("manager.List - resource [%s] - namespace [%s] - obj [%+v]", resource, namespace, obj)

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
		log.Tracef("manager.Get - resource [%s] - namespace [%s] - name [%s] - error [%s]", resource, namespace, name, err)
		return err
	}

	log.Tracef("manager.Get - resource [%s] - namespace [%s] - name [%s] - obj [%+v]", resource, namespace, name, obj)
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
		log.Tracef("manager.Create - resource [%s] - namespace [%s] - error [%s]", resource, namespace, err)
	}

	log.Tracef("manager.Create - resource [%s] - namespace [%s] - obj [%+v]", resource, namespace, obj)

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
		log.Tracef("manager.Delete - resource [%s] - namespace [%s] - name [%s] - error [%s]", resource, namespace, name, err)
	}

	log.Tracef("manager.Delete - resource [%s] - namespace [%s] - name [%s]", resource, namespace, name)

	return err
}
