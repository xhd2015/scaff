# Scenario

**Feature**: scaff fix project/agents

```
# scaffold root AGENTS.md from go.mod metadata
fix executor -> project/agents -> AGENTS.md with build/test sections
```

## Preconditions

- A Go project fixture is prepared per leaf case (typically `github.com/xhd2015/myapp`).

## Steps

1. Materialize AGENTS.md state for the scenario.
2. Run `scaff fix project/agents` with optional `--dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoModGitHubScaffold(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

// markProjectAgentsTree keeps hierarchical child packages importing this package live.
func markProjectAgentsTree() {}
```