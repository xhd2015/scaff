# Scenario

**Feature**: fix project/layout/cmd --dry-run previews create

```
# --dry-run shows would-create without writing
project/layout/cmd fix --dry-run -> stdout preview, cmd main unchanged
```

## Preconditions

- `cmd/` directory does not exist.

## Steps

1. Run `scaff fix project/layout/cmd --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markProjectLayoutCmdTree()
	markFixTree()
	req.Args = []string{"fix", "project/layout/cmd", "--dry-run"}
	return nil
}
```