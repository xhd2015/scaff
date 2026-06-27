# Scenario

**Feature**: git.hooks.install requires git repo

```
# hook scaffold present but no .git -> error
git.hooks.install fix -> non-git directory error
```

## Preconditions

- `script/git-hooks/main.go` exists.
- Directory is **not** a git repository.

## Steps

1. Write hook scaffold without `git init`.
2. Run `scaff fix git.hooks.install`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGitHooksMain(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "git.hooks.install"}
	return nil
}
```