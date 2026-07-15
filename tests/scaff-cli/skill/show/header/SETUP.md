# Scenario

**Feature**: `--show --header` prints only YAML frontmatter delimiters

```
# header-only mode strips skill/topic body
scaff skill --show --header -> ---\n<header>\n---\n (no body)
```

## Preconditions

- Root skill has YAML frontmatter with `name: scaff`.

## Steps

1. Run `scaff skill --show --header`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"skill", "--show", "--header"}
	return nil
}
```
