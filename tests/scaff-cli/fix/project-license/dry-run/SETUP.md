# Scenario

**Feature**: fix project/license --dry-run

```
# --dry-run reports would-create without writing LICENSE
project/license fix --dry-run -> preview only
```

## Preconditions

- Project has no `LICENSE`.

## Steps

1. Run `scaff fix project/license --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markProjectLicenseTree()
	markFixTree()
	req.Args = []string{"fix", "project/license", "--dry-run"}
	return nil
}
```