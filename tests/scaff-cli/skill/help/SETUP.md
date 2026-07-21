# Scenario

**Feature**: `scaff skill --help` documents skill surface and topics

```
# skill-level help from skillcmd (+ Available topics index)
scaff skill --help -> usage + --show/--install/--list + topic index
```

## Preconditions

- Skill subcommand is registered and TreeFS topics are available for the index.

## Steps

1. Leaf runs `scaff skill --help`.

```go
func Setup(t *testing.T, req *Request) error {
	markSkillTree()
	req.Args = []string{"skill", "--help"}
	return nil
}

// markHelpTree keeps hierarchical child packages importing this package live.
func markHelpTree() {}
```
