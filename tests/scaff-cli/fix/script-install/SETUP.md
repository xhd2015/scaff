# Scenario

**Feature**: scaff fix script/install

```
# create script/install/install.go stub when missing
fix executor -> script/install -> build-then-install helper stub
```

## Preconditions

- Project directory exists with `go.mod`.

## Steps

1. Materialize `script/install/install.go` state for the scenario.
2. Run `scaff fix script/install` with case-specific flags.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

```