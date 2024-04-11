package cli

import (
	"github.com/torchiaf/kubectl-rancherx/pkg/scheme"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client struct {
	KubeClient    *kubernetes.Clientset
	DynamicClient *dynamic.DynamicClient
	RestClient    *rest.RESTClient
}

func toRestClient(runtimeScheme *runtime.Scheme, config *rest.Config) (*rest.RESTClient, error) {
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: scheme.GroupName, Version: scheme.GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.NewCodecFactory(runtimeScheme)
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	return rest.UnversionedRESTClientFor(config)
}
