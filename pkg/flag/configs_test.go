package flag

import (
	"encoding/json"
	"strings"

	apiv3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"context"
	"testing"
)

func TestMergeValues(t *testing.T) {

	type args struct {
		ctx context.Context
		obj apiv3.Project
		cfg *CommonConfig
	}

	type test struct {
		name string
		args args
		want string
	}

	test1 := test{
		name: "should merge Set value: --set key1=value1",
		args: args{
			ctx: context.TODO(),
			obj: apiv3.Project{
				Spec: apiv3.ProjectSpec{
					Description: "",
				},
			},
			cfg: &CommonConfig{
				Set: []string{
					"spec.description=bar1",
				},
			},
		},
		want: "bar1",
	}

	t.Run(test1.name, func(t *testing.T) {
		err := MergeValues(test1.args.ctx, &test1.args.obj, test1.args.cfg.Set)
		if err != nil {
			t.Errorf("MergeValues() error = %v", err)
		}
		description := test1.args.obj.Spec.Description
		if description != "bar1" {
			t.Errorf("MergeValues() description = %v, want %v", description, test1.want)
		}
	})

	test2 := test{
		name: "should not merge not existent Set value: --set key1=value1",
		args: args{
			ctx: context.TODO(),
			obj: apiv3.Project{},
			cfg: &CommonConfig{
				Set: []string{
					"spec.pippo=bar1",
				},
			},
		},
		want: "",
	}

	t.Run(test2.name, func(t *testing.T) {
		err := MergeValues(test2.args.ctx, &test2.args.obj, test2.args.cfg.Set)
		if err != nil {
			t.Errorf("MergeValues() error = %v", err)
		}

		jsonData, err := json.Marshal(test2.args.obj)
		if err != nil {
			t.Errorf("MergeValues() error = %v", err)
		}

		stringValues := string(jsonData[:])

		if strings.Contains(stringValues, "pippo") {
			t.Errorf("MergeValues() found = pippo, want %v", test2.want)
		}
	})

	test3 := test{
		name: "should merge multiple Set values: --set key1=value1,key2=value2",
		args: args{
			ctx: context.TODO(),
			obj: apiv3.Project{
				ObjectMeta: v1.ObjectMeta{Labels: make(map[string]string)},
				Spec: apiv3.ProjectSpec{
					Description: "",
				},
			},
			cfg: &CommonConfig{
				Set: []string{
					"spec.description=bar1",
					"metadata.generateName=-p",
				},
			},
		},
	}

	t.Run(test3.name, func(t *testing.T) {
		err := MergeValues(test3.args.ctx, &test3.args.obj, test3.args.cfg.Set)
		if err != nil {
			t.Errorf("MergeValues() error = %v", err)
		}

		description := test3.args.obj.Spec.Description
		if description != "bar1" {
			t.Errorf("MergeValues() description = %v, want %v", description, "bar1")
		}

		generateName := test3.args.obj.ObjectMeta.GenerateName
		if generateName != "-p" {
			t.Errorf("MergeValues() generateName = %v, want %v", generateName, "-p")
		}
	})
}
