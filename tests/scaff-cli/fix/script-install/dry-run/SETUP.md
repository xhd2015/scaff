# Scenario

**Feature**: fix script/install --dry-run previews create

```
# --dry-run shows would-create without writing
script/install fix --dry-run -> stdout preview, install.go unchanged
```

## Preconditions

- `script/install/install.go` does not exist.

## Steps

1. Run `scaff fix script/install --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markScriptInstallTree()
	markFixTree()
	req.Args = []string{"fix", "script/install", "--dry-run"}
	return nil
}
```