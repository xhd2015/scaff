# Scenario

**Feature**: fix script/github/release-assets is idempotent

```
# existing release-assets stub -> no-op
script/github/release-assets fix -> nothing to do
```

## Preconditions

- `script/github/release-assets/main.go` already exists.

## Steps

1. Write existing stub (marker content).
2. Run `scaff fix script/github/release-assets`.

```go
func Setup(t *testing.T, req *Request) error {
	markScriptGithubReleaseAssetsTree()
	markFixTree()
	if err := writeScriptGithubReleaseAssets(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "script/github/release-assets"}
	return nil
}
```
