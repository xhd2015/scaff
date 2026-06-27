# Scenario

**Feature**: scaff fix script.generate

```
# create script/generate/main.go stub when missing
fix executor -> script.generate -> no-op stub
```

## Preconditions

- Project directory exists.

## Steps

1. Materialize `script/generate/main.go` state for the scenario.
2. Run `scaff fix script.generate`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}
```