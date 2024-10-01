package rancher

import (
	"context"
	"fmt"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/torchiaf/kubectl-rancherx/pkg/flag"
	"github.com/torchiaf/kubectl-rancherx/pkg/manager"
	"github.com/torchiaf/kubectl-rancherx/pkg/output"
	"github.com/torchiaf/kubectl-rancherx/pkg/prompt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

const (
	project = "projects"
)

var Resources = []string{
	"project", "projects",
}

type projectData struct {
	Name        string
	DisplayName string
	ClusterName string
	Set         []string
}

func GetProject(ctx context.Context, client *rest.RESTClient, name string, clusterName string) (*v3.Project, error) {

	projects := &v3.ProjectList{}

	err := manager.List(ctx, client, project, clusterName, projects)
	if err != nil {
		return &v3.Project{}, err
	}

	for _, project := range projects.Items {
		if name == project.Spec.DisplayName {
			return &project, nil
		}
	}

	return &v3.Project{}, fmt.Errorf("project %q not found in cluster %q", name, clusterName)
}

func ListProjects(ctx context.Context, client *rest.RESTClient, clusterName string) (*v3.ProjectList, error) {

	projects := &v3.ProjectList{}

	err := manager.List(ctx, client, project, clusterName, projects)
	if err != nil {
		return &v3.ProjectList{}, err
	}

	return projects, nil
}

func CreateProjectI(ctx context.Context, client *rest.RESTClient, name string) error {

	clusters, _ := ListClusters(ctx, client)

	var items []string
	for _, cluster := range clusters.Items {
		items = append(items, cluster.Name)
	}

	clusterNamePromptContent := prompt.PromptContent{
		ErrorMsg: "Please provide a Cluster Name.",
		Label:    "Which cluster do you want to create the project in?",
	}
	clusterName := prompt.PromptGetSelect(clusterNamePromptContent, items)

	displayNamePromptContent := prompt.PromptContent{
		ErrorMsg: "Please provide a Display Name.",
		Label:    "What name would you like to assign to the project?",
	}
	displayName := prompt.PromptGetInput(displayNamePromptContent)

	data := projectData{
		Name:        name,
		ClusterName: clusterName,
		DisplayName: displayName,
		Set:         nil,
	}

	return createProject(ctx, client, data)
}

func CreateProject(ctx context.Context, client *rest.RESTClient, name string, cfg *ProjectConfig) error {

	data := projectData{
		Name:        name,
		DisplayName: cfg.DisplayName,
		ClusterName: cfg.ClusterName,
		Set:         cfg.Common.Set,
	}

	return createProject(ctx, client, data)
}

func createProject(ctx context.Context, client *rest.RESTClient, data projectData) error {

	projects := &v3.ProjectList{}

	err := manager.List(ctx, client, project, data.ClusterName, projects)
	if err != nil {
		return err
	}

	for _, project := range projects.Items {
		if data.DisplayName == project.Spec.DisplayName {
			return fmt.Errorf("project %q already exists in cluster %q", data.DisplayName, data.ClusterName)
		}
	}

	obj := &v3.Project{
		// ObjectMeta: v1.ObjectMeta{
		// 	Name: name,
		// },
		ObjectMeta: v1.ObjectMeta{
			GenerateName: "p-",
		},
		Spec: v3.ProjectSpec{
			ClusterName: data.ClusterName,
			DisplayName: data.DisplayName,
		},
	}

	err = flag.MergeValues(ctx, &obj, data.Set)
	if err != nil {
		return err
	}

	return manager.Create(ctx, client, project, data.ClusterName, obj)
}

func DeleteProject(ctx context.Context, client *rest.RESTClient, name string, clusterName string) error {

	projects := &v3.ProjectList{}

	err := manager.List(ctx, client, project, clusterName, projects)
	if err != nil {
		return err
	}

	var projectName = ""

	for _, project := range projects.Items {
		if name == project.Spec.DisplayName {
			projectName = project.Name
			break
		}
	}

	if len(projectName) == 0 {
		return fmt.Errorf("project %q not found in cluster %q", name, clusterName)
	}

	return manager.Delete(ctx, client, project, clusterName, projectName)
}

func PrintProject(ctx context.Context, items []v3.Project, cfg *ProjectConfig, def func(item v3.Project) string) error {
	for _, item := range items {
		item.ObjectMeta.ManagedFields = nil

		err := output.Print(ctx, item, cfg.Common.Output, def)
		if err != nil {
			return err
		}
	}

	return nil
}
