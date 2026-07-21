# Scenario

**Feature**: lint --profile node checks node_modules

```
# polyglot signals overridden to node profile
--profile node -> git/ignore expects node_modules/
```

## Preconditions

- Project has both `go.mod` and `package.json` (would auto-detect polyglot).
- `--profile node` overrides detection.

## Steps

1. Write `go.mod` and `package.json`.
2. Run `scaff lint --profile node`.

```go
func Setup(t *testing.T, req *Request) error {
	markProfileOverrideTree()
	markLintTree()
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	if err := writePackageJSON(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"lint", "--profile", "node"}
	return nil
}
```