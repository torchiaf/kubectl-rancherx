## kubectl-rancherx get project

Get projects

```
kubectl-rancherx get project [flags]
```

### Examples

```
kubectl rancherx get project [--cluster-name] [projectName]
```

### Options

```
      --cluster-name string   ClusterName is the name of the cluster the project belongs to. Immutable.
  -h, --help                  help for project
  -o, --output string         Output format. One of: (json, yaml)
```

### Options inherited from parent commands

```
      --log-file string   print logs to file
      --no-color          disable colors in logs output
  -v, --verbosity int     level of log verbosity
```

### SEE ALSO

* [kubectl-rancherx get](kubectl-rancherx_get.md)	 - Display one or many Rancher resources.

###### Auto generated by spf13/cobra on 15-Jan-2025
