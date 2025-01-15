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
		RunE:          ValidateSubCommand(rancher.ProjectResources),
		SilenceErrors: true,
	}

	cmd.AddCommand(
		newGetProjectsCmd(client.RestClient),
		newGetProjectRoleTemplateBindingsCmd(client.RestClient),
	)

	return cmd
}

func newGetProjectsCmd(client *rest.RESTClient) *cobra.Command {
	cfg := &rancher.ProjectConfig{}

	cmd := &cobra.Command{
		Use:               "project",
		Aliases:           []string{"projects"},
		Short:             "Get projects",
		Example:           `kubectl rancherx get project [--cluster-name] [--display-name] [projectName]`,
		ValidArgsFunction: NoFileCompletions,
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

func newGetProjectRoleTemplateBindingsCmd(client *rest.RESTClient) *cobra.Command {
	cfg := &rancher.ProjectRoleTemplateBindingConfig{}

	cmd := &cobra.Command{
		Use:               "projectroletemplatebinding",
		Aliases:           []string{"projectroletemplatebindings"},
		Short:             "Get projectroletemplatebindings",
		Example:           `kubectl rancherx get projectroletemplatebindings [--project] [projectroletemplatebindings]`,
		ValidArgsFunction: NoFileCompletions,
		RunE: func(c *cobra.Command, args []string) error {

			projects, err := rancher.ListProjects(c.Context(), client, "")
			if err != nil {
				return fmt.Errorf("getting projects: %w", err)
			}

			projectName := ""
			projectMap := make(map[string]v3.Project)

			for i := 0; i < len(projects.Items); i++ {
				if cfg.ProjectName != "" && cfg.ProjectName == projects.Items[i].Spec.DisplayName {
					projectName = projects.Items[i].Name
				}
				projectMap[projects.Items[i].Name] = projects.Items[i]
			}

			if cfg.ProjectName != "" && projectName == "" {
				fmt.Printf("Project %q not found.\n", cfg.ProjectName)
				return nil
			}

			projectRoleTemplateBindings, err := rancher.ListProjectRoleTemplateBindings(c.Context(), client, projectName)
			if err != nil {
				return fmt.Errorf("getting projectRoleTemplateBinding: %w", err)
			}

			items := []rancher.ProjectRoleTemplateBindingOutput{}

			if len(args) > 0 {
				prtbMap := make(map[string]v3.ProjectRoleTemplateBinding)
				for i := 0; i < len(projectRoleTemplateBindings.Items); i++ {
					prtbMap[projectRoleTemplateBindings.Items[i].Name] = projectRoleTemplateBindings.Items[i]
				}

				for _, arg := range args {
					if prtbMap[arg].Name != "" { // is not empty
						items = append(items, rancher.ProjectRoleTemplateBindingOutput{
							ProjectRoleTemplateBinding: prtbMap[arg],
							ProjectDisplayName:         projectMap[prtbMap[arg].Namespace].Spec.DisplayName,
						})
					} else {
						fmt.Printf("ProjectRoleTemplateBinding %q not found.\n", arg)
					}
				}
			} else {
				slices.SortFunc(projectRoleTemplateBindings.Items, func(a, b v3.ProjectRoleTemplateBinding) int {
					return strings.Compare(a.Name, b.Name)
				})

				for i := 0; i < len(projectRoleTemplateBindings.Items); i++ {
					items = append(items, rancher.ProjectRoleTemplateBindingOutput{
						ProjectRoleTemplateBinding: projectRoleTemplateBindings.Items[i],
						ProjectDisplayName:         projectMap[projectRoleTemplateBindings.Items[i].Namespace].Spec.DisplayName,
					})
				}

				if len(items) == 0 {
					if cfg.ProjectName != "" {
						fmt.Printf("No pojectRoleTemplateBindings found for %q project.\n", cfg.ProjectName)
					} else {
						fmt.Print("No pojectRoleTemplateBindings found.\n")
					}

					return nil
				}
			}

			rancher.PrintProjectRoleTemplateBindings(
				c.Context(),
				items,
				cfg,
			)

			return nil
		},
		SilenceErrors: true,
	}

	cmd.Flags().StringVar(&cfg.ProjectName, "project-name", "", "ProjectName is the name of the project of the ProjectRoleTemplateBinding. Immutable.")
	cmd.Flags().StringVarP(&cfg.Common.Output, "output", "o", "", "Output format. One of: (json, yaml)")

	// cmd.MarkFlagRequired("project-name")
	cmd.RegisterFlagCompletionFunc("project-name", ProjectFlagValidator(client))

	return cmd
}
