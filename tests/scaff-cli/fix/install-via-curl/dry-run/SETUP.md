# Scenario

**Feature**: fix install.via.curl --dry-run

```
# --dry-run reports would-create without writing installer
install.via.curl fix --dry-run -> preview only
```

## Preconditions

- Project has no `install-via-curl.sh`.

## Steps

1. Run `scaff fix install.via.curl --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "install.via.curl", "--dry-run"}
	return nil
}
```