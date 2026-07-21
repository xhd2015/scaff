# Scenario

**Feature**: scaff fix script/build

```
# create script/build/build.go stub when missing
fix executor -> script/build -> native go build helper stub
```

## Preconditions

- Project directory exists with `go.mod`.

## Steps

1. Materialize `script/build/build.go` state for the scenario.
2. Run `scaff fix script/build` with case-specific flags.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

// markScriptBuildTree keeps hierarchical child packages importing this package live.
func markScriptBuildTree() {}
```