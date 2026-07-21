# Scenario

**Feature**: fix github/release creates release scaffold

```
# no release scripts -> create main + lib with module substitutions
github/release fix -> script/github/release + script/github/lib
```

## Preconditions

- Project has `go.mod` (`github.com/xhd2015/myapp`) and no release scaffold files.

## Steps

1. Ensure `script/github/release/main.go` and `script/github/lib/build_release.go` are absent.
2. Run `scaff fix github/release`.

```go
func Setup(t *testing.T, req *Request) error {
	markGithubReleaseTree()
	markFixTree()
	req.Args = []string{"fix", "github/release"}
	return nil
}
```