package rancher

import (
	f "github.com/torchiaf/kubectl-rancherx/pkg/flag"
)

type ProjectConfig struct {
	SetFlag     f.SetFlag
	DisplayName string
	ClusterName string
}
