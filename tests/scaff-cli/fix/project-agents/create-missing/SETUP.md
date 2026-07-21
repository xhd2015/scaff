# Scenario

**Feature**: fix project/agents creates AGENTS.md

```
# no AGENTS.md -> create with module substitutions
project/agents fix -> AGENTS.md with build and test sections
```

## Preconditions

- Project has `go.mod` (`github.com/xhd2015/myapp`) and no `AGENTS.md`.

## Steps

1. Ensure `AGENTS.md` is absent.
2. Run `scaff fix project/agents`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "project/agents"}
	return nil
}
```