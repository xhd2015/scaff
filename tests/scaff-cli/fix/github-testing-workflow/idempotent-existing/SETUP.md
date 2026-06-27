# Scenario

**Feature**: fix github.testing.workflow no-op when test.yml exists

```
# existing test.yml -> no overwrite
github.testing.workflow fix -> nothing to do
```

## Preconditions

- Project already has `.github/workflows/test.yml`.

## Steps

1. Write existing `test.yml`.
2. Run `scaff fix github.testing.workflow`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeTestWorkflow(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "github.testing.workflow"}
	return nil
}
```