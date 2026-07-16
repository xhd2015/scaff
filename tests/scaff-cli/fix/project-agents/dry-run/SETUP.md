# Scenario

**Feature**: fix project/agents --dry-run

```
# --dry-run reports would-create without writing AGENTS.md
project/agents fix --dry-run -> preview only
```

## Preconditions

- Project has no `AGENTS.md`.

## Steps

1. Run `scaff fix project/agents --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "project/agents", "--dry-run"}
	return nil
}
```