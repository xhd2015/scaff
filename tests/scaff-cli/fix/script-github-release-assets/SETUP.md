# Scenario

**Feature**: scaff fix script/github/release-assets

```
# create script/github/release-assets/main.go stub when missing
fix executor -> script/github/release-assets -> release-assets helper stub
```

## Preconditions

- Project directory exists with `go.mod`.
- Rule is opt-in fix only (not default lint).
- Related to `github/release` but a separate scaffold under `script/github/release-assets/`.

## Steps

1. Materialize `script/github/release-assets/main.go` state for the scenario.
2. Run `scaff fix script/github/release-assets` with case-specific flags.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	return nil
}
```
