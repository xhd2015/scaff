# Scenario

**Feature**: fix install/via-curl creates installer script

```
# no install-via-curl.sh -> create bash installer with GitHub release URLs
install/via-curl fix -> install-via-curl.sh at repo root
```

## Preconditions

- Project has `go.mod` and no `install-via-curl.sh`.

## Steps

1. Run `scaff fix install/via-curl`.

```go
func Setup(t *testing.T, req *Request) error {
	markInstallViaCurlTree()
	markFixTree()
	req.Args = []string{"fix", "install/via-curl"}
	return nil
}
```