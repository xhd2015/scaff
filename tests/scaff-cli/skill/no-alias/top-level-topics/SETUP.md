# Scenario

**Feature**: `scaff topics` is not a skill-list alias

```
# top-level topics must not list skill topics
scaff topics -> unknown command, exit non-zero
```

## Preconditions

- No top-level `topics` command is registered.

## Steps

1. Run `scaff topics`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"topics"}
	return nil
}
```
