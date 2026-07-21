# Scenario

**Feature**: lint reports missing doctest tree

```
# go.mod only, no tests/<name>-cli/DOCTEST.md -> tests/doctest missing
Rule tests/doctest -> missing status for tests/myapp-cli/DOCTEST.md
```

## Preconditions

- Project has `go.mod` (`github.com/xhd2015/myapp`) and no `tests/myapp-cli/DOCTEST.md`.

## Steps

1. Write `go.mod` to the project directory.
2. Run `scaff lint`.

```go
func Setup(t *testing.T, req *Request) error {
	markIssuesFoundTree()
	markLintTree()
	if err := writeGoModGitHubScaffold(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"lint"}
	return nil
}
```