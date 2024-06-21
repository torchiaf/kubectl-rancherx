package cli

import (
	apiv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	"github.com/torchiaf/kubectl-rancherx/pkg/log"
)

func NewRootCmd() (*cobra.Command, error) {
	cfg := &log.LogConfig{}

	rootCmd := &cobra.Command{
		Use:   "kubectl-rancherx",
		Short: "kubectl-rancherx helps to create k8s objects in a Rancher cluster",
		Long: `
A very simple cli.`,
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return log.InitLogger(cmd.Context(), cfg)
		},
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
		newVersionCmd(kubeClient),
		newPodsCmd(kubeClient),
		newGetCmd(client),
		newCreateCmd(client),
		newDeleteCmd(client),
	)

	rootCmd.PersistentFlags().IntVarP(&cfg.LogLevel, "verbosity", "v", 0, "level of log verbosity")
	rootCmd.PersistentFlags().StringVarP(&cfg.LogFileName, "log-file", "l", "", "print logs to file")

	return rootCmd, nil
}
