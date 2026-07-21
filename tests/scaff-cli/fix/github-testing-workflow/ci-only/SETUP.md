# Scenario

**Feature**: fix creates test.yml when only ci.yml exists

```
# ci.yml present does not satisfy lint; fix still creates test.yml
github/testing-workflow fix -> test.yml alongside ci.yml
```

## Preconditions

- Project has `ci.yml` but no `test.yml`.

## Steps

1. Write `ci.yml` only.
2. Run `scaff fix github/testing-workflow`.

```go
func Setup(t *testing.T, req *Request) error {
	markGithubTestingWorkflowTree()
	markFixTree()
	if err := writeCiWorkflow(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "github/testing-workflow"}
	return nil
}
```