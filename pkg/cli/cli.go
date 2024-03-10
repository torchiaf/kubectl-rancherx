package cli

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "kubectl-rancherx",
		Short: "kubectl-rancherx helps to create k8s objects in a Rancher cluster",
		Long: `
A very simple cli.`,
	}

	rootCmd.AddCommand(
		print(),
	)

	return rootCmd, nil
}
