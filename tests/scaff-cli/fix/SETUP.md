# Scenario

**Feature**: scaff fix applies one rule

```
# fix executor repairs a single slash-form rule idempotently
scaff fix <rule> [flags] -> filesystem changes + stdout
```

## Preconditions

- The scaff binary and temp project directory are ready from the root setup.

## Steps

1. Descendant setups materialize the project fixture and select the rule.
2. Run `scaff fix <rule>` with case-specific flags.

## Context

- Exit 0 on success or no-op; exit 1 on fix failure; exit 2 on unknown rule.

```go
func Setup(t *testing.T, req *Request) error {
	if req.RunDir == "" {
		req.RunDir = req.ProjectDir
	}
	return nil
}

```