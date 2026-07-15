# Scenario

**Feature**: fix git/hooks is idempotent

```
# existing hook runner -> no-op
git/hooks fix -> nothing to do
```

## Preconditions

- `script/git-hooks/main.go` already exists.

## Steps

1. Write existing hook runner.
2. Run `scaff fix git/hooks`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGitHooksMain(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "git/hooks"}
	return nil
}
```