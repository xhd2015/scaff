# Scenario

**Feature**: fix project/layout/cmd is idempotent when cmd/ exists

```
# cmd/ already present -> no-op
project/layout/cmd fix -> nothing to do
```

## Preconditions

- `cmd/myapp/main.go` already exists with a custom marker.

## Steps

1. Write existing cmd entry main.
2. Run `scaff fix project/layout/cmd`.

```go
func Setup(t *testing.T, req *Request) error {
	markProjectLayoutCmdTree()
	markFixTree()
	if err := writeCmdMyappCustom(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "project/layout/cmd"}
	return nil
}
```