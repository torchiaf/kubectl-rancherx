package cli

var resources = []string{
	"project", "projects",
}

type ProjectConfig struct {
	DisplayName string
	ClusterName string
}
