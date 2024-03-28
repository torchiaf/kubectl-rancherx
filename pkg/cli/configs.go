package cli

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

type Client struct {
	KubeClient    *kubernetes.Clientset
	DynamicClient *dynamic.DynamicClient
}
