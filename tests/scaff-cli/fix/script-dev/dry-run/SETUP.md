# Scenario

**Feature**: fix script/dev --dry-run previews create

```
# --dry-run shows would-create without writing
script/dev fix --dry-run -> stdout preview, dev main unchanged
```

## Preconditions

- `script/dev/main.go` does not exist.

## Steps

1. Run `scaff fix script/dev --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markScriptDevTree()
	markFixTree()
	req.Args = []string{"fix", "script/dev", "--dry-run"}
	return nil
}
```