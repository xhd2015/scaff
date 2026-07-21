# Scenario

**Feature**: scaff fix project/layout/cmd

```
# scaffold cmd/<name>/main.go from go.mod module name
fix executor -> project/layout/cmd -> cmd entry main.go
```

## Preconditions

- Project directory exists with `go.mod` (`github.com/xhd2015/myapp`).

## Steps

1. Materialize `cmd/` layout state for the scenario.
2. Run `scaff fix project/layout/cmd` with case-specific flags.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoModGitHubScaffold(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

// markProjectLayoutCmdTree keeps hierarchical child packages importing this package live.
func markProjectLayoutCmdTree() {}
```