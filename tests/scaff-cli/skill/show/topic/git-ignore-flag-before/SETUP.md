# Scenario

**Feature**: show `git/ignore` with `--show` before path

```
# flag-before-path order
scaff skill --show git/ignore -> name: scaff/git/ignore / rule git/ignore
```

## Preconditions

- Topic `git/ignore` is embedded.

## Steps

1. Run `scaff skill --show git/ignore`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"skill", "--show", "git/ignore"}
	return nil
}
```
