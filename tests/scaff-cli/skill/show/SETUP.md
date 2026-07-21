# Scenario

**Feature**: `scaff skill --show` prints root SKILL.md or nested TOPIC.md

```
# show root or topic-path content from embedded TreeFS
scaff skill --show [path] | scaff skill <path> --show -> skill body on stdout
```

## Preconditions

- Root `docs/SKILL.md` and nested `docs/<path>/TOPIC.md` are embedded.

## Steps

1. Descendant leaves set topic path, flag order, and optional `--header`.

```go
func Setup(t *testing.T, req *Request) error {
	// Keep parent skill package import live under hierarchical gen.
	// Leaves narrow Args for root, topic, header, or unknown path.
	if req.RunDir == "" {
		req.RunDir = req.ProjectDir
	}
	return nil
}

```
