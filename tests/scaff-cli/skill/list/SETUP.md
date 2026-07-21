# Scenario

**Feature**: `scaff skill --list` enumerates skill name and topic paths

```
# list action prints skill name then sorted nested topic paths
scaff skill --list -> "scaff\n" + topic paths (one per line)
```

## Preconditions

- Multi-topic TreeFS is wired with the full topic inventory.

## Steps

1. Run `scaff skill --list`.

```go
func Setup(t *testing.T, req *Request) error {
	markSkillTree()
	req.Args = []string{"skill", "--list"}
	return nil
}

// markListTree keeps hierarchical child packages importing this package live.
func markListTree() {}
```
