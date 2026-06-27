# Scenario

**Feature**: git.hooks.install patches git hooks

```
# git repo + hook scaffold -> patch pre-commit and pre-push
git.hooks.install fix -> # scaff hooks marker in .git/hooks/
```

## Preconditions

- Project is a git repository.
- `script/git-hooks/main.go` exists.

## Steps

1. Initialize git repo and write hook scaffold.
2. Run `scaff fix git.hooks.install`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGitHooksMain(req.ProjectDir); err != nil {
		return err
	}
	if err := initGitRepo(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "git.hooks.install"}
	return nil
}
```