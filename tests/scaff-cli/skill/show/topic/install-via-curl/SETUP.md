# Scenario

**Feature**: show `install-via-curl` topic

```
scaff skill --show install-via-curl -> name: scaff/install-via-curl / install/via-curl
```

## Preconditions

- Topic `install-via-curl` is embedded.

## Steps

1. Run `scaff skill --show install-via-curl`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"skill", "--show", "install-via-curl"}
	return nil
}
```
