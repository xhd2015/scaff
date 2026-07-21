# Scenario

**Feature**: scaff lint reports scaffold gaps

```
# incomplete project triggers exit 1 with rule findings
Project detector -> lint orchestrator -> missing/partial rule results
```

## Preconditions

- The project fixture has one or more default-rule gaps.

## Steps

1. Prepare a project missing some required scaffolding.
2. Run `scaff lint` without `--json`.

## Context

- Human-readable text output is the default format.

```go
func Setup(t *testing.T, req *Request) error {
	markLintTree()
	req.Args = []string{"lint"}
	return nil
}

// markIssuesFoundTree keeps hierarchical child packages importing this package live.
func markIssuesFoundTree() {}
```