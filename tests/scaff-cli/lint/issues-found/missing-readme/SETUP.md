# Scenario

**Feature**: lint reports missing README.md

```
# go.mod only, no README.md -> project/readme missing
Rule project/readme -> missing status for README.md
```

## Preconditions

- Project has `go.mod` and no `README.md`.

## Steps

1. Write `go.mod` to the project directory.
2. Run `scaff lint`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"lint"}
	return nil
}
```