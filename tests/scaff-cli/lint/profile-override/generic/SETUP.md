# Scenario

**Feature**: lint generic profile uses universal patterns only

```
# no go.mod or package.json -> generic profile
git.ignore -> only .DS_Store, .vscode/, *.swp, *~ patterns
```

## Preconditions

- Project has no `go.mod` or `package.json`.

## Steps

1. Leave project directory empty (or with a README only).
2. Run `scaff lint`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeFile(req.ProjectDir, "README.md", "# app\n"); err != nil {
		return err
	}
	req.Args = []string{"lint"}
	return nil
}
```