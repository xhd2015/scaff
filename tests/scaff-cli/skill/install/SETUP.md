# Scenario

**Feature**: `scaff skill --install` installs embedded skill files

```
# skillcmd install with optional --dry-run and target dir
scaff skill --install [--dry-run] [<dir>] -> plan or write SKILL.md + TOPIC.md tree
```

## Preconditions

- Skill install uses skill name `scaff` and nested `TOPIC.md` extras from TreeFS.

## Steps

1. Leaf chooses dry-run / target dir args.

```go
func Setup(t *testing.T, req *Request) error {
	// Leaves set --install flags; prefer --dry-run + positional dir.
	if req.RunDir == "" {
		req.RunDir = req.ProjectDir
	}
	return nil
}

```
