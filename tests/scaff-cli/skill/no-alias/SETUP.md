# Scenario

**Feature**: no top-level install/topics aliases (skill-only surface)

```
# Shape 3 skill-only CLI: install/topics are unknown top-level commands
scaff install | scaff topics -> unknown command (not skill install/list)
```

## Preconditions

- Product intentionally omits top-level `install` and `topics` aliases.
- Skill install/list remain under `scaff skill …`.

## Steps

1. Leaf runs a forbidden top-level command.

```go
func Setup(t *testing.T, req *Request) error {
	// Leaves set Args to top-level install or topics (not under skill).
	if req.RunDir == "" {
		req.RunDir = req.ProjectDir
	}
	return nil
}
```
