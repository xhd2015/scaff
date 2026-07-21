# Scenario

**Feature**: `--install --dry-run` into a temp dir previews skill layout

```
# dry-run to positional dir (skillcmd) without writing
scaff skill --install --dry-run <ProjectDir>
  -> [dry-run] create SKILL.md + nested path/TOPIC.md under dir
```

## Preconditions

- `req.ProjectDir` is an empty temp directory used as the install target.
- Prefer dry-run to avoid home / `.agents` pollution.

## Steps

1. Run `scaff skill --install --dry-run` with `req.ProjectDir` as positional dir.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"skill", "--install", "--dry-run", req.ProjectDir}
	return nil
}
```
