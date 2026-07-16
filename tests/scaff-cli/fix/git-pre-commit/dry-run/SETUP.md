# Scenario

**Feature**: fix git/pre-commit --dry-run previews create

```
# --dry-run shows would create without writing
git/pre-commit fix --dry-run -> stdout preview, no file
```

## Preconditions

- `script/git/pre-commit/main.go` does not exist.

## Steps

1. Run `scaff fix git/pre-commit --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "git/pre-commit", "--dry-run"}
	return nil
}
```
