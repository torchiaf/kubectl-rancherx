package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func print() *cobra.Command {
	return &cobra.Command{
		Use:   "print",
		Short: "Print Hello World!",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello World!\n")
		},
	}
}

func version(kubeClient kubernetes.Interface) *cobra.Command {
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

func pods(kubeClient kubernetes.Interface) *cobra.Command {
	return &cobra.Command{
		Use:   "pods",
		Short: "Print pods",
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

func projects(kubeClient *dynamic.DynamicClient) *cobra.Command {
	return &cobra.Command{
		Use:   "projects",
		Short: "Print projects",
		RunE: func(c *cobra.Command, args []string) error {

			projects, err := kubeClient.Resource(schema.GroupVersionResource{
				Group:    "management.cattle.io",
				Version:  "v3",
				Resource: "projects",
			}).Namespace("c-m-79djmg9n").List(c.Context(), v1.ListOptions{})

			if err != nil {
				return err
			}

			for i := 0; i < len(projects.Items); i++ {
				fmt.Println(projects.Items[i].GetName())
			}

			return nil

		},
	}
}
