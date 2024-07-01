package manager

import (
	"context"
	"log/slog"

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
		log.Error(
			ctx,
			"error listing resource",
			slog.Group("args",
				"resource", resource,
				"namespace", namespace,
			),
			"err", err,
		)

		return err
	}

	log.Trace(
		ctx,
		"listing resource",
		slog.Group("args",
			"resource", resource,
			"namespace", namespace,
			"obj", obj,
		),
	)

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
		log.Error(
			ctx,
			"error getting resource",
			slog.Group("args",
				"resource", resource,
				"namespace", namespace,
				"name", name,
			),
			"err", err,
		)
		return err
	}

	log.Trace(
		ctx,
		"getting resource",
		slog.Group("args",
			"resource", resource,
			"namespace", namespace,
			"name", name,
			"obj", obj,
		),
	)

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
		log.Error(
			ctx,
			"error creating resource",
			slog.Group("args",
				"resource", resource,
				"namespace", namespace,
			),
			"err", err,
		)

		return err
	}

	log.Trace(
		ctx,
		"creating resource",
		slog.Group("args",
			"resource", resource,
			"namespace", namespace,
			"obj", obj,
		),
	)

	return nil
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
		log.Error(
			ctx,
			"error deleting resource",
			slog.Group("args",
				"resource", resource,
				"namespace", namespace,
				"name", name,
			),
			"err", err,
		)

		return err
	}

	log.Trace(
		ctx,
		"deleting resource",
		slog.Group("args",
			"resource", resource,
			"namespace", namespace,
			"name", name,
		),
	)

	return nil
}
