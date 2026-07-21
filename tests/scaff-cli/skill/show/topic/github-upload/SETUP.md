# Scenario

**Feature**: show `github/upload` topic (docs-only)

```
scaff skill --show github/upload -> name: scaff/github/upload
```

## Preconditions

- Topic `github/upload` is embedded (docs-only; no new fix rule).

## Steps

1. Run `scaff skill --show github/upload`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"skill", "--show", "github/upload"}
	return nil
}
```
