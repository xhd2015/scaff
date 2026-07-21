# Scenario

**Feature**: show root skill body (multi-topic index)

```
# no topic path -> root SKILL.md
scaff skill --show -> frontmatter name: scaff + index body with retrieve examples
```

## Preconditions

- Root skill documents multi-topic retrieval via `scaff skill --show`.
- Root body must not document install target flags (`--cursor`, `--global`).

## Steps

1. Run `scaff skill --show`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"skill", "--show"}
	return nil
}
```
