package rancher

import (
	"github.com/torchiaf/kubectl-rancherx/pkg/flag"
)

type ProjectConfig struct {
	Set         flag.Set
	DisplayName string
	ClusterName string
}
