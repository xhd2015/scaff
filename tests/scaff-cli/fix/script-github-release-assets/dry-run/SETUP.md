# Scenario

**Feature**: fix script/github/release-assets --dry-run previews create

```
# --dry-run shows would-create without writing
script/github/release-assets fix --dry-run -> stdout preview, main.go unchanged
```

## Preconditions

- `script/github/release-assets/main.go` does not exist.

## Steps

1. Run `scaff fix script/github/release-assets --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markScriptGithubReleaseAssetsTree()
	markFixTree()
	req.Args = []string{"fix", "script/github/release-assets", "--dry-run"}
	return nil
}
```
