package output

import (
	"context"
	"io"
	"os"
	"testing"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPrintTable(t *testing.T) {
	type args struct {
		items []v3.Project
		table Table[v3.Project]
	}

	type test struct {
		name string
		args args
		want string
	}

	tests := []test{
		{
			name: "Print empty table",
			args: args{
				items: []v3.Project{},
				table: Table[v3.Project]{
					Header: []string{"foo"},
					Row: func(item v3.Project) []string {
						return []string{
							item.Name,
						}
					},
				},
			},
			want: "FOO   \n",
		},
		{
			name: "Print 1 row table",
			args: args{
				items: []v3.Project{{ObjectMeta: v1.ObjectMeta{Name: "pippo"}}},
				table: Table[v3.Project]{
					Header: []string{"header1"},
					Row: func(item v3.Project) []string {
						return []string{
							item.Name,
						}
					},
				},
			},
			want: "HEADER1   \npippo     \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := PrintTable(context.Background(), tt.args.items, tt.args.table)

			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = rescueStdout

			result := string(out)

			if err != nil || result != tt.want {
				t.Errorf("PrintTable() error = %v, want %v, is %v", err, tt.want, result)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	type args struct {
		outputType string
		items      []v3.Project
		table      Table[v3.Project]
	}

	type test struct {
		name string
		args args
		want string
	}

	tests := []test{
		{
			name: "Print Json output",
			args: args{outputType: "json", items: []v3.Project{{ObjectMeta: v1.ObjectMeta{Name: "pippo"}}}},
			want: `{
  "metadata": {
    "name": "pippo",
    "creationTimestamp": null
  },
  "spec": {
    "description": "",
    "enableProjectMonitoring": false
  },
  "status": {
    "conditions": null,
    "podSecurityPolicyTemplateId": ""
  }
}
`,
		},
		{
			name: "Print empty yaml",
			args: args{outputType: "yaml", items: []v3.Project{}},
			want: "[]\n\n",
		},
		{
			name: "Do not print empty table",
			args: args{outputType: "", items: []v3.Project{}},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := Print(context.Background(), tt.args.outputType, tt.args.items, tt.args.table)

			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = rescueStdout

			result := string(out)

			if err != nil || result != tt.want {
				t.Errorf("PrintTable() error = %v, want %v, is %v", err, tt.want, result)
			}
		})
	}
}
