package rancher

import (
	"context"
	"fmt"

	apiv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/torchiaf/kubectl-rancherx/pkg/manager"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type RancherResource struct {
	Project string
}

var resource = &RancherResource{
	Project: "projects",
}

func GetProject(ctx context.Context, client *rest.RESTClient, name string, clusterName string) (*apiv3.Project, error) {

	projects := &apiv3.ProjectList{}

	err := manager.List(ctx, client, resource.Project, clusterName, projects)
	if err != nil {
		return &apiv3.Project{}, err
	}

	for _, project := range projects.Items {
		if name == project.Spec.DisplayName {
			return &project, nil
		}
	}

	return &apiv3.Project{}, fmt.Errorf("project %q not found in cluster %q", name, clusterName)
}

func ListProjects(ctx context.Context, client *rest.RESTClient, clusterName string) (*apiv3.ProjectList, error) {

	projects := &apiv3.ProjectList{}

	err := manager.List(ctx, client, resource.Project, clusterName, projects)
	if err != nil {
		return &apiv3.ProjectList{}, err
	}

	return projects, nil
}

func CreateProject(ctx context.Context, client *rest.RESTClient, name string, displayName string, clusterName string) error {

	projects := &apiv3.ProjectList{}

	err := manager.List(ctx, client, resource.Project, clusterName, projects)
	if err != nil {
		return err
	}

	for _, project := range projects.Items {
		if displayName == project.Spec.DisplayName {
			return fmt.Errorf("project %q already exists in cluster %q", displayName, clusterName)
		}
	}

	obj := &apiv3.Project{
		// ObjectMeta: v1.ObjectMeta{
		// 	Name: name,
		// },
		ObjectMeta: v1.ObjectMeta{
			GenerateName: "p-",
		},
		Spec: apiv3.ProjectSpec{
			ClusterName: clusterName,
			DisplayName: displayName,
		},
	}

	return manager.Create(ctx, client, resource.Project, clusterName, obj)
}

func DeleteProject(ctx context.Context, client *rest.RESTClient, name string, clusterName string) error {

	projects := &apiv3.ProjectList{}

	err := manager.List(ctx, client, resource.Project, clusterName, projects)
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

	return manager.Delete(ctx, client, resource.Project, clusterName, projectName)
}
