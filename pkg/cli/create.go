package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	rancher "github.com/torchiaf/kubectl-rancherx/pkg/rancher"
	"k8s.io/client-go/rest"
)

func newCreateCmd(client *Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "create",
		Short:         "Create a Rancher resource from a file or from stdin.",
		Args:          cobra.ExactArgs(1),
		RunE:          ValidateSubCommand(rancher.ProjectResources),
		SilenceErrors: true,
	}

	cmd.AddCommand(
		newCreateProjectCmd(client.RestClient),
	)

	return cmd
}

func newCreateProjectCmd(client *rest.RESTClient) *cobra.Command {
	cfg := &rancher.ProjectConfig{}

	cmd := &cobra.Command{
		Use:               "project",
		Aliases:           []string{"projects"},
		Short:             "Create a project",
		Example:           `kubectl rancherx create project [--display-name] [--cluster-name] projectName`,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: NoFileCompletions,
		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			if !cfg.Interactive {
				c.MarkFlagRequired("display-name")
				c.MarkFlagRequired("cluster-name")
			}

			return nil
		},
		RunE: func(c *cobra.Command, args []string) error {

			var err error
			projectName := args[0]

			if cfg.Interactive {
				err = rancher.CreateProjectI(c.Context(), client, projectName)
			} else {
				err = rancher.CreateProject(c.Context(), client, projectName, cfg)
			}

			if err != nil {
				return fmt.Errorf("creating project: %w", err)
			}

			fmt.Printf("Project: %q created \n", projectName)

			return nil
		},
		SilenceErrors: true,
	}

	cmd.Flags().StringArrayVar(&cfg.Common.Set, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	cmd.Flags().StringVar(&cfg.DisplayName, "display-name", "", "DisplayName is the human-readable name for the project.")
	cmd.Flags().StringVar(&cfg.ClusterName, "cluster-name", "", "ClusterName is the name of the cluster the project belongs to. Immutable.")
	cmd.Flags().BoolVarP(&cfg.Interactive, "interactive", "i", false, "Enable the interactive mode.")

	cmd.RegisterFlagCompletionFunc("cluster-name", ClustersFlagValidator(client))
	cmd.RegisterFlagCompletionFunc("display-name", NoFileCompletions)

	return cmd
}
