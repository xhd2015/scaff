# Scenario

**Feature**: fix install/via-curl creates installer script

```
# no install.sh -> create bash installer with GitHub release URLs
install/via-curl fix -> install.sh at repo root
```

## Preconditions

- Project has `go.mod` and no `install.sh`.

## Steps

1. Run `scaff fix install/via-curl`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "install/via-curl"}
	return nil
}
```