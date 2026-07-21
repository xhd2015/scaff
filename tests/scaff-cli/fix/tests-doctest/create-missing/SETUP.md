# Scenario

**Feature**: fix tests/doctest creates doctest tree

```
# no tests/<name>-cli/ -> create DOCTEST.md + SETUP.md
tests/doctest fix -> tests/myapp-cli/ with module substitutions
```

## Preconditions

- Project has `go.mod` (`github.com/xhd2015/myapp`) and no `tests/myapp-cli/` tree.

## Steps

1. Ensure doctest tree is absent.
2. Run `scaff fix tests/doctest`.

```go
func Setup(t *testing.T, req *Request) error {
	markTestsDoctestTree()
	markFixTree()
	req.Args = []string{"fix", "tests/doctest"}
	return nil
}
```