package rancher

import (
	"github.com/torchiaf/kubectl-rancherx/pkg/flag"
)

type ProjectConfig struct {
	DisplayName string
	ClusterName string
	Interactive bool
	Common      flag.CommonConfig
}
