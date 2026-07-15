# Scenario

**Feature**: scaff fix git/hooks/install

```
# patch .git/hooks when git/hooks scaffold exists
fix executor -> git/hooks/install -> pre-commit/pre-push markers
```

## Preconditions

- Project directory and git/hooks scaffold state vary per leaf.

## Steps

1. Materialize git repo and hook scaffold state.
2. Run `scaff fix git/hooks/install`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}
```