# Scenario

**Feature**: fix script/github/release-assets creates stub

```
# missing script/github/release-assets/main.go -> create release-assets helper
script/github/release-assets fix -> main.go with help + Proposed behavior
```

## Preconditions

- `script/github/release-assets/main.go` does not exist.

## Steps

1. Run `scaff fix script/github/release-assets`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "script/github/release-assets"}
	return nil
}
```
