# Scenario

**Feature**: lint reports missing LICENSE

```
# go.mod only, no LICENSE -> project/license missing
Rule project/license -> missing status for LICENSE
```

## Preconditions

- Project has `go.mod` and no `LICENSE`.

## Steps

1. Write `go.mod` to the project directory.
2. Run `scaff lint`.

```go
func Setup(t *testing.T, req *Request) error {
	markIssuesFoundTree()
	markLintTree()
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"lint"}
	return nil
}
```