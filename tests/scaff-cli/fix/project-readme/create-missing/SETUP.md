# Scenario

**Feature**: fix project/readme creates README.md

```
# no README.md -> create with module substitutions
project/readme fix -> README.md with go install line
```

## Preconditions

- Project has `go.mod` (`github.com/xhd2015/myapp`) and no `README.md`.

## Steps

1. Ensure `README.md` is absent.
2. Run `scaff fix project/readme`.

```go
func Setup(t *testing.T, req *Request) error {
	markProjectReadmeTree()
	markFixTree()
	req.Args = []string{"fix", "project/readme"}
	return nil
}
```