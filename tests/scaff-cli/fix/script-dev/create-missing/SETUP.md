# Scenario

**Feature**: fix script/dev creates dev wrapper stub

```
# missing script/dev/main.go -> create go run . --dev wrapper
script/dev fix -> script/dev/main.go
```

## Preconditions

- `script/dev/main.go` does not exist.

## Steps

1. Run `scaff fix script/dev`.

```go
func Setup(t *testing.T, req *Request) error {
	markScriptDevTree()
	markFixTree()
	req.Args = []string{"fix", "script/dev"}
	return nil
}
```