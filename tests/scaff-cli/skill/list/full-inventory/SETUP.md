# Scenario

**Feature**: full topic inventory appears in `--list` output

```
# every path/TOPIC.md under the skill tree is listed once, sorted
scaff skill --list -> scaff + fix, git/*, github/*, install-via-curl, lint, overview, script/*
```

## Preconditions

- Parent sets `skill --list`.
- Product embeds all inventory topics as nested `TOPIC.md` files.

## Steps

1. Assert stdout matches the complete sorted inventory (leaf Assert).

```go
func Setup(t *testing.T, req *Request) error {
	// Args inherited from skill/list parent: skill --list
	if len(req.Args) == 0 {
		req.Args = []string{"skill", "--list"}
	}
	return nil
}
```
