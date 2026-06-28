# Scenario

**Feature**: fix script.build creates stub

```
# missing script/build/build.go -> create build helper stub
script.build fix -> script/build/build.go
```

## Preconditions

- `script/build/build.go` does not exist.

## Steps

1. Run `scaff fix script.build`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "script.build"}
	return nil
}
```