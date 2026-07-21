# Scenario

**Feature**: lint --json emits structured report

```
# JSON consumer reads LintReport from stdout
scaff lint --json -> valid JSON with rule results
```

## Preconditions

- Project has scaffold gaps (empty Go project).

## Steps

1. Write `go.mod` only.
2. Run `scaff lint --json`.

```go
func Setup(t *testing.T, req *Request) error {
	markJsonOutputTree()
	markLintTree()
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"lint", "--json"}
	return nil
}
```