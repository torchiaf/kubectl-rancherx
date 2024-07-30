package cli

import (
	"fmt"
	"slices"
	"strings"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"

	rancher "github.com/torchiaf/kubectl-rancherx/pkg/rancher"
)

func newGetCmd(client *Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "get",
		Short:         "Display one or many Rancher resources.",
		Args:          cobra.ExactArgs(1),
		RunE:          ValidateSubCommand(rancher.Resources),
		SilenceErrors: true,
	}

	cmd.AddCommand(
		newGetProjectsCmd(client.RestClient),
	)

	return cmd
}

func newGetProjectsCmd(client *rest.RESTClient) *cobra.Command {
	cfg := &rancher.ProjectConfig{}

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

			items := []v3.Project{}

			if len(args) > 0 {
				projectMap := make(map[string]v3.Project)
				for i := 0; i < len(projects.Items); i++ {
					projectMap[projects.Items[i].Spec.DisplayName] = projects.Items[i]
				}

				for _, arg := range args {
					if projectMap[arg].Name != "" { // is not empty project
						items = append(items, projectMap[arg])
					} else {
						fmt.Printf("Project %q not found.\n", arg)
					}
				}
			} else {
				slices.SortFunc(projects.Items, func(a, b v3.Project) int {
					return strings.Compare(a.Spec.DisplayName, b.Spec.DisplayName)
				})

				items = append(items, projects.Items...)

				if len(items) == 0 {
					fmt.Printf("No projects found in %q cluster.\n", cfg.ClusterName)
					return nil
				}
			}

			rancher.PrintProject(
				c.Context(),
				items,
				cfg,
				func(item v3.Project) string {
					return fmt.Sprintf("%s\t%s", item.Name, item.Spec.DisplayName)
				},
			)

			return nil
		},
		SilenceErrors: true,
	}

	cmd.Flags().StringVar(&cfg.ClusterName, "cluster-name", "", "ClusterName is the name of the cluster the project belongs to. Immutable.")
	cmd.Flags().StringVarP(&cfg.Common.Output, "output", "o", "", "Output format. One of: (json, yaml)")

	cmd.MarkFlagRequired("cluster-name")
	cmd.RegisterFlagCompletionFunc("cluster-name", ClustersFlagValidator(client))

	return cmd
}
