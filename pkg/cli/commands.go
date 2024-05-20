package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func newPodsCmd(kubeClient kubernetes.Interface) *cobra.Command {
	return &cobra.Command{
		Hidden: true,
		Use:    "pods",
		Short:  "Print pods",
		RunE: func(c *cobra.Command, args []string) error {
			pods, err := kubeClient.CoreV1().Pods("default").List(c.Context(), v1.ListOptions{})
			if err != nil {
				return err
			}

			fmt.Printf("Pods: ")
			for i := 0; i < len(pods.Items); i++ {
				fmt.Println(pods.Items[i].Name)
			}

			return nil
		},
	}
}
