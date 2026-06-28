# Scenario

**Feature**: fix script.build is idempotent

```
# existing build stub -> no-op
script.build fix -> nothing to do
```

## Preconditions

- `script/build/build.go` already exists.

## Steps

1. Write existing stub.
2. Run `scaff fix script.build`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeScriptBuild(req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "script.build"}
	return nil
}
```