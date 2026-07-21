# Scenario

**Feature**: scaff fix github/testing-workflow

```
# create .github/workflows/test.yml when missing
fix executor -> github/testing-workflow -> test.yml template
```

## Preconditions

- A Go project fixture is prepared per leaf case.

## Steps

1. Materialize workflow file state for the scenario.
2. Run `scaff fix github/testing-workflow` with optional `--dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

// markGithubTestingWorkflowTree keeps hierarchical child packages importing this package live.
func markGithubTestingWorkflowTree() {}
```