package cli

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/dynamic"

	rancher "github.com/torchiaf/kubectl-rancherx/pkg/rancher"
)

func create(client *Client) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a Rancher resource from a file or from stdin.",
	}

	createCmd.AddCommand(
		createProject(client.DynamicClient),
	)

	return createCmd
}

type createProjectConfig struct {
	DisplayName string
	ClusterName string
}

func createProject(dynamic *dynamic.DynamicClient) *cobra.Command {
	cfg := &createProjectConfig{}

	cmd := &cobra.Command{
		Use:     "project <name>",
		Short:   "Create a project",
		Example: `kubectl rancherx create project [--display-name] [--cluster-name] projectName`,
		Args:    cobra.ExactArgs(1),
		RunE: func(c *cobra.Command, args []string) error {

			projectName := args[0]

			rancher.CreateProject(c.Context(), dynamic, projectName, cfg.DisplayName, cfg.ClusterName)

			return nil
		},
	}

	cmd.Flags().StringVar(&cfg.DisplayName, "display-name", "", "DisplayName is the human-readable name for the project.")
	cmd.Flags().StringVar(&cfg.ClusterName, "cluster-name", "", "ClusterName is the name of the cluster the project belongs to. Immutable.")

	return cmd
}
