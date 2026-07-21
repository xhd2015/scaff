# Scenario

**Feature**: fix git/ignore never adds vendor/

```
# vendored deps may be committed; vendor/ is excluded from patterns
git/ignore fix -> no vendor/ line
```

## Preconditions

- Project has no `.gitignore`.

## Steps

1. Run `scaff fix git/ignore`.

```go
func Setup(t *testing.T, req *Request) error {
	markGitIgnoreTree()
	markFixTree()
	req.Args = []string{"fix", "git/ignore"}
	return nil
}
```