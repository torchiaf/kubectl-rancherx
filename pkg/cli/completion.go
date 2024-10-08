package cli

import (
	apiv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/spf13/cobra"
	"github.com/torchiaf/kubectl-rancherx/pkg/manager"
	"k8s.io/client-go/rest"
)

var emptyCompletions = []string{""}

func NoFileCompletions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return emptyCompletions, cobra.ShellCompDirectiveNoFileComp
}

func ProjectArgValidator(client *rest.RESTClient) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		projects := &apiv3.ProjectList{}

		err := manager.List(cmd.Context(), client, "projects", "", projects)
		if err != nil {
			return emptyCompletions, cobra.ShellCompDirectiveNoFileComp
		}

		var list []string

		for _, project := range projects.Items {
			list = append(list, project.Name)
		}

		return list, cobra.ShellCompDirectiveNoFileComp
	}
}

func ClustersFlagValidator(client *rest.RESTClient) func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		clusters := &apiv3.ClusterList{}

		err := manager.List(cmd.Context(), client, "clusters", "", clusters)
		if err != nil {
			return emptyCompletions, cobra.ShellCompDirectiveNoFileComp
		}

		var list []string

		for _, cluster := range clusters.Items {
			list = append(list, cluster.Spec.DisplayName)
		}

		return list, cobra.ShellCompDirectiveNoFileComp
	}
}
