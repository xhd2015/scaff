# Scenario

**Feature**: fix script/install is idempotent

```
# existing install stub -> no-op
script/install fix -> nothing to do
```

## Preconditions

- `script/install/install.go` already exists.

## Steps

1. Write existing stub.
2. Run `scaff fix script/install`.

```go
func Setup(t *testing.T, req *Request) error {
	markScriptInstallTree()
	markFixTree()
	if err := writeScriptInstall(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "script/install"}
	return nil
}
```