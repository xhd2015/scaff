# Scenario

**Feature**: scaff fix project/license

```
# scaffold root LICENSE from go.mod metadata
fix executor -> project/license -> MIT LICENSE with year + owner
```

## Preconditions

- A Go project fixture is prepared per leaf case (typically `github.com/xhd2015/myapp`).

## Steps

1. Materialize LICENSE file state for the scenario.
2. Run `scaff fix project/license` with optional `--dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoModGitHubScaffold(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

// markProjectLicenseTree keeps hierarchical child packages importing this package live.
func markProjectLicenseTree() {}
```