# Scenario

**Feature**: scaff fix unknown rule

```
# invalid rule id -> usage error listing available rules
fix executor -> unknown rule -> exit 2
```

## Preconditions

- Project directory exists (content irrelevant).

## Steps

1. Run `scaff fix unknown.rule`.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "unknown.rule"}
	return nil
}
```