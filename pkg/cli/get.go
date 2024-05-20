package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"

	rancher "github.com/torchiaf/kubectl-rancherx/pkg/rancher"
)

func newGetCmd(client *Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "get",
		Short:         "Display one or many Rancher resources.",
		Args:          cobra.ExactArgs(1),
		RunE:          ValidateSubCommand(resources),
		SilenceErrors: true,
	}

	cmd.AddCommand(
		newGetProjectsCmd(client.RestClient),
	)

	return cmd
}

func newGetProjectsCmd(client *rest.RESTClient) *cobra.Command {
	cfg := &ProjectConfig{}

	cmd := &cobra.Command{
		Use:     "project",
		Aliases: []string{"projects"},
		Short:   "Get projects",
		Example: `kubectl rancherx get project [--cluster-name] [projectName]`,
		RunE: func(c *cobra.Command, args []string) error {

			projects, err := rancher.ListProjects(c.Context(), client, cfg.ClusterName)
			if err != nil {
				return fmt.Errorf("getting projects: %w", err)
			}

			if len(projects.Items) == 0 {
				fmt.Printf("No resources found in %q cluster.\n", cfg.ClusterName)
			}

			if len(args) > 0 {
				projectMap := make(map[string]string)
				for i := 0; i < len(projects.Items); i++ {
					projectMap[projects.Items[i].Spec.DisplayName] = fmt.Sprintf("%s  %s", projects.Items[i].Name, projects.Items[i].Spec.DisplayName)
				}

				for _, arg := range args {
					fmt.Printf("%s\n", projectMap[arg])
				}
			} else {
				for _, project := range projects.Items {
					fmt.Printf("%s\n", fmt.Sprintf("%s  %s", project.Name, project.Spec.DisplayName))
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
