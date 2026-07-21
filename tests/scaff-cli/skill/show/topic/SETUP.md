# Scenario

**Feature**: show nested topic by slash path (both flag orders)

```
# topic path resolves to docs/<path>/TOPIC.md
scaff skill --show <path> | scaff skill <path> --show -> topic body
```

## Preconditions

- Nested topics use `TOPIC.md` (not nested `SKILL.md`).
- Frontmatter `name` uses `scaff/<path>` form.

## Steps

1. Leaf sets path and flag order; Assert checks identity markers.

```go
func Setup(t *testing.T, req *Request) error {
	markShowTree()
	markSkillTree()
	// Leaves set Args with topic path and --show order.
	if req.RunDir == "" {
		req.RunDir = req.ProjectDir
	}
	return nil
}

// markShowTopicTree keeps hierarchical child packages importing this package live.
func markShowTopicTree() {}
```
