# Scenario

**Feature**: scaff fix git/hooks

```
# scaffold script/git-hooks/main.go with install + no-op hooks
fix executor -> git/hooks -> hook runner without sub-check dirs
```

## Preconditions

- Project directory exists.

## Steps

1. Materialize `script/git-hooks/main.go` state for the scenario.
2. Run `scaff fix git/hooks`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

```