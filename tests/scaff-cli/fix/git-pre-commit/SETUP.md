# Scenario

**Feature**: scaff fix git/pre-commit

```
# scaffold script/git/pre-commit/main.go ensure + git add helper
fix executor -> git/pre-commit -> pre-commit script
```

## Preconditions

- Project directory exists.

## Steps

1. Materialize `script/git/pre-commit/main.go` state for the scenario.
2. Run `scaff fix git/pre-commit`.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

// markGitPreCommitTree keeps hierarchical child packages importing this package live.
func markGitPreCommitTree() {}
```
