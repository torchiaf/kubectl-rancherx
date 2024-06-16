package rancher

import (
	"context"
	"encoding/json"
	"fmt"

	// "github.com/itchyny/gojq"
	"github.com/tidwall/sjson"

	apiv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/torchiaf/kubectl-rancherx/pkg/flag"
	"github.com/torchiaf/kubectl-rancherx/pkg/manager"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func patchStruct[T comparable](obj *T, set flag.Set) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	newValue := string(data)

	for k, v := range set {
		newValue, err = sjson.Set(newValue, k, v)
		if err != nil {
			return err
		}
	}

	json.Unmarshal([]byte(newValue), &obj)
	if err != nil {
		return err
	}

	return nil
}

const (
	project = "projects"
)

var Resources = []string{
	"project", "projects",
}

func GetProject(ctx context.Context, client *rest.RESTClient, name string, clusterName string) (*apiv3.Project, error) {

	projects := &apiv3.ProjectList{}

	err := manager.List(ctx, client, project, clusterName, projects)
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

	err := manager.List(ctx, client, project, clusterName, projects)
	if err != nil {
		return &apiv3.ProjectList{}, err
	}

	return projects, nil
}

func CreateProject(ctx context.Context, client *rest.RESTClient, name string, cfg *ProjectConfig) error {

	projects := &apiv3.ProjectList{}

	err := manager.List(ctx, client, project, cfg.ClusterName, projects)
	if err != nil {
		return err
	}

	for _, project := range projects.Items {
		if cfg.DisplayName == project.Spec.DisplayName {
			return fmt.Errorf("project %q already exists in cluster %q", cfg.DisplayName, cfg.ClusterName)
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
			ClusterName: cfg.ClusterName,
			DisplayName: cfg.DisplayName,
		},
	}

	if cfg.Set != nil {
		err := patchStruct(&obj, cfg.Set)
		if err != nil {
			return err
		}
	}

	return manager.Create(ctx, client, project, cfg.ClusterName, obj)
}

func DeleteProject(ctx context.Context, client *rest.RESTClient, name string, clusterName string) error {

	projects := &apiv3.ProjectList{}

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
