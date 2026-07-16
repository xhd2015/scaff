# Scenario

**Feature**: fix tests/doctest --dry-run

```
# --dry-run reports would-create without writing doctest tree
tests/doctest fix --dry-run -> preview only
```

## Preconditions

- Project has no `tests/myapp-cli/` tree.

## Steps

1. Run `scaff fix tests/doctest --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "tests/doctest", "--dry-run"}
	return nil
}
```