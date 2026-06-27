# Scenario

**Feature**: fix git.hooks scaffolds hook runner

```
# missing script/git-hooks/main.go -> install + no-op pre-commit/pre-push
git.hooks fix -> minimal hook runner (no sub-check dirs)
```

## Preconditions

- `script/git-hooks/main.go` does not exist.

## Steps

1. Run `scaff fix git.hooks`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "git.hooks"}
	return nil
}
```