# Scenario

**Feature**: fix github/testing-workflow --dry-run

```
# --dry-run reports would-create without writing test.yml
github/testing-workflow fix --dry-run -> preview only
```

## Preconditions

- Project has no `test.yml`.

## Steps

1. Run `scaff fix github/testing-workflow --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markGithubTestingWorkflowTree()
	markFixTree()
	req.Args = []string{"fix", "github/testing-workflow", "--dry-run"}
	return nil
}
```