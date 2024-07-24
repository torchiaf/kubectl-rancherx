# kubectl-rancherx
A kubectl plugin for Rancher

### Generate docs

```bash
go run pkg/docgen/main.go
```

[Commands list](docs/kubectl-rancherx.md)

### Unit Tests

```bash
go test ./... -coverprofile=coverage.out
```

### E2e tests

#### Add new tests

```bash
ginkgo generate --template rancher_template foo
```
