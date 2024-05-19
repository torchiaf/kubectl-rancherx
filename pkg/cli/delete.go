package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"

	rancher "github.com/torchiaf/kubectl-rancherx/pkg/rancher"
)

func newDeleteCmd(client *Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "delete",
		Short:         "Delete Rancher resources by resources and names.",
		Args:          cobra.ExactArgs(1),
		RunE:          ValidateSubCommand(resources),
		SilenceErrors: true,
	}

	cmd.AddCommand(
		newDeleteProjectCmd(client.RestClient),
	)

	return cmd
}

func newDeleteProjectCmd(client *rest.RESTClient) *cobra.Command {
	cfg := &ProjectConfig{}

	cmd := &cobra.Command{
		Use:               "project",
		Aliases:           []string{"projects"},
		Short:             "Delete a project",
		Example:           `kubectl rancherx delete project [--cluster-name] projectName`,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: ProjectArgValidator(client),
		RunE: func(c *cobra.Command, args []string) error {

			projectName := args[0]

			err := rancher.DeleteProject(c.Context(), client, projectName, cfg.ClusterName)

			if err != nil {
				return fmt.Errorf("deleting project: %w", err)
			}

			fmt.Printf("Project: %q deleted \n", projectName)

			return nil
		},
		SilenceErrors: true,
	}

	cmd.Flags().StringVar(&cfg.ClusterName, "cluster-name", "", "ClusterName is the name of the cluster the project belongs to. Immutable.")
	cmd.MarkFlagRequired("cluster-name")

	return cmd
}
