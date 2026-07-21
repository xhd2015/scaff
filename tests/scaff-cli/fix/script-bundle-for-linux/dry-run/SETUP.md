# Scenario

**Feature**: fix script/bundle/for-linux --dry-run previews create

```
# --dry-run shows would-create without writing
script/bundle/for-linux fix --dry-run -> stdout preview, main.go unchanged
```

## Preconditions

- `script/bundle/for-linux/main.go` does not exist.

## Steps

1. Run `scaff fix script/bundle/for-linux --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "script/bundle/for-linux", "--dry-run"}
	return nil
}
```