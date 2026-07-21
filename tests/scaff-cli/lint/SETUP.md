# Scenario

**Feature**: scaff lint read-only audit

```
# project detector + lint orchestrator evaluate default rules
scaff lint [flags] -> LintReport -> stdout/stderr + exit code
```

## Preconditions

- The scaff binary and temp project directory are ready from the root setup.

## Steps

1. Descendant setups materialize the project fixture for the lint scenario.
2. Run `scaff lint` with case-specific flags.

## Context

- Default lint rules: `git/ignore`, `github/testing-workflow` only.
- Exit 0 when all default rules pass; exit 1 when issues are found.

```go
func Setup(t *testing.T, req *Request) error {
	if req.Args == nil {
		req.Args = []string{"lint"}
	}
	return nil
}

```