# kubectl-rancherx
A kubectl plugin for Rancher

### Generate docs

```bash
go run pkg/docgen/main.go
```

[Commands list](docs/kubectl-rancherx.md)

### Tests

```bash
go test ./... -coverprofile=coverage.out
```