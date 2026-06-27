# Scenario

**Feature**: scaff lint JSON report

```
# structured JSON output for automation consumers
scaff lint --json -> LintReport JSON on stdout
```

## Preconditions

- The project has lint issues to report.

## Steps

1. Prepare a project with scaffold gaps.
2. Run `scaff lint --json`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"lint", "--json"}
	return nil
}
```