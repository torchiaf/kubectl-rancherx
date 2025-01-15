package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
)

func newVersionCmd(kubeClient kubernetes.Interface) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the Kubernetes version",
		RunE: func(c *cobra.Command, args []string) error {
			sv, err := kubeClient.Discovery().ServerVersion()
			if err != nil {
				return err
			}

			fmt.Printf("Kubernetes Version: %s\n", sv.String())

			return nil
		},
	}
}
