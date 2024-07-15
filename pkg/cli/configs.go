package cli

import (
	apiv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"k8s.io/apimachinery/pkg/runtime"
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
	config.ContentConfig.GroupVersion = &apiv3.SchemeGroupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.NewCodecFactory(runtimeScheme)
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	return rest.UnversionedRESTClientFor(config)
}
