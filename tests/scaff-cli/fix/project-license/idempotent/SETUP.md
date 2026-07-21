# Scenario

**Feature**: fix project/license no-op when LICENSE exists

```
# LICENSE already present -> no overwrite
project/license fix -> nothing to do
```

## Preconditions

- `LICENSE` already exists with a custom marker.

## Steps

1. Write custom `LICENSE`.
2. Run `scaff fix project/license`.

```go
func Setup(t *testing.T, req *Request) error {
	markProjectLicenseTree()
	markFixTree()
	if err := writeLICENSE(req.ProjectDir, "CUSTOM_LICENSE\n\nExisting license text.\n"); err != nil {
		return err
	}
	req.Args = []string{"fix", "project/license"}
	return nil
}
```