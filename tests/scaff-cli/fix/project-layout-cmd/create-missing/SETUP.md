# Scenario

**Feature**: fix project/layout/cmd creates cmd entry

```
# no cmd/ directory -> create cmd/myapp/main.go
project/layout/cmd fix -> cmd/myapp/main.go
```

## Preconditions

- `cmd/` directory does not exist.

## Steps

1. Run `scaff fix project/layout/cmd`.

```go
func Setup(t *testing.T, req *Request) error {
	markProjectLayoutCmdTree()
	markFixTree()
	req.Args = []string{"fix", "project/layout/cmd"}
	return nil
}
```