# Scenario

**Feature**: scaff fix git/ignore

```
# merge missing .gitignore patterns per profile
fix executor -> git/ignore -> append-only .gitignore update
```

## Preconditions

- A Go project fixture is prepared per leaf case.

## Steps

1. Materialize `.gitignore` state for the scenario.
2. Run `scaff fix git/ignore` with optional `--dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

```