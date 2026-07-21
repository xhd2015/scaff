# Scenario

**Feature**: fix git/pre-commit scaffolds pre-commit helper

```
# missing script/git/pre-commit/main.go -> ensure + git add stub
git/pre-commit fix -> brief pre-commit helper
```

## Preconditions

- `script/git/pre-commit/main.go` does not exist.

## Steps

1. Run `scaff fix git/pre-commit`.

```go
func Setup(t *testing.T, req *Request) error {
	markGitPreCommitTree()
	markFixTree()
	req.Args = []string{"fix", "git/pre-commit"}
	return nil
}
```
