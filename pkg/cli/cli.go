package cli

import (
	apiv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func NewRootCmd() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "kubectl-rancherx",
		Short: "kubectl-rancherx helps to create k8s objects in a Rancher cluster",
		Long: `
A very simple cli.`,
	}

	rancherXScheme := runtime.NewScheme()

	// scheme.AddToScheme(customScheme)
	apiv3.AddToScheme(rancherXScheme)

	config, err := genericclioptions.NewConfigFlags(true).ToRESTConfig()
	if err != nil {
		return nil, err
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	restClient, err := toRestClient(rancherXScheme, config)
	if err != nil {
		return nil, err
	}

	client := &Client{
		kubeClient,
		dynamicClient,
		restClient,
	}

	rootCmd.AddCommand(
		version(kubeClient),
		pods(kubeClient),        // dev only
		projects(dynamicClient), // dev only
		create(client),
	)

	return rootCmd, nil
}
