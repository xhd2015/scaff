# Scenario

**Feature**: fix script.generate creates stub

```
# missing script/generate/main.go -> create no-op stub
script.generate fix -> generator entrypoint stub
```

## Preconditions

- `script/generate/main.go` does not exist.

## Steps

1. Run `scaff fix script.generate`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "script.generate"}
	return nil
}
```