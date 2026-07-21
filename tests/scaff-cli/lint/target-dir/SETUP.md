# Scenario

**Feature**: scaff lint with --dir

```
# audit a subdirectory as the project root
scaff lint --dir <subdir> -> rules evaluated relative to subdir
```

## Preconditions

- A nested directory contains its own project markers.

## Steps

1. Prepare a parent directory with a nested project subdirectory.
2. Run `scaff lint --dir <subdir>` from the parent.

```go
func Setup(t *testing.T, req *Request) error {
	markLintTree()
	req.RunDir = req.ProjectDir
	return nil
}

// markTargetDirTree keeps hierarchical child packages importing this package live.
func markTargetDirTree() {}
```