# Scenario

**Feature**: fix project/agents no-op when AGENTS.md exists

```
# AGENTS.md already present -> no overwrite
project/agents fix -> nothing to do
```

## Preconditions

- `AGENTS.md` already exists with a custom marker.

## Steps

1. Write custom `AGENTS.md`.
2. Run `scaff fix project/agents`.

```go
func Setup(t *testing.T, req *Request) error {
	markProjectAgentsTree()
	markFixTree()
	if err := writeAGENTS(req.ProjectDir, "# CUSTOM_AGENTS\n\nExisting agent instructions.\n"); err != nil {
		return err
	}
	req.Args = []string{"fix", "project/agents"}
	return nil
}
```