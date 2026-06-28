# Scenario

**Feature**: fix script.bundle.for-linux creates stub

```
# missing script/bundle/for-linux/main.go -> create bundle helper stub
script.bundle.for-linux fix -> script/bundle/for-linux/main.go
```

## Preconditions

- `script/bundle/for-linux/main.go` does not exist.

## Steps

1. Run `scaff fix script.bundle.for-linux`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "script.bundle.for-linux"}
	return nil
}
```