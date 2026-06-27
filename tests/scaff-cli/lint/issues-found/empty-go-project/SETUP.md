# Scenario

**Feature**: lint empty Go project reports only default rules

```
# go.mod only -> git.ignore + github.testing.workflow issues
lint orchestrator -> two default rules only (no opt-in rules)
```

## Preconditions

- Project contains only `go.mod`.

## Steps

1. Write `go.mod` to the project directory.
2. Run `scaff lint`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"lint"}
	return nil
}
```