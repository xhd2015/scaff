# Scenario

**Feature**: fix tests/doctest no-op when doctest tree exists

```
# DOCTEST.md already present -> no overwrite
tests/doctest fix -> nothing to do
```

## Preconditions

- `tests/myapp-cli/DOCTEST.md` already exists with a custom marker.

## Steps

1. Write custom doctest tree.
2. Run `scaff fix tests/doctest`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeDoctestTree(req.ProjectDir, "myapp"); err != nil {
		return err
	}
	custom := "# CUSTOM_DOCTEST\n\n" + readProjectFile(t, req, "tests/myapp-cli/DOCTEST.md")
	if err := writeFile(req.ProjectDir, "tests/myapp-cli/DOCTEST.md", custom); err != nil {
		return err
	}
	req.Args = []string{"fix", "tests/doctest"}
	return nil
}
```