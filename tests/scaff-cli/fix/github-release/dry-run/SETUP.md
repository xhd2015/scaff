# Scenario

**Feature**: fix github/release --dry-run

```
# --dry-run reports would-create without writing release files
github/release fix --dry-run -> preview only
```

## Preconditions

- Project has no release scaffold files.

## Steps

1. Run `scaff fix github/release --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "github/release", "--dry-run"}
	return nil
}
```