package rancher

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/torchiaf/kubectl-rancherx/pkg/flag"
	"github.com/torchiaf/kubectl-rancherx/pkg/manager"
	"github.com/torchiaf/kubectl-rancherx/pkg/output"
	"k8s.io/client-go/rest"
)

const (
	projectRoleTemplateBinding = "projectRoleTemplateBindings"
)

var ProjectRoleTemplateBindingResource = []string{
	"projectRoleTemplateBinding", "projectRoleTemplateBindings",
}

type ProjectRoleTemplateBindingConfig struct {
	ProjectName string
	// Interactive bool
	Common flag.CommonConfig
}

type ProjectRoleTemplateBindingOutput struct {
	v3.ProjectRoleTemplateBinding
	ProjectDisplayName string
}

func ListProjectRoleTemplateBindings(ctx context.Context, client *rest.RESTClient, projectName string) (*v3.ProjectRoleTemplateBindingList, error) {

	projectRoleTemplateBindings := &v3.ProjectRoleTemplateBindingList{}

	err := manager.List(ctx, client, projectRoleTemplateBinding, projectName, projectRoleTemplateBindings)
	if err != nil {
		return &v3.ProjectRoleTemplateBindingList{}, err
	}

	return projectRoleTemplateBindings, nil
}

func PrintProjectRoleTemplateBindings(ctx context.Context, projectRoleTemplateBindings []ProjectRoleTemplateBindingOutput, cfg *ProjectRoleTemplateBindingConfig) error {
	table := output.Table[ProjectRoleTemplateBindingOutput]{
		Header: []string{
			"name",
			"project-name",
			"role-template",
		},
		Row: func(item ProjectRoleTemplateBindingOutput) []string {
			return []string{
				item.Name,
				item.ProjectDisplayName,
				item.RoleTemplateName,
			}
		},
	}
	return output.Print(ctx, cfg.Common.Output, projectRoleTemplateBindings, table)
}
