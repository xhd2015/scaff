# Scenario

**Feature**: fix git/pre-commit is idempotent

```
# existing pre-commit helper -> no-op
git/pre-commit fix -> nothing to do
```

## Preconditions

- `script/git/pre-commit/main.go` already exists.

## Steps

1. Write existing pre-commit helper.
2. Run `scaff fix git/pre-commit`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGitPreCommitMain(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "git/pre-commit"}
	return nil
}
```
