# Scenario

**Feature**: fix github.release no-op when scaffold exists

```
# release main + lib already present -> no overwrite
github.release fix -> nothing to do
```

## Preconditions

- Both release scaffold files already exist with custom markers.

## Steps

1. Write custom release `main.go` and `build_release.go`.
2. Run `scaff fix github.release`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGithubReleaseMain(req.ProjectDir, "// CUSTOM_RELEASE_MAIN\n"); err != nil {
		return err
	}
	if err := writeGithubReleaseLib(req.ProjectDir, "package lib\n\n// CUSTOM_RELEASE_LIB\n"); err != nil {
		return err
	}
	req.Args = []string{"fix", "github.release"}
	return nil
}
```