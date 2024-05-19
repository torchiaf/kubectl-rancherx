package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"

	rancher "github.com/torchiaf/kubectl-rancherx/pkg/rancher"
)

func newCreateCmd(client *Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "create",
		Short:         "Create a Rancher resource from a file or from stdin.",
		Args:          cobra.ExactArgs(1),
		RunE:          ValidateSubCommand(resources),
		SilenceErrors: true,
	}

	cmd.AddCommand(
		newCreateProjectCmd(client.RestClient),
	)

	return cmd
}

func newCreateProjectCmd(client *rest.RESTClient) *cobra.Command {
	cfg := &ProjectConfig{}

	cmd := &cobra.Command{
		Use:               "project",
		Aliases:           []string{"projects"},
		Short:             "Create a project",
		Example:           `kubectl rancherx create project [--display-name] [--cluster-name] projectName`,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: NoFileCompletions,
		RunE: func(c *cobra.Command, args []string) error {

			projectName := args[0]

			err := rancher.CreateProject(c.Context(), client, projectName, cfg.DisplayName, cfg.ClusterName)

			if err != nil {
				return fmt.Errorf("creating project: %w", err)
			}

			fmt.Printf("Project: %q created \n", projectName)

			return nil
		},
		SilenceErrors: true,
	}

	cmd.Flags().StringVar(&cfg.DisplayName, "display-name", "", "DisplayName is the human-readable name for the project.")
	cmd.Flags().StringVar(&cfg.ClusterName, "cluster-name", "", "ClusterName is the name of the cluster the project belongs to. Immutable.")
	cmd.MarkFlagRequired("display-name")
	cmd.MarkFlagRequired("cluster-name")
	cmd.RegisterFlagCompletionFunc("cluster-name", ClustersFlagValidator(client))

	return cmd
}
