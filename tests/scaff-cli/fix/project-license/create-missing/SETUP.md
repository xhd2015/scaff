# Scenario

**Feature**: fix project/license creates LICENSE

```
# no LICENSE -> create MIT license with metadata substitutions
project/license fix -> LICENSE with year and owner
```

## Preconditions

- Project has `go.mod` (`github.com/xhd2015/myapp`) and no `LICENSE`.

## Steps

1. Ensure `LICENSE` is absent.
2. Run `scaff fix project/license`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "project/license"}
	return nil
}
```