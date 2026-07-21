# Scenario

**Feature**: scaff fix project/readme

```
# scaffold root README.md from go.mod metadata
fix executor -> project/readme -> README.md with install line
```

## Preconditions

- A Go project fixture is prepared per leaf case (typically `github.com/xhd2015/myapp`).

## Steps

1. Materialize README file state for the scenario.
2. Run `scaff fix project/readme` with optional `--dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoModGitHubScaffold(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

// markProjectReadmeTree keeps hierarchical child packages importing this package live.
func markProjectReadmeTree() {}
```