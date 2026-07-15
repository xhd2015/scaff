# Scenario

**Feature**: fix git/ignore appends missing patterns

```
# partial .gitignore -> append only missing lines
git/ignore fix -> merge without overwriting existing lines
```

## Preconditions

- Project has partial `.gitignore` missing `.vscode/`, `*.swp`, `*~`.

## Steps

1. Write partial `.gitignore`.
2. Run `scaff fix git/ignore`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writePartialGitignore(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "git/ignore"}
	return nil
}
```