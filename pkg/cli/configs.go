package cli

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client struct {
	KubeClient    *kubernetes.Clientset
	DynamicClient *dynamic.DynamicClient
	RestClient    *rest.RESTClient
}
