package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"

	rancher "github.com/torchiaf/kubectl-rancherx/pkg/rancher"
)

func get(client *Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "get",
		Short:         "Display one or many Rancher resources.",
		Args:          cobra.ExactArgs(1),
		RunE:          ValidateSubCommand(resources),
		SilenceErrors: true,
	}

	cmd.AddCommand(
		getProjects(client.RestClient),
	)

	return cmd
}

func getProjects(client *rest.RESTClient) *cobra.Command {
	cfg := &ProjectConfig{}

	cmd := &cobra.Command{
		Use:     "project",
		Aliases: []string{"projects"},
		Short:   "Get projects",
		Example: `kubectl rancherx get project [--cluster-name] [projectName]`,
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) > 0 {
				project, err := rancher.GetProject(c.Context(), client, args[0], cfg.ClusterName)
				if err != nil {
					return fmt.Errorf("getting projects: %w", err)
				}

				fmt.Printf("%s\n", project.Spec.DisplayName)
			} else {
				projects, err := rancher.ListProjects(c.Context(), client, cfg.ClusterName)
				if err != nil {
					return fmt.Errorf("getting projects: %w", err)
				}

				if len(projects.Items) == 0 {
					fmt.Printf("No resources found in %q cluster.\n", cfg.ClusterName)
				}

				for _, project := range projects.Items {
					fmt.Printf("%s\n", project.Spec.DisplayName)
				}
			}

			return nil
		},
		SilenceErrors: true,
	}

	cmd.Flags().StringVar(&cfg.ClusterName, "cluster-name", "", "ClusterName is the name of the cluster the project belongs to. Immutable.")
	cmd.MarkFlagRequired("cluster-name")
	cmd.RegisterFlagCompletionFunc("cluster-name", ClustersFlagValidator(client))

	return cmd
}
