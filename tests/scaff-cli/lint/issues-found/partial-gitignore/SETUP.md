# Scenario

**Feature**: lint reports partial git/ignore

```
# .gitignore missing .vscode/ -> git/ignore partial/missing
Rule git/ignore -> partial status for missing pattern
```

## Preconditions

- Project has `go.mod` and a partial `.gitignore` missing `.vscode/`.

## Steps

1. Write `go.mod` and partial `.gitignore`.
2. Run `scaff lint`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	if err := writePartialGitignore(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"lint"}
	return nil
}
```