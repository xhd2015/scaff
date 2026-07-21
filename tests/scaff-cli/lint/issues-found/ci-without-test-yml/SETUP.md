# Scenario

**Feature**: ci.yml does not satisfy github/testing-workflow

```
# ci.yml present but test.yml missing -> workflow rule still fails
Rule github/testing-workflow -> requires test.yml specifically
```

## Preconditions

- Project has complete `.gitignore`, `go.mod`, and `ci.yml` but no `test.yml`.

## Steps

1. Write `go.mod`, complete `.gitignore`, and `ci.yml`.
2. Run `scaff lint`.

```go
func Setup(t *testing.T, req *Request) error {
	markIssuesFoundTree()
	markLintTree()
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	if err := writeCompleteGoGitignore(req.ProjectDir); err != nil {
		return err
	}
	if err := writeCiWorkflow(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"lint"}
	return nil
}
```