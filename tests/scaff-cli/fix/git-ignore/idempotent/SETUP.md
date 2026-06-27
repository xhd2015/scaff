# Scenario

**Feature**: fix git.ignore is idempotent

```
# complete .gitignore -> no-op on second run
git.ignore fix -> nothing to do
```

## Preconditions

- Project has complete Go `.gitignore`.

## Steps

1. Write complete `.gitignore`.
2. Run `scaff fix git.ignore`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeCompleteGoGitignore(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "git.ignore"}
	return nil
}
```