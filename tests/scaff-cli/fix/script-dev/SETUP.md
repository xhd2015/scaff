# Scenario

**Feature**: scaff fix script/dev

```
# create script/dev/main.go go run . --dev wrapper when missing
fix executor -> script/dev -> dev helper stub
```

## Preconditions

- Project directory exists with `go.mod`.

## Steps

1. Materialize `script/dev/main.go` state for the scenario.
2. Run `scaff fix script/dev` with case-specific flags.

```go
func Setup(t *testing.T, req *Request) error {
	markFixTree()
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

// markScriptDevTree keeps hierarchical child packages importing this package live.
func markScriptDevTree() {}
```