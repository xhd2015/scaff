# Scenario

**Feature**: git/hooks/install requires scaffold

```
# no script/git-hooks/main.go -> error with hint
git/hooks/install fix -> exit 1, hint scaff fix git/hooks
```

## Preconditions

- `script/git-hooks/main.go` does not exist.

## Steps

1. Run `scaff fix git/hooks/install`.

```go
func Setup(t *testing.T, req *Request) error {
	markGitHooksInstallTree()
	markFixTree()
	req.Args = []string{"fix", "git/hooks/install"}
	return nil
}
```