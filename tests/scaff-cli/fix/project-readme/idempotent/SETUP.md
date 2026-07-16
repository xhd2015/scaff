# Scenario

**Feature**: fix project/readme no-op when README exists

```
# README.md already present -> no overwrite
project/readme fix -> nothing to do
```

## Preconditions

- `README.md` already exists with a custom marker.

## Steps

1. Write custom `README.md`.
2. Run `scaff fix project/readme`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeREADME(req.ProjectDir, "# CUSTOM_README\n\nExisting content.\n"); err != nil {
		return err
	}
	req.Args = []string{"fix", "project/readme"}
	return nil
}
```