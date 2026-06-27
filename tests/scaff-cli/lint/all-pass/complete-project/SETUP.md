# Scenario

**Feature**: lint passes on complete scaffold

```
# complete gitignore + test.yml -> all default rules ok
lint orchestrator -> exit 0 all good
```

## Preconditions

- Project has `go.mod`, complete Go `.gitignore`, and `test.yml`.

## Steps

1. Write complete project fixtures.
2. Run `scaff lint`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	if err := writeCompleteGoGitignore(req.ProjectDir); err != nil {
		return err
	}
	if err := writeTestWorkflow(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"lint"}
	return nil
}
```