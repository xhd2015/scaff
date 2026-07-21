# Scenario

**Feature**: scaff fix tests/doctest

```
# scaffold tests/<name>-cli doctest tree from go.mod metadata
fix executor -> tests/doctest -> DOCTEST.md + SETUP.md
```

## Preconditions

- A Go project fixture is prepared per leaf case (typically `github.com/xhd2015/myapp`).

## Steps

1. Materialize doctest tree state for the scenario.
2. Run `scaff fix tests/doctest` with optional `--dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoModGitHubScaffold(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

// markTestsDoctestTree keeps hierarchical child packages importing this package live.
func markTestsDoctestTree() {}
```