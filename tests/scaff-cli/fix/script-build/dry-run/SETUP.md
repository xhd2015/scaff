# Scenario

**Feature**: fix script/build --dry-run previews create

```
# --dry-run shows would-create without writing
script/build fix --dry-run -> stdout preview, build.go unchanged
```

## Preconditions

- `script/build/build.go` does not exist.

## Steps

1. Run `scaff fix script/build --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "script/build", "--dry-run"}
	return nil
}
```