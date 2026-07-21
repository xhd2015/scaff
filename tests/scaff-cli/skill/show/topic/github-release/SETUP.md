# Scenario

**Feature**: show `github/release` topic

```
scaff skill --show github/release -> name: scaff/github/release / github/release
```

## Preconditions

- Topic `github/release` is embedded.

## Steps

1. Run `scaff skill --show github/release`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"skill", "--show", "github/release"}
	return nil
}
```
