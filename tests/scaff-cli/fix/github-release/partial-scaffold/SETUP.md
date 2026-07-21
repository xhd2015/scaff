# Scenario

**Feature**: fix github/release creates missing lib only

```
# release main exists, lib missing -> create lib without overwriting main
github/release fix -> append missing build_release.go
```

## Preconditions

- `script/github/release/main.go` exists with custom content.
- `script/github/lib/build_release.go` is absent.

## Steps

1. Write custom release `main.go`.
2. Run `scaff fix github/release`.

```go
func Setup(t *testing.T, req *Request) error {
	markGithubReleaseTree()
	markFixTree()
	customMain := "// CUSTOM_RELEASE_MAIN\n"
	if err := writeGithubReleaseMain(req.ProjectDir, customMain); err != nil {
		return err
	}
	req.Args = []string{"fix", "github/release"}
	return nil
}
```