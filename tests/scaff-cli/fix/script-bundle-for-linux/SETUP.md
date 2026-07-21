# Scenario

**Feature**: scaff fix script/bundle/for-linux

```
# create script/bundle/for-linux/main.go stub when missing
fix executor -> script/bundle/for-linux -> cross-compile bundle helper stub
```

## Preconditions

- Project directory exists with `go.mod`.

## Steps

1. Materialize `script/bundle/for-linux/main.go` state for the scenario.
2. Run `scaff fix script/bundle/for-linux` with case-specific flags.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

// markScriptBundleForLinuxTree keeps hierarchical child packages importing this package live.
func markScriptBundleForLinuxTree() {}
```