# Scenario

**Feature**: fix git/ignore --dry-run previews append

```
# --dry-run shows would-append lines without writing
git/ignore fix --dry-run -> stdout preview, .gitignore unchanged
```

## Preconditions

- Project has partial `.gitignore`.

## Steps

1. Write partial `.gitignore`.
2. Run `scaff fix git/ignore --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writePartialGitignore(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "git/ignore", "--dry-run"}
	return nil
}
```